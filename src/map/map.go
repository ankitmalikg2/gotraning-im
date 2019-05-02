package main

import "fmt"
import "reflect"

func main() {
	fmt.Println("Hello ankit malik")
	var testmap = make(map[string]int)
	testmap["first"] = 100

	var test2 = map[string]int{
		"second": 20,
		"value":  100,
	}

	fmt.Println(testmap)
	fmt.Println(test2)
	fmt.Println(reflect.ValueOf(test2).MapKeys)

}
