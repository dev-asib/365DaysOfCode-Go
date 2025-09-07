package main

import "fmt"

var a = 10
var b = 20

func add(x int, y int) {
	sum := x + y
	fmt.Println(sum)
}

func addTwoNumbersx(x int, y int) {
	sum := a + b
	fmt.Println(sum)
}

func main() {

	fmt.Println("Hi, Go")

	c := 20
	d := 30

	add(c, d)

	addTwoNumbersx(c, d)

	var name = "Asib"

	fmt.Println(name)

	m := b
	fmt.Println(m)
}
