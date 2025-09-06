package main

import "fmt"

func main() {

	// First Program
	fmt.Println("Hi! Go")

	// Variable Delclaration
	var name = "Asib"
	fmt.Println(name)

	var age int = 25
	fmt.Println(age)

	country := "Bangladesh"
	fmt.Println(country)

	// Data Types
	var mobileBrand string = "Tecno"
	fmt.Println(mobileBrand)

	var price int = 2500
	fmt.Println(price)

	var height float32 = 6.7
	fmt.Println(height)

	var isTecno = true
	fmt.Println(isTecno)

	taka := 200

	if taka == 20 {
		fmt.Println("You will get 2 Potatos")
	} else if taka == 200 {
		fmt.Println("You will get 1 Potatos")
	} else if taka <= 0 {
		fmt.Print("You have no money")
	} else {
		fmt.Println("You have balance")
	}

	userName := "Tamim"

	switch userName {
	case "Asib":
		fmt.Print("You are a right user.")
	case "Samy", "Sabbir":
		fmt.Println("You are a Asib's Bhai")
	default:
		fmt.Println("You are not a valid person")
	}

	a := 2
	b := 5

	c := 10
	d := 20

	addTwoNumbers(a, b)
	addTwoNumbers(10, 20)
	goDeveloperName()
	sum := addNumb(a, b)
	fmt.Print(sum)
	mul, info := multipleReturn(c, d)

	fmt.Println(mul)
	fmt.Println(info)
	welcomeCourse("Asib")
	friends("Samy, ", "Sabbir")

	/// User input

	var myName string

	fmt.Scanln(&myName)

	fmt.Println("My name is - ", myName)

	var myAge int = 24
	fmt.Println(myAge)
	fmt.Scanln(&myAge)
	fmt.Println("My Age is ", myAge)

}

func addTwoNumbers(numb1 int, numb2 int) {
	sum := numb1 + numb2

	fmt.Println(sum)
}

func goDeveloperName() {
	fmt.Println("Asib")
}

/// Single Return Type Function

func addNumb(a int, b int) int {
	return a + b
}

/// Multiple Return Type Function

func multipleReturn(n1 int, n2 int) (int, string) {
	return n1 * n2, "Asib"
}

func welcomeCourse(name string) {
	fmt.Println("Welcome to the Golang course", name, "Bhai")
}

func friends(name1 string, name2 string) {
	fmt.Print("Our friends: ", name1, name2)
}
