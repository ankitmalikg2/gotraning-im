package main

import "fmt"

func fish(c chan<- string) {
	c <- "fish"

}

func chicken(c chan string) {
	c <- "chicken"
}

func egg(c chan string) {
	c <- "egg"
}

func main() {
	c := make(chan string, 1)
	d := make(chan string)
	e := make(chan string)
	go fish(c)
	go chicken(d)
	go chicken(c)
	go egg(e)

	fmt.Println(<-c)
	fmt.Println(<-d)
	fmt.Println(<-e)

	fmt.Println(<-c)
}
