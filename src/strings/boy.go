package main

import s "strings"
import "fmt"

func main() {
	fmt.Println("Hello ankit malik")
	fmt.Println("Contains:  ", s.Contains("test", "es"))
	a := "boy"
	fmt.Println(reverse(a, 0, len(a)-1))

	b := "i am a boy"
	c := reverse(b, 0, len(b)-1)
	d := ReverseWords(c)
	fmt.Println(c)
	fmt.Println(d)
	// var d []int

	// for(int i=0;i<n)

}

func reverse(s string, start int, end int) string {
	chars := []rune(s)
	fmt.Println(chars)
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func ReverseWords(s string) string {
	str := ""
	chars := []rune(s)
	last := 0
	for i := 0; i < (len(chars) - 2); i++ {
		// fmt.Println(i)
		if chars[i] == 32 {
			fmt.Println(i)
			str += reverse(s, last, i)
			last = i + 1
		}
	}
	return str
}
