package main

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

func (lb *LB) getNextServer() *Server {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	lb.Current = (lb.Current + 1) % len(lb.Servers)
    return lb.Servers[lb.Current]
}

func (lb *LB) handleRequest(w http.ResponseWriter, r *http.Request) {

	server := lb.getNextServer()

	serverURL, err := url.Parse(server.ServerURL)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(serverURL)
	proxy.ServeHTTP(w, r)
}

func main() {
	lb := LB{
        Servers: []*Server{
            {ServerURL: "http://localhost:7777"},
            {ServerURL: "http://localhost:8888"},
			// {ServerURL: "http://localhost:9999"},
        },
    }

    http.HandleFunc("/", lb.handleRequest) 
    fmt.Println("Server is running on port 8080...")
    http.ListenAndServe(":8080", nil)
}
