package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	beforeTime := time.Now()
	for scanner.Scan() {
		requestURL := scanner.Text()[4:]
		// if err != nil {
		// 	fmt.Println("someting went wrong parsing the strings")
		// 	break
		// }
		resp, err := http.Get(requestURL)
		if err != nil {
			fmt.Println("Error", err)
			break
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

			fmt.Println("Received event: ", string(line))
		}
	}
	afterTime := time.Now()
	executionTime := afterTime.Sub(beforeTime)
	fmt.Printf("executionTime: %v\n", executionTime)

}
