package main

// import (
// 	"fmt"
// 	"sync"
// )

// type ClassInfo struct {
// 	maxEnrollment int
// 	enrollment    int
// 	enrolledIds   []int
// }

// type DB struct {
// 	c map[int]*ClassInfo
// 	sync.Mutex
// }
// var DB map[int]*ClassInfo = make(map[int]*ClassInfo)

// func (db *DB) UpdateDB(studentID int, classNumber int, cacheIndex int) (bool, int) {
// 	db.Lock()
// 	defer db.Unlock()
// 	class := db.c[classNumber]
// 	if class.enrollment == class.maxEnrollment {
// 		// TODO send message to all caches saying that you need to update this
// 		fmt.Println("something went wrong")
// 		return false, classNumber
// 	} else {
// 		class.enrolledIds = append(class.enrolledIds, studentID)
// 		return true, -1
// 	}
	
// }
