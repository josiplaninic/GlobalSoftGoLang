package main

import "fmt"

// 1. Declare two integers, `firstNumber` and `secondNumber`, assign values 20 and 30 to them.
// Swap values of `firstNumber` and `secondNumber` without using the third variable
// After all, print values of `firstNumber` and `secondNumber`.


// 2. Declare two variables, `firstName` and `lastName` assign them with wanted values.
// Declare constant named `fullname`
// Combine constant and both strings into a full name by concatenating strings with a space in between and print them out.

func main() {

	//first task

	var firstNumber int = 20
	var secondNumber int = 30

	fmt.Printf("Before the swap %d %d \n", firstNumber, secondNumber)
	
	
	firstNumber += secondNumber
	secondNumber = firstNumber - secondNumber
	firstNumber -= secondNumber 
	
	fmt.Printf("After the swap %d %d \n", firstNumber, secondNumber)
	
	
	//second task
	
	var firstName string="Marko"
	var lastName string="Markovic"
	const FULLNAME string = "Full name:"
	var fullname string = FULLNAME + " "+ firstName + " " + lastName

	fmt.Println(fullname)

	


}