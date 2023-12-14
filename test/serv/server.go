package main

import (
    "fmt"
    // "log"
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
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        enrollmentInfo := processQuery(r)

        fmt.Println(enrollmentInfo.StudentID, enrollmentInfo.ClassNum)
        fmt.Printf("Request received on port 7777")
        fmt.Fprintf(w, "Hello from server on port 7777")
    })

    http.ListenAndServe(":7777", nil)
}
