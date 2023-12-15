package main

import (
	"bufio"
	"fmt"
	"io"
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

	resp, err := http.Get(urlWithParams)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	respReader := bufio.NewReader(resp.Body)
    for {
        line, err := respReader.ReadBytes('\n')
		
        if err != nil {
			if err == io.EOF {
				break
			}
            fmt.Println("Error reading line:", err)
            break
        }

        fmt.Print("Received event: ", string(line))
    }

}
