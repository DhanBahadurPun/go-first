package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

// func main() {
// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		fmt.Println("-> ")

// 		userInput, _ := reader.ReadString('\n')

// 		userInput = strings.Replace(userInput, "\n", "", -1)

// 		if userInput == "quit" {
// 			break
// 		} else {
// 			fmt.Println(userInput)
// 		}
// 	}
// }

// func main() {
// 	err := keyboard.Open()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer func() {
// 		_ = keyboard.Close()
// 	}()

// 	fmt.Println("Press any key on the keyboard. Press ESC to quit.")

// 	for {
// 		char, key, err := keyboard.GetSingleKey()

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		if key != 0 {
// 			fmt.Println("You pressed!", char, key)
// 		} else {
// 			fmt.Println("You pressed!", char, key)
// 		}

// 		if key == keyboard.KeyEsc {
// 			fmt.Println("Program exited>>>>")
// 			break
// 		}

// 	}
// }

func main() {
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)

	coffees[1] = "cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("------")
	fmt.Println("1 -- Cappucino")
	fmt.Println("2 -- Latte")
	fmt.Println("3 -- Americano")
	fmt.Println("4 -- Mocha")
	fmt.Println("5 -- Macchiato")
	fmt.Println("6 -- Espresso")
	fmt.Println("7 -- Quit the program")

	for {
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {

			break
		}

		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))

	}

	fmt.Println("Program exited......")
}
