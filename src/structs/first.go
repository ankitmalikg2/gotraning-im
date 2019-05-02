package main

import "fmt"

type product struct {
	name   string
	price  int
	height int
	weight int
}

func (p *product) area() int {
	return p.height * p.weight
}

func main() {
	fmt.Println("Hello ankit malik")

	p := product{name: "ankit", price: 10, height: 12, weight: 21}
	fmt.Println(p.name)

	p.price--
	fmt.Println(p)

	fmt.Println(p.area())

	fmt.Println(product{name: "ram"})
}
