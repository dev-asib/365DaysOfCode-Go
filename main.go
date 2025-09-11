package main

import "fmt"

func main() {

	func(a, b int) {
		result := a + b
		fmt.Println(result)

	}(20, 30)

}

func init() {
	fmt.Println("I will be called first")

}
