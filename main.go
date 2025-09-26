package main

import "fmt"

func main(){
processOperation(5, 7, add)
sum := callAdd()
sum(10, 20)

}


func processOperation(a int, b int, op func (c int, d int)){
op(a, b)
}

func add(x int, y int){
	z := x+y
	fmt.Println(z)
}


func callAdd() func ( int,  int){
	return add
}

/*
Higher Order Function
01. Parameter -> Function
02. Function -> Return
03. Both
*/