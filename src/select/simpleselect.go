package main

import "fmt"

func main() {
	m1 := make(chan string)
	m2 := make(chan string)

	go func() { m1 <- "first" }()
	go func() { m2 <- "second" }()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-m1:
			fmt.Println(msg1)
		case msg2 := <-m2:
			fmt.Println(msg2)

		}
	}

}
