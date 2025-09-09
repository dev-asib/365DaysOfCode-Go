package main

import "fmt"

var a int = 20

func main() {
	fmt.Println("Hello Init Function")
	fmt.Println(a)
}

func init() {
	fmt.Println("This is init function")
	fmt.Println(a)
	a = 30
}
