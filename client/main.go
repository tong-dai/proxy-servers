package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your nine digit PUID: ")
	studentID, _ := reader.ReadString('\n')
	studentID = studentID[:len(studentID)-1]

	fmt.Print("Enter the five digit class number: ")
	classNum, _ := reader.ReadString('\n')
	classNum = classNum[:len(classNum)-1]

	baseUrl := "http://localhost:8080/"

	params := url.Values{}
	params.Add("studentID", studentID)
	params.Add("classNum", classNum)

	urlWithParams := baseUrl + "?" + params.Encode()

	req, err := http.Get(urlWithParams)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer req.Body.Close()
}
