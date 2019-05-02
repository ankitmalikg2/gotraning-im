package main

import (
	"errors"
	"fmt"
)

type prod interface {
	pricecal() (int, error)
	area() int
}

type product struct {
	name   string
	price  int
	height int
	width  int
	weight int
}

func (p product) area() int {
	return p.height * p.weight
}
func (p product) pricecal() (int, error) {
	if p.price == 0 {
		return 0, errors.New("value is zero")
	}
	return 2 * p.price, nil
}

func measure(p prod) {
	fmt.Println(p)
	fmt.Println(p.area())
	a, err := p.pricecal()
	if err == nil {
		fmt.Println(a)
	} else {
		fmt.Println("error", err.Error())
	}

}

func main() {
	fmt.Println("Hello ankit malik8r7iu")

	p := product{name: "ankit", price: 10, height: 12, width: 15, weight: 21}

	measure(p)

}
