package main
import (
    // "fmt"
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

