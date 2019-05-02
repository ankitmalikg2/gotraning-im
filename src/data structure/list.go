package main

import "fmt"

func main() {
	fmt.Println("Hello ankit malik")

	a := []int{1, 2, 5, 6, 8, 11, 10}

	for i, num := range a {
		fmt.Print(i, num, " -- ")
	}
	fmt.Println()

	for i := (len(a) - 1); i >= 0; i-- {
		fmt.Print(i, a[i], "-")
	}
	fmt.Println()

}
