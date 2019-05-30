package main

func main() {
	str := "ankit"

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)

}

func video(str string, c chan string) {
	c <- str
}

func image(str string, c chan string) {
	c <- str
}

func search(str string, c chan string) {
	c <- str
}

func news(str string, c chan string) {
	c <- str
}
