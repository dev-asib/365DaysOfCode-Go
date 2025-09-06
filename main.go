package main

import "fmt"

func main() {
	printWelcomeMessage()
	name := getUserName()
	num1, num2 := getTwoNumbers()
	sum := add(num1, num2)
	display(name, sum)
	printGoodbyeMessage()
}

func printWelcomeMessage() {
	fmt.Println("Welcome to our applicatoin")
}

func getUserName() string {
	var userName string
	fmt.Println("Enter your name - ")
	fmt.Scanln(&userName)
	return userName
}

func getTwoNumbers() (int, int) {
	var numb1 int
	var numb2 int

	fmt.Println("Enter num 1 - ")
	fmt.Scanln(&numb1)

	fmt.Println("Enter numb 2 - ")
	fmt.Scanln(&numb2)

	return numb1, numb2
}

func add(numb1 int, numb2 int) int {
	sum := numb1 + numb2

	return sum
}

func display(userName string, sum int) {
	fmt.Println("Hi, ", userName)
	fmt.Println("Summation = ", sum)
}

func printGoodbyeMessage() {
	fmt.Println("Thanks for using my application")
	fmt.Println("Good Bye")
}
