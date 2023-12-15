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

	// go func() {
	// 	fmt.Println("Server started on: http://localhost:8888")
	// 	http.ListenAndServe(":8888", server2)
	// }()

	// go func() {
	// 	fmt.Println("Server started on: http://localhost:9999")
	// 	http.ListenAndServe(":9999", server3)
	// }()

	// Running Main Server
	http.ListenAndServe("localhost:9000", main_server)
}

func handlerServer1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/event-stream")

	eInfo := &EnrollmentInfo{StudentID: r.URL.Query().Get("studentID"), ClassNum: strings.TrimSpace(r.URL.Query().Get("classNum"))}
	
	srv := load_balancer.Servers[0]
	
	srv.Lock()
	defer srv.Unlock()

	var message string
	class, found := srv.Classes[eInfo.ClassNum]
	for {
		if !found {
			message = "You failed to enroll in class " + eInfo.ClassNum
			fmt.Fprintln(w, message)
			w.(http.Flusher).Flush()
			return
		} 
		// write to about successfully enrollment
		class.Enrollment++
		success := database.UpdateDB(load_balancer, eInfo.StudentID, eInfo.ClassNum, 0)
		if success {
			fmt.Println("successful update of the database")
		} else {
			message = "Class enrollment revoked for class number " + eInfo.ClassNum
			fmt.Fprintln(w, message)
			w.(http.Flusher).Flush()
		}
	}
}

func handlerServer2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hellor???")

	enrollmentInfo := processQuery(r)
	// if success {
	//     for key, value := range db.DB.C {

	//     }
	// }
	fmt.Println(enrollmentInfo.StudentID, enrollmentInfo.ClassNum)
	fmt.Println("Running on Port :8888")
	message := "byE BYE BYE"
	w.Write([]byte(message))
}

func handlerServer3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hellor???")

	enrollmentInfo := processQuery(r)
	// if success {
	//     for key, value := range db.DB.C {

	//     }
	// }

	fmt.Println(enrollmentInfo.StudentID, enrollmentInfo.ClassNum)
	fmt.Println("Running on Port :9999")

}

// func Enroll(servers []*Server, i int, w http.ResponseWriter, r *http.Request) {
// 	studentID, err := strconv.Atoi(r.URL.Query().Get("studentID"))
// 	if err != nil {
// 		log.Panic("something went wrong converting studentID to an int")
// 	}
// 	classNumber, err := strconv.Atoi(r.URL.Query().Get("classNumber"))
// 	if err != nil {
// 		log.Panic("something went converting classNumber to an int")
// 	}
// 	servers[i].Lock()
// 	defer servers[i].Unlock()
// 	class, found := servers[i].classes[classNumber]
// 	if found {
// 		fmt.Println("found the class")
// 		class.enrollment++
// 		success := db.UpdateDB(studentID, classNumber, servers, i)
// 		if success {
// 			fmt.Fprintf(w, "Successfully enrolled in %v", classNumber)
// 		} else {
// 			fmt.Fprintf(w, "Sorry, you were not enrolled in class %v", classNumber)
// 		}
// 	} else {
// 		fmt.Fprintf(w, "Sorry, you were not actually enrolled in class %v", classNumber)
// 	}
// }
