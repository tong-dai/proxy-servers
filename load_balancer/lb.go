package lb

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type LB struct {
	Servers []*Server
	Current int
	mu      sync.Mutex
}

type CacheClassInfo struct {
	MaxEnrollment int
	Enrollment    int
}

type Server struct {
	ServerURL string
	Index     int
	sync.Mutex
	Classes map[string]*CacheClassInfo
}

var load_balancer *LB = &LB{
	Servers: []*Server{
		{ServerURL: "http://localhost:7777", Index: 0, Classes: createServerClassInfo(1000, 15)},
		{ServerURL: "http://localhost:8888", Index: 1, Classes: createServerClassInfo(1000, 15)},
		{ServerURL: "http://localhost:9999", Index: 2, Classes: createServerClassInfo(1000, 15)},
	},
	Current: 0,
}

func createServerClassInfo(numClasses int, capacity int) map[string]*CacheClassInfo {
	classes := make(map[string]*CacheClassInfo)
	for i := 0; i < numClasses; i++ {
		classes[fmt.Sprint(i)] = &CacheClassInfo{MaxEnrollment: capacity, Enrollment: 0}
	}
	return classes
}

func GetLB() *LB {
	return load_balancer
}

func (lb *LB) DeleteClass(classNum string) {
	servers := lb.Servers
	for i := 0; i < len(servers); i++ {
		// servers[i].Lock()
		_, found := servers[i].Classes[classNum]
		if found {
			delete(servers[i].Classes, classNum)
		}
		// servers[i].Unlock()
	}
}

func (lb *LB) UpdateServer(classNum string, enrollment int, serverIndex int) bool {

	lb.Servers[serverIndex].Lock()
	class, found := lb.Servers[serverIndex].Classes[classNum]
	if !found {
		return false
	}
	if enrollment == class.MaxEnrollment {
		delete(lb.Servers[serverIndex].Classes, classNum)
	} else {
		lb.Servers[serverIndex].Classes[classNum].Enrollment = enrollment
	}
	lb.Servers[serverIndex].Unlock()
	return true

}

func (lb *LB) GetNextServer() *Server {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	lb.Current = (lb.Current + 1) % len(lb.Servers)
	return lb.Servers[lb.Current]
}

func (lb *LB) HandleRequest(w http.ResponseWriter, r *http.Request) {
	server := lb.GetNextServer()

	serverURL, err := url.Parse(server.ServerURL)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(serverURL)
	proxy.ServeHTTP(w, r)
}
