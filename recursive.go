package main

import "fmt"

func main() {
	var a string = "This is a recursion call function"
	output(a)

}

func output(a string) {
	for i := range a {
		fmt.Printf("Character in string in hexadecimal: %x\n ", a[i])
	}
}
