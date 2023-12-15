package main

import (
	// db "cos316/td_ec_final_project/database"
	// server "cos316/td_ec_final_project/load_balancer_server"

	// lb "cos316/td_ec_final_project/load_balancer"
	"fmt"
	"net/http"
	"strconv"
)

type EnrollmentInfo struct {
	StudentID string
	ClassNum  string
}

func processQuery(r *http.Request) bool {
	studentID, err := strconv.Atoi(r.URL.Query().Get("studentID"))
	if err != nil {
		fmt.Println("Something went wrong converting studentID to int")
	}
	classNum, err := strconv.Atoi(r.URL.Query().Get("classNum"))
	if err != nil {
		fmt.Println("Something went wrong converting classNum to id")
	}
	// success := db.UpdateDB(studentID, classNum, server.Load_balancer.Servers, server.Load_balancer.GetNextServerIndex())
	fmt.Printf("studentID %v", studentID)
	fmt.Printf("classNum %v", classNum)
	return true
	// return &EnrollmentInfo{StudentID: studentID, ClassNum: classNum}
}

func main() {

	// Create Main Server
	fmt.Println("Server started on: http://localhost:9000")
	main_server := http.NewServeMux()

	//Creating Sub Domains
	server1 := http.NewServeMux()
	server1.HandleFunc("/", handlerServer1)

	server2 := http.NewServeMux()
	server2.HandleFunc("/", handlerServer2)

	server3 := http.NewServeMux()
	server3.HandleFunc("/", handlerServer3)

	//Running Sub Domains
	go func() {
		fmt.Println("Server started on: http://localhost:7777")
		http.ListenAndServe(":7777", server1)
	}()

	go func() {
		fmt.Println("Server started on: http://localhost:8888")
		http.ListenAndServe(":8888", server2)
	}()

	go func() {
		fmt.Println("Server started on: http://localhost:9999")
		http.ListenAndServe(":9999", server2)
	}()

	//Running Main Server
	http.ListenAndServe("localhost:9000", main_server)
}

func handlerServer1(w http.ResponseWriter, r *http.Request) {
	success := processQuery(r)
	// if success {
	//     for key, value := range db.DB.C {

	//     }
	// }
	// fmt.Println(enrollmentInfo.StudentID, enrollmentInfo.ClassNum)
	fmt.Println(success)
	fmt.Println("Running on Port :7777")
}

func handlerServer2(w http.ResponseWriter, r *http.Request) {
	success := processQuery(r)
	fmt.Println(success)

	// fmt.Println(enrollmentInfo.StudentID, enrollmentInfo.ClassNum)
	fmt.Println("Running on Port :8888")
}

func handlerServer3(w http.ResponseWriter, r *http.Request) {
	success := processQuery(r)
	fmt.Println(success)

	// fmt.Println(enrollmentInfo.StudentID, enrollmentInfo.ClassNum)
	fmt.Println("Running on Port :9999")
}
