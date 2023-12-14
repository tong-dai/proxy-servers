package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type ClassInfo struct {
	maxEnrollment int
	enrollment    int
	enrolledIds   []int
}

type CacheClassInfo struct {
	maxEnrollment int
	enrollment    int
}

type DB struct {
	c map[int]*ClassInfo
	sync.Mutex
}
type Server struct {
	serverURL string
	index     int
	sync.Mutex
	classes map[int]*CacheClassInfo
}

var db *DB = &DB{c: make(map[int]*ClassInfo)}

// var DB map[int]*ClassInfo = make(map[int]*ClassInfo)

func DeleteClass(servers []*Server, classNumber int) {
	for i := 0; i < len(servers); i++ {
		servers[i].Lock()
		_, found := servers[i].classes[classNumber]
		if found {
			delete(servers[i].classes, classNumber)
		}
		servers[i].Unlock()
	}
}

// send the new enrollemnt to all servers except for the one that sent the last request
func UpdateServers(servers []*Server, classNumber int, enrollment int, serverIndex int) {
	for i := 0; i < len(servers); i++ {
		if i != serverIndex {
			servers[i].Lock()
			class, found := servers[i].classes[classNumber]
			if !found {
				continue
			}
			if enrollment == class.maxEnrollment {
				delete(servers[i].classes, classNumber)
			} else {
				servers[i].classes[classNumber].enrollment = enrollment
			}
			servers[i].Unlock()
		}
	}
}

func UpdateDB(studentID int, classNumber int, servers []*Server, serverIndex int) bool {
	db.Lock()
	defer db.Unlock()
	class := db.c[classNumber]
	if class.enrollment == class.maxEnrollment {
		defer DeleteClass(servers, classNumber)
		fmt.Printf("something went wrong enrolling student %v in class %v", studentID, classNumber)
		return false
	} else {
		//TODO maybe add a delay here to imitate accessing an actual database?
		class.enrolledIds = append(class.enrolledIds, studentID)
		// TODO figure out when to call this method => how often?
		defer UpdateServers(servers, classNumber, class.enrollment, serverIndex)
		return true
	}

}

func Enroll(servers []*Server, i int, w http.ResponseWriter, r *http.Request) {
	studentID, err := strconv.Atoi(r.URL.Query().Get("studentID"))
	if err != nil {
		log.Panic("something went wrong converting studentID to an int")
	}
	classNumber, err := strconv.Atoi(r.URL.Query().Get("classNumber"))
	if err != nil {
		log.Panic("something went converting classNumber to an int")
	}
	servers[i].Lock()
	defer servers[i].Unlock()
	class, found := servers[i].classes[classNumber]
	if found {
		fmt.Println("found the class")
		class.enrollment++
		success := UpdateDB(studentID, classNumber, servers, i)
		if success {
			fmt.Fprintf(w, "Successfully enrolled in %v", classNumber)
		} else {
			fmt.Fprintf(w, "Sorry, you were not enrolled in class %v", classNumber)
		}
	} else {
		fmt.Fprintf(w, "Sorry, you were not actually enrolled in class %v", classNumber)
	}
}
