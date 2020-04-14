package main

import "fmt"

func soma(num1 int, num2 int) int{
	return num1 + num2
}

func main(){
	
	num1:= 5;
	num2:= 5;
	
	result:= soma(num1, num2)

	fmt.Println("A soma entre os dois números é:", result)
}