package main

import "fmt"

// func f(a string) string {
// 	if a == "ping" {
// 		return "pong"
// 	}
// 	return a
// }

func main() {
	msg := make(chan string)

	go func() { msg <- "ping" }()

	m := <-msg
	fmt.Println(m)
}
