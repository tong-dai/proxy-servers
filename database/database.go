package db

import (
	lb "cos316/td_ec_final_project/load_balancer"
	"fmt"
	"sync"
	"time"
)

type ClassInfo struct {
	maxEnrollment int
	enrollment    int
	enrolledIds   []string
}

// type CacheClassInfo struct {
// 	maxEnrollment int
// 	enrollment    int
// }

//	type Server struct {
//		serverURL string
//		index     int
//		sync.Mutex
//		classes map[int]*CacheClassInfo
//	}
type DB struct {
	C map[string]*ClassInfo
	sync.Mutex
}

var database *DB = &DB{C: createDB(1000)}

func createDB(numClasses int) map[string]*ClassInfo {
	classInfo := make(map[string]*ClassInfo)
	for i := 0; i < numClasses; i++ {
		classInfo[fmt.Sprint(i)] = &ClassInfo{enrollment: 0, maxEnrollment: 15, enrolledIds: make([]string, 0)}
	}
	return classInfo
}

func GetDB() *DB {
	return database
}

// var DB map[int]*ClassInfo = make(map[int]*ClassInfo)

// func DeleteClass(servers []*lb.Server, classNumber int) {
// 	for i := 0; i < len(servers); i++ {
// 		servers[i].Lock()
// 		_, found := servers[i].Classes[classNumber]
// 		if found {
// 			delete(servers[i].Classes, classNumber)
// 		}
// 		servers[i].Unlock()
// 	}
// }

// send the new enrollemnt to all servers except for the one that sent the last request
// func UpdateServers(servers []*lb.Server, classNumber int, enrollment int, serverIndex int) {
// 	for i := 0; i < len(servers); i++ {
// 		if i != serverIndex {
// 			servers[i].Lock()
// 			class, found := servers[i].Classes[classNumber]
// 			if !found {
// 				continue
// 			}
// 			if enrollment == class.MaxEnrollment {
// 				delete(servers[i].Classes, classNumber)
// 			} else {
// 				servers[i].Classes[classNumber].Enrollment = enrollment
// 			}
// 			servers[i].Unlock()
// 		}
// 	}
// }

func (db *DB) UpdateDB(load_balancer *lb.LB, studentID string, classNumber string, serverIndex int) bool {
	db.Lock()
	defer db.Unlock()
	// delay to simulate accessing a database
	time.Sleep(15 * time.Millisecond)
	class := db.C[classNumber]

	if class.enrollment >= class.maxEnrollment {
		load_balancer.DeleteClass(classNumber)
		// fmt.Printf("something went wrong enrolling student %v in class %v\n", studentID, classNumber)
		return false
	} else {
		//TODO maybe add a delay here to imitate accessing an actual database?
		time.Sleep(22 * time.Millisecond)
		class.enrolledIds = append(class.enrolledIds, studentID)
		class.enrollment++
		// fmt.Printf("length of class.enrolledIds %d\n", len(class.enrolledIds))
		// TODO figure out when to call this method => how often?
		// defer load_balancer.UpdateServer(classNumber, class.enrollment, serverIndex)
		return true
	}

}
func (db *DB) UpdateDBNoCache(studentID string, classNumber string) bool {
	db.Lock()
	defer db.Unlock()
	// delay to simulate accessing a database
	time.Sleep(15 * time.Millisecond)
	class := db.C[classNumber]

	if class.enrollment >= class.maxEnrollment {
		// fmt.Printf("something went wrong enrolling student %v in class %v\n", studentID, classNumber)
		return false
	} else {
		//TODO maybe add a delay here to imitate accessing an actual database?
		time.Sleep(22 * time.Millisecond)
		class.enrolledIds = append(class.enrolledIds, studentID)
		class.enrollment++
		// fmt.Printf("length of class.enrolledIds %d\n", len(class.enrolledIds))
		// TODO figure out when to call this method => how often?
		// defer load_balancer.UpdateServer(classNumber, class.enrollment, serverIndex)
		return true
	}

}
