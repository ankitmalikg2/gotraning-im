package main

import "fmt"

func main() {
	fmt.Println("Hello ankit malik")

	a := [5]int{1, 2, 5, 6, 8}

	fmt.Println(a)

	fmt.Println(a[:len(a)-1])
	fmt.Println(a[:2])
	fmt.Println(a[2:])

	for _, num := range a {
		fmt.Print(num, " ")
	}
}
