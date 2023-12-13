package main

import (
	// "net/http"
	"bufio"
	"fmt"
	"net/http"
	"os"
)

// type EnrollmentRequest struct {
// 	StudentID	int			// student's 9-digit id
// 	ClassNumber	int			// class number
// }

// func sendEnrollRequest(stuentID []byte, classNumber []byte) bool {

// }

// Validate student ID and class number input
func isValidInput(input []byte, expectedLength int) bool {
	if len(input) > 0 {
		input = input[:len(input)-1]
	}

	if len(input) != expectedLength {
		return false
	}

	for _, b := range input {
		if b > 57 || b < 48 {
			return false
		}
	}
	return true
}

// Runs the client
func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Student ID: ")
	studentID, _ := reader.ReadBytes('\n')
	fmt.Println(studentID)
	// if (!isValidInput(studentID, 9)) {
	// 	fmt.Println("Please enter your 9-digit student ID correctly.")
	// 	return
	// }

	fmt.Println("Enter Class Number: ")
	classNumber, _ := reader.ReadBytes('\n')
	fmt.Println(classNumber)
	response, err := http.Get("http://localhost:8080/enroll")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)

	// if (!isValidInput(classNumber, 5)) {
	// 	fmt.Println("Please enter the 5-digit class number correctly.")
	// 	return
	// }

}
