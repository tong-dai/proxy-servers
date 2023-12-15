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
	Classes map[int]*CacheClassInfo
}

var load_balancer *LB = &LB{
	Servers: []*Server{
		{ServerURL: "http://localhost:7777", Index: 0, Classes: createServerClassInfo(3, 5)},
		{ServerURL: "http://localhost:8888", Index: 1, Classes: createServerClassInfo(3, 5)},
		{ServerURL: "http://localhost:9999", Index: 1, Classes: createServerClassInfo(3, 5)},
	},
	Current: 0,
}

func createServerClassInfo(numClasses int, capacity int) map[int]*CacheClassInfo {
	classes := make(map[int]*CacheClassInfo)
	for i := 0; i < numClasses; i++ {
		classes[i] = &CacheClassInfo{MaxEnrollment: capacity, Enrollment: 0}
	}
	return classes
}

func GetLB() *LB {
	return load_balancer
}

func (lb *LB) GetNextServer() *Server {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	lb.Current = (lb.Current + 1) % len(lb.Servers)
	return lb.Servers[lb.Current]
}

func (lb *LB) HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Do we get to handle request?")
	server := lb.GetNextServer()

	serverURL, err := url.Parse(server.ServerURL)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(serverURL)
	proxy.ServeHTTP(w, r)
}
