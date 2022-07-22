package main

import (
	"challenge/packagechallenge"
)

var myVar = "Hi, I am from the main package level!"

func main() {
	// variables
	// declare a package level variable for the main
	// package named myVar

	// declare a block variable for the main function
	// called blockVar

	// declare a package level variable in the packageone
	// package named PackageVar

	// create an exported function in packageone call PrintMe

	// in the main function, print out the values of myVar,
	// blockVar, and PackageVar on one line
	// function in packageone

	var blockVar = "Hi, I am from the main function block level"

	packagechallenge.PrintMe(myVar, blockVar)
}
