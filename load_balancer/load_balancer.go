package load_balancer

import (
	"fmt"
	"sync"
)

type LB struct {
	Caches []*Cache // represents the different caches
	Index int // represents the index to which to send the incoming request
	sync.Mutex

}
func (lb *LB) SendReqToCache(student int, class int) (bool, int) {
	lb.Lock()
	lb.Caches[0] = new(Cache)
	fmt.Println(lb.Index)
	lb.Index++
	lb.Unlock()
	// returning the cache to which the request was sent
	return true, lb.Index-1
}

