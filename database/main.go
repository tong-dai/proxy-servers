package main

import (
	// "encoding/json"
	lb "cos316/td_ec_final_project/load_balancer"
	"fmt"
	"net/http"
	"sync"
)

type ClassInfo struct {
	maxEnrollment int
	enrollment    int
	enrolledIds   []int
}
type DB struct {
	c map[int]*ClassInfo
	sync.Mutex
}

func (db *DB) UpdateDB(studentID int, classNumber int, cacheIndex int) (bool, int) {
	db.Lock()
	defer db.Unlock()
	class := db.c[classNumber]
	if class.enrollment == class.maxEnrollment {
		// TODO send message to all caches saying that you need to update this
		fmt.Println("something went wrong")
		return false, classNumber
	} else {
		class.enrolledIds = append(class.enrolledIds, studentID)
		return true, -1
	}

}

func createServer(numCaches int, numClasses int) *lb.LB {
	server := new(lb.LB)
	// fills caches with requisite data
	server.Caches = make([]*lb.Cache, numCaches)
	for i := 0; i < len(server.Caches); i++ {
		currClassMap := make(map[int]*lb.ClassInfo)
		for j := 0; j < numClasses; j++ {
			currClassInfo := &lb.ClassInfo{Enrollment: 0, MaxEnrollment: 5, ClassNumber: i}
			currClassMap[j] = currClassInfo
		}
		server.Caches = append(server.Caches, &lb.Cache{Classes: currClassMap})
	}
	return server
}
func createDB(numClasses int) *DB {
	db := new(DB)
	db.c = make(map[int]*ClassInfo)
	for i := 0; i < numClasses; i++ {
		db.c[i] = &ClassInfo{maxEnrollment: 5, enrollment: 0, enrolledIds: make([]int, 0)}
	}
	return db
}

func main() {
	// TODO add code that fills the db with the requisite classes
	// db := createDB(5)
	// server := createServer(4, 5)
	fmt.Println("Server running")

	http.HandleFunc("/enroll", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("recevied request")
		// TODO implement reading the http request
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// cached, cacheIndex := server.SendReqToCache(1, 1)
		// if cached {
		// 	fmt.Println("it was cached")
		// 	// success response
		// 	w.Write(make([]byte, 1))
		// 	success, classNumber := db.UpdateDB(1, 1, cacheIndex)
		// 	// if the class is full, it should be deleted from the cache
		// 	if !success {
		// 		delete(server.Caches[cacheIndex].Classes, classNumber)
		// 		fmt.Printf("the backing store said class %v was full\n", classNumber)
		// 		// error response
		// 		w.Write(make([]byte, 5))
		// 	} else {
		// 		// no need to make an http response
		// 		fmt.Printf("the backing store said that class %v was not full\n", classNumber)
		// 	}
		// } else {
		// 	fmt.Println("it was not cached")
		// 	// error response
		// 	w.Write(make([]byte, 5))
		// }
		// var body map[string]interface{}
		// json.NewDecoder(r.Body).Decode(body)
		// fmt.Println("received response", body)
		var response []byte = make([]byte, 5)
		for i := 0; i < 5; i++ {
			response[i] = byte(i)
		}
		w.Write(response)
	})
	// http.HandleFunc("/enroll/:classNumber", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("Query params", r.URL.Query())
	// 	var response []byte = make([]byte, 10)
	// 	for i := 0; i < 10; i++ {
	// 		response[i] = byte(i)
	// 	}
	// 	w.Write(response)

	// })
	http.ListenAndServe(":8080", nil)
}
