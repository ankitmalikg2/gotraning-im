// time.Sleep(100 * time.Millisecond)
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	// ch1 := make(chan string)
	// ch2 := make(chan string)
	// ch3 := make(chan string)
	// ch4 := make(chan string)

	counter := make(chan int, 1)
	counter <- 1

	for i := 0; i < 10; i++ {
		distribute_data("ankit"+strconv.Itoa(i), counter)
	}

}

func distribute_data(s string, counter chan int) {
	c := <-counter
	fmt.Println(c)
	if c == 1 {
		go lb1(s, counter)
	} else if c == 2 {
		go lb2(s, counter)
	} else if c == 3 {
		go lb3(s, counter)
	} else if c == 4 {
		go lb4(s, counter)
	}
}

func lb1(s string, counter chan int) {
	fmt.Println(s, "one")
	time.Sleep(1000 * time.Millisecond)
	counter <- 2
}

func lb2(s string, counter chan int) {
	fmt.Println(s, "two")
	time.Sleep(2000 * time.Millisecond)
	counter <- 3
}

func lb3(s string, counter chan int) {
	fmt.Println(s, "three")
	time.Sleep(3000 * time.Millisecond)
	counter <- 4
}

func lb4(s string, counter chan int) {
	fmt.Println(s, "four")
	time.Sleep(4000 * time.Millisecond)
	counter <- 1
}
