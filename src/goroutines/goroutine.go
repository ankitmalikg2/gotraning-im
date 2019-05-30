package main

import "fmt"

func f(a string) {
	for i := 0; i < 10; i++ {

		fmt.Println(i, " -- ", a)
	}
}

func main() {
	g := func(msg string) {
		fmt.Println(msg)
	}
	fmt.Println("Hello ankit malik")
	go f("direct")
	go f("second")

	g("ankit at g function")
	func(msg string) {
		fmt.Println(msg)
	}("panic here and there")

	fmt.Scanln()
	fmt.Println("ankit malik")
}
