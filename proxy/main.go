package main

import (
	db "cos316/td_ec_final_project/database"
	lb "cos316/td_ec_final_project/load_balancer"
	"fmt"
	"net/http"
	"strings"
)

type EnrollmentInfo struct {
	StudentID string
	ClassNum  string
}

var load_balancer *lb.LB = lb.GetLB()
var database *db.DB = db.GetDB()

func processQuery(r *http.Request) *EnrollmentInfo {
	studentID := r.URL.Query().Get("studentID")
	classNum := r.URL.Query().Get("classNum")
	fmt.Println(r.URL.Query())
	fmt.Println("--------------")
	fmt.Printf("classnum: %s", classNum)
	fmt.Printf("studentId: %s", studentID)
	return &EnrollmentInfo{StudentID: studentID, ClassNum: classNum}
}

func main() {

	// Create Main Server
	fmt.Println("Server started on: http://localhost:9000")
	main_server := http.NewServeMux()

	// Creating Sub Domains
	server1 := http.NewServeMux()
	server1.HandleFunc("/", handlerServer1)

	server2 := http.NewServeMux()
	server2.HandleFunc("/", handlerServer2)

	server3 := http.NewServeMux()
	server3.HandleFunc("/", handlerServer3)

	// Running Sub Domains
	go func() {
		fmt.Println("Server started on: http://localhost:7777")
		http.ListenAndServe(":7777", server1)
	}()

	fmt.Println(load_balancer.Servers[0].Classes["0"].Enrollment)
	fmt.Println(load_balancer.Servers[0].Classes["0"].MaxEnrollment)

	go func() {
		fmt.Println("Server started on: http://localhost:8888")
		http.ListenAndServe(":8888", server2)
	}()

	go func() {
		fmt.Println("Server started on: http://localhost:9999")
		http.ListenAndServe(":9999", server3)
	}()

	// Running Main Server
	http.ListenAndServe("localhost:9000", main_server)
}

func handlerServer1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/event-stream")

	fmt.Println("server0 -----------")
	// eInfo := processQuery(r)
	eInfo := &EnrollmentInfo{StudentID: r.URL.Query().Get("studentID"), ClassNum: strings.TrimSpace(r.URL.Query().Get("classNum"))}
	
	srv := load_balancer.Servers[0]
	
	srv.Lock()
	defer srv.Unlock()
	
	fmt.Println(len(srv.Classes))


	// for key := range srv.Classes {
	// 	fmt.Printf("key %s", key)
	// }
	fmt.Printf("classNum: %s\n", eInfo.ClassNum)
	class, found := srv.Classes[eInfo.ClassNum]
	for {
		if !found || class.Enrollment == class.MaxEnrollment {
			msg := "you cannot enroll in that class"
			fmt.Println(msg)
			w.(http.Flusher).Flush()
			fmt.Println("-------------------")
		return
		}
	}
	


	// class.Enrollment++
	// fmt.Printf("new enrollment %d\n", class.Enrollment)
	// success := database.UpdateDB(load_balancer, eInfo.StudentID, eInfo.ClassNum, 0)
	// if success {
	// 	fmt.Println("successful update of the database")
	// } else {
	// 	msg := "you were not enrolled in class " + eInfo.ClassNum
	// 	w.Write([]byte(msg))
	// }
	// fmt.Println(eInfo.StudentID, eInfo.ClassNum)

	// fmt.Println("Running on Port :7777")
	// message := "byE BYE BYE"
	// w.Write([]byte(message))
	fmt.Println("Running on Port :7777")
	message := "byE BYE BYE"
	w.Write([]byte(message))
	fmt.Println("-------------------")
}

func handlerServer2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server1 -----------")
	// eInfo := processQuery(r)
	eInfo := &EnrollmentInfo{StudentID: r.URL.Query().Get("studentID"), ClassNum: strings.TrimSpace(r.URL.Query().Get("classNum"))}
	srv := load_balancer.Servers[1]
	srv.Lock()
	defer srv.Unlock()
	fmt.Printf("classNum: %s\n", eInfo.ClassNum)
	class, found := srv.Classes[eInfo.ClassNum]
	if !found || class.Enrollment == class.MaxEnrollment {
		msg := "you cannot enroll in that class"
		fmt.Println(msg)
		w.Write([]byte(msg))
		fmt.Println("-------------------")
		return
	}
	class.Enrollment++
	fmt.Printf("new enrollment %d\n", class.Enrollment)
	success := database.UpdateDB(load_balancer, eInfo.StudentID, eInfo.ClassNum, 1)
	if success {
		fmt.Println("successful update of the database")
	} else {
		msg := "you were not enrolled in class " + eInfo.ClassNum
		w.Write([]byte(msg))
	}
	fmt.Println(eInfo.StudentID, eInfo.ClassNum)

	fmt.Println("Running on Port :8888")
	message := "byE BYE BYE"
	w.Write([]byte(message))
	fmt.Println("-------------------")
}

func handlerServer3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server2 -----------")
	// eInfo := processQuery(r)
	eInfo := &EnrollmentInfo{StudentID: r.URL.Query().Get("studentID"), ClassNum: strings.TrimSpace(r.URL.Query().Get("classNum"))}
	srv := load_balancer.Servers[2]
	srv.Lock()
	defer srv.Unlock()
	fmt.Printf("classNum: %s\n", eInfo.ClassNum)
	class, found := srv.Classes[eInfo.ClassNum]
	if !found || class.Enrollment == class.MaxEnrollment {
		msg := "you cannot enroll in that class"
		fmt.Println(msg)
		w.Write([]byte(msg))
		fmt.Println("-------------------")
		return
	}
	class.Enrollment++
	fmt.Printf("new enrollment %d\n", class.Enrollment)
	success := database.UpdateDB(load_balancer, eInfo.StudentID, eInfo.ClassNum, 2)
	if success {
		fmt.Println("successful update of the database")
	} else {
		msg := "you were not enrolled in class " + eInfo.ClassNum
		w.Write([]byte(msg))
	}
	fmt.Println(eInfo.StudentID, eInfo.ClassNum)

	fmt.Println("Running on Port :9999")
	message := "byE BYE BYE"
	w.Write([]byte(message))
	fmt.Println("-------------------")

}