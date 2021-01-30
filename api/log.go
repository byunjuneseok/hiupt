package api

import "fmt"

func Logger(body string) {
	fmt.Println("Received: ", body)
}
