package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < 20000; i++ {
		fmt.Fprintf(os.Stdout, "GET http://localhost:8080?studentID=%d&classNum=%d\n", i/4, i%250)
	}
}
