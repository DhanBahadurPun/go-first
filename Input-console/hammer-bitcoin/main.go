package main

import (
	"fmt"
	"game/game"
)

func main() {
	playAgain := true
	for playAgain {
		game.Play()
		playAgain = game.GetYesOrNo("Woudl you like to play again (y/n)?")
	}

	fmt.Println("")
	fmt.Println("Goodbye")
}
