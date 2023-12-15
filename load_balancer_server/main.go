package main

import (
	lb "cos316/td_ec_final_project/load_balancer"
	"fmt"
	"net/http"
)




func main() {
	load_balancer := lb.GetLB()
	fmt.Println(load_balancer.Servers)
	http.HandleFunc("/", load_balancer.HandleRequest)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
