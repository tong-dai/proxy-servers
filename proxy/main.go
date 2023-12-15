package main

import (
	"fmt"
	"net/http"
)

type EnrollmentInfo struct {
	StudentID string
	ClassNum  string
}

//	func processQuery(r *http.Request) bool {
//		studentID, err := strconv.Atoi(r.URL.Query().Get("studentID"))
//		if err != nil {
//			fmt.Println("Something went wrong converting studentID to int")
//		}
//		classNum, err := strconv.Atoi(r.URL.Query().Get("classNum"))
//		if err != nil {
//			fmt.Println("Something went wrong converting classNum to id")
//		}
//		// success := db.UpdateDB(studentID, classNum, server.Load_balancer.Servers, server.Load_balancer.GetNextServerIndex())
//		fmt.Printf("studentID %v", studentID)
//		fmt.Printf("classNum %v", classNum)
//		return true
//		// return &EnrollmentInfo{StudentID: studentID, ClassNum: classNum}
//	}
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
		http.ListenAndServe(":9999", server3)
	}()

	//Running Main Server
	http.ListenAndServe("localhost:9000", main_server)
}

func handlerServer1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hellor???")
	enrollmentInfo := processQuery(r)
	// if success {
	//     for key, value := range db.DB.C {

	//     }
	// }
	fmt.Println(enrollmentInfo.StudentID, enrollmentInfo.ClassNum)
	
	fmt.Println("Running on Port :7777")
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
