package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and press ENTER when ready."

func main() {
	// one way - declare, then assign (two steps)

	var firstNumber int
	firstNumber = 9

	// another way, declare type and name and assign value
	var secondNumber = 7

	// one step variable: declare name, assign value and let GO figure out the type of variable
	substraction := 4

	var answer int

	reader := bufio.NewReader(os.Stdin)

	// display a welcome/instructions

	fmt.Println("Guess the Number Game")
	fmt.Println("-----------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10 ")
	reader.ReadString('\n')

	// take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now substract", substraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	answer = firstNumber*secondNumber - substraction
	fmt.Println("The Answer: ", answer)
}
