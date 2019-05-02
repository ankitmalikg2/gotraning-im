package main

import "fmt"

func zerovalue(i int) {
	i = 0
}

func zeroptr(i *int) {
	*i = 0
}

func main() {
	fmt.Println("Hello ankit malik")
	a := 1

	zerovalue(a)
	fmt.Println(a)

	zeroptr(&a)
	fmt.Println(a)
}
