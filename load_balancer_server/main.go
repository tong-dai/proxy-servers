package main

import (
	lb "cos316/td_ec_final_project/load_balancer"
	"fmt"
	"net/http"
)

func createServerClassInfo(numClasses int, capacity int) map[int]*lb.CacheClassInfo {
	classes := make(map[int]*lb.CacheClassInfo)
	for i := 0; i < numClasses; i++ {
		classes[i] = &lb.CacheClassInfo{MaxEnrollment: capacity, Enrollment: 0}
	}
	return classes
}



var Load_balancer = lb.LB{
	Servers: []*lb.Server{
		{ServerURL: "http://localhost:7777", Index: 0, Classes: createServerClassInfo(3, 5)},
		{ServerURL: "http://localhost:8888", Index: 1, Classes: createServerClassInfo(3, 5)},
		// {ServerURL: "http://localhost:9999"},
	},
}
func main() {

	http.HandleFunc("/", Load_balancer.HandleRequest)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
	