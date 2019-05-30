package main

import "fmt"
import "time"

func main() {
	requests := make(chan int, 5)

	for i := 0; i < 4; i++ {
		requests <- (i + 22)
	}
	close(requests)

	limiter := time.Tick(900 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println(req, time.Now())
	}

}
