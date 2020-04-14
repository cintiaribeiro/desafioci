package main

import "fmt"

func soma(n1 int, n2 int) int{
	return n1 + n2
}

func main(){
	
	num1:= 5;
	num2:= 5;
	
	result:= soma(num1, num2)

	fmt.Println("A soma entre os dois números é:", result)
}