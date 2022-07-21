package main

import (
	"fmt"
	"myapp/packageone"
)

// package level variable is avialable in every
var one = "One"

func main() {
	// block level variable have higher precedence! if same name for both
	var one = "the number one!"
	fmt.Println(one)
	myFunc()

	newString := packageone.PublicVar

	fmt.Println("From packageone", newString)

}

func myFunc() {

	fmt.Println(one)
}

// go mod init myapp
