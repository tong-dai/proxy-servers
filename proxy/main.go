package main

import (
    "fmt"
    "net/http"
)


type EnrollmentInfo struct {
    StudentID   string
    ClassNum    string
}

func processQuery(r *http.Request) *EnrollmentInfo {
    studentID := r.URL.Query().Get("studentID")
    classNum := r.URL.Query().Get("classNum")
    return &EnrollmentInfo{StudentID: studentID, ClassNum: classNum}
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
    enrollmentInfo := processQuery(r)

    fmt.Println(enrollmentInfo.StudentID, enrollmentInfo.ClassNum)
    fmt.Println("Running on Port :7777")
}

func handlerServer2(w http.ResponseWriter, r *http.Request) {
    enrollmentInfo := processQuery(r)

    fmt.Println(enrollmentInfo.StudentID, enrollmentInfo.ClassNum)
	fmt.Println("Running on Port :8888")
}

func handlerServer3(w http.ResponseWriter, r *http.Request) {
    enrollmentInfo := processQuery(r)

    fmt.Println(enrollmentInfo.StudentID, enrollmentInfo.ClassNum)
	fmt.Println("Running on Port :9999")
}
