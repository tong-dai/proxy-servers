package main

import (
	db "cos316/td_ec_final_project/database"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	database := db.GetDB()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		studentID := strings.TrimSpace(r.URL.Query().Get("studentID"))
		classNum := strings.TrimSpace(r.URL.Query().Get("classNum"))
		success := database.UpdateDBNoCache(studentID, classNum)
		if success {
			// fmt.Println("successfully updated database")
			fmt.Fprintf(w, "Successfully enrolled in class%s\n", classNum)
		} else {
			// fmt.Printf("failed to enroll in class %s\n", classNum)
			fmt.Fprintf(w, "failed to enroll in class %s\n", classNum)
		}
	})
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
