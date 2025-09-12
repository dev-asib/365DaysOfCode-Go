package main

import "fmt"

var sum = func(a, b int) {
	fmt.Println(a + b)
}

func main() {

	sum(11, 22)

	add := func(a int, b int) {
		fmt.Println(a + b)
	}

	add(10, 20)

	sum(33, 44)

}

func init() {
	sum(100, 200)
}
