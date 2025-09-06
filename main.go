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

}
