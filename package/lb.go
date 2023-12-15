package load_balancer

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

type Server struct {
	ServerURL	string
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