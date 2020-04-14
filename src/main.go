package main

import "fmt"

func soma(a int, b int) int{
	return a + b
}

func main(){
	
	num1:= 5;
	num2:= 5;
	
	result:= soma(num1, num2)

	fmt.Println("A soma entre os dois números é:", result)
}