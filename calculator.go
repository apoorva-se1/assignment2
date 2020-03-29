package main

import "fmt"

func main() {
	var number1, number2 int
	fmt.Print("Please enter First number: ")
	fmt.Scanln(&number1)
	fmt.Print("Please enter Second number: ")
	fmt.Scanln(&number2)

	var mulRes = mul(number1, number2)
	var addRes = add(number1, number2)
	var subRes = sub(number1, number2)
	var divRes = div(number1, number2)

	defer output(mulRes, addRes, subRes, divRes)
}

func mul(a1, a2 int) int {
	res := a1 * a2
	return res
}

func add(a1, a2 int) int {
	res := a1 + a2
	return res
}

func sub(a1, a2 int) int {
	res := a1 - a2
	return res
}
func div(a1, a2 int) int {
	res := a1 / a2
	return res
}

func output(a1, a2, a3, a4 int) {
	fmt.Println("Result of multiplication: ", a1)
	fmt.Println("Result of addition: ", a2)
	fmt.Println("Result of subtraction: ", a3)
	fmt.Println("Result of division: ", a4)
}
