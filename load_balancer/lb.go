package lb
import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type LB struct {
	Servers		[]*Server
	Current		int
	mu			sync.Mutex
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


var lb *LB = &LB{
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

func (lb *LB) GetLB() *LB {
	return lb
}

func (lb *LB) GetNextServerIndex() int {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	lb.Current = (lb.Current + 1) % len(lb.Servers)
    return lb.Current
}

func (lb *LB) HandleRequest(w http.ResponseWriter, r *http.Request) {

	server := lb.Servers[lb.GetNextServerIndex()]

	serverURL, err := url.Parse(server.ServerURL)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(serverURL)
	proxy.ServeHTTP(w, r)
}