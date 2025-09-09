// package main

// import (
// 	"mathlib"
// )

// func main() {
// 	mathlib.AddNumbers()
// 	mathlib.AddTwoNumbers(5, 10)
// }

package main

import "fmt"

var a int = 10
var b int = 20

func main() {
	addNumbers(a, b)

}

func addNumbers(x, y int) {
	result := x + y
	printNum(result)
}

func printNum(num int) {
	fmt.Println(num)
}
