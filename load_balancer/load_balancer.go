package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendReqToCache(student int, class int) bool {
	return true
}

func main() {
	fmt.Println("Server running")
	http.HandleFunc("/enroll", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var body map[string]interface{}
		json.NewDecoder(r.Body).Decode(body)
		fmt.Println("received response", body)
		var response []byte = make([]byte, 5)
		for i := 0; i < 5; i++ {
			response[i] = byte(i)
		}
		w.Write(response)
	})
	http.HandleFunc("/enroll/:classNumber", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Query params", r.URL.Query())
		var response []byte = make([]byte, 10)
		for i := 0; i < 10; i++ {
			response[i] = byte(i)
		}
		w.Write(response)

	})
	http.ListenAndServe(":8080", nil)
}
