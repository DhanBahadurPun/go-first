package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready."

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// one way - declare, then assign (two steps)
	var firstNumber int
	firstNumber = rand.Intn(8) + 2

	// another way, declare type and name and assign value
	var secondNumber = rand.Intn(8) + 2

	// one step variable: declare name, assign value and let GO figure out the type of variable
	substraction := rand.Intn(8) + 2

	var answer = firstNumber*secondNumber - substraction

	PlayTheGames(firstNumber, secondNumber, substraction, answer)

}

func PlayTheGames(firstNumber int, secondNumber int, substraction int, answer int) {
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
	fmt.Println("The Answer: ", answer)

}
