package main

import (
	"fmt"
	"net/http"
	"cos316/td_ec_final_project/package"
)



func main() {
	lb := load_balancer.LB{
        Servers: []*load_balancer.Server{
            {ServerURL: "http://localhost:7777"},
            {ServerURL: "http://localhost:8888"},
			{ServerURL: "http://localhost:9999"},
        },
    }

    http.HandleFunc("/", lb.HandleRequest) 
    fmt.Println("Server is running on port 8080...")
    http.ListenAndServe(":8080", nil)
}
