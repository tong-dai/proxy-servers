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