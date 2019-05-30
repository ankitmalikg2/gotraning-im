package main

import "fmt"

type product struct {
	name  string
	price int
}

var products = make(map[int]product)

type behaviour interface {
	create(id int, pr product) int
	update(d int, pr product) int
	remove(id int) int
	getproduct(id int) product
}

func create(id int, pr product) int {
	products[id] = pr
	return 1
}

func update(id int, pr product) int {
	products[id] = pr
	return 1
}

func remove(id int) int {
	delete(products, id)
	return 1
}

// func getproduct(id int) product {
// 	if val, ok := products[id]; ok {
// 		return val
// 	}
// 	return product{}
// }

func main() {
	fmt.Println("Hello ankit malik")
	p1 := product{name: "ankit", price: 20}
	p2 := product{name: "ankit", price: 30}

	create(1, p1)
	create(2, p2)
	fmt.Println(products)

	remove(2)

	fmt.Println(products)

	// fmt.Println(getproduct(1))
	// fmt.Println(getproduct(2))

}
