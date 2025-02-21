package main

import (
	"fmt"
	"math/rand"
)

// you have to put the type returned in the function decleration
func getName() string {
	name := ""
	fmt.Println("Welcome to Josh's Casino!")
	fmt.Printf("Enter your name: ")
	//the & gives the memory address of the variable
	// this means the Scanln function can store the input in the variable and replace the existing value
	_, err := fmt.Scanln(&name)
	// nil is none, err
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Printf("Hello, %s! Let's play some games!\n", name)
	return name
}

func getBet(balance uint) uint {
	bid := 0
	//There is no while loop in go (wtf i'm gonna cry i love while loops)
	for true {
		fmt.Printf("You have $%d. How much would you like to bet? ", balance)
		_, err := fmt.Scanln(&bid)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		if bid <= int(balance) {
			break
		}
		fmt.Println("You don't have enough money to bet that much!")
	}
	return uint(bid)
}

func continueGame() bool {
	choice := ""
	for true {
		fmt.Printf("Would you like to continue playing? (y/n) ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println(err)
			return false
		}
		if choice == "y" {
			return true
		} else if choice == "n" {
			return false
		} else {
			fmt.Println("Please enter y or n.")
		}
	}
	return false
}

// slice is an image of an array, it's a reference to an array and can be resized (which arrays can't be)
func generateSymbolArray(symbols map[string]uint) []string {
	symbolArray := []string{}
	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArray = append(symbolArray, symbol)
		}
	}
	return symbolArray
}

func getRandomNumber(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func getSpin(reel []string, rows int, cols int) [][]string {
	spin := [][]string{}
	for i := 0; i < rows; i++ {
		spin = append(spin, []string{})
	}
	//this is a nested loop, the start with cols instead of rows on purpose
	for col := 0; col < cols; col++ {
		selected := map[int]bool{}
		for row := 0; row < rows; row++ {
			for true {
				randomIndex := getRandomNumber(0, len(reel)-1)
				if _, exists := selected[randomIndex]; !exists {
					selected[randomIndex] = true
					spin[row] = append(spin[row], reel[randomIndex])
					break
				}
			}
		}
	}
	return spin
}

func printSpin(spin [][]string) {
	for _, row := range spin {
		for j, symbol := range row {
			fmt.Printf("%s ", symbol)
			if j == len(row)-1 {
				fmt.Printf(" | ")
			}
		}
		fmt.Println()
	}
}

func checkWin(spin [][]string, multipliers map[string]uint8) []int {
	wins := []int{}
	for _, row := range spin {
		win := true
		checkSymbol := row[0]
		for _, symbol := range row[1:] {
			if symbol != checkSymbol {
				win = false
				break
			}
		}
		if win {
			wins = append(wins, int(multipliers[checkSymbol]))
		} else {
			wins = append(wins, 0)
		}
	}
	return wins
}

func main() {
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}
	multipliers := map[string]uint8{
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}
	symbolArr := generateSymbolArray(symbols)
	//we're starting the user with 200 dollars
	balance := uint(200)
	getName()
	bet := uint(0)
	for balance > 0 {
		bet = getBet(balance)
		if bet == 0 { // Corrected this line
			break
		}
		balance -= bet

		spin := getSpin(symbolArr, 3, 3)
		printSpin(spin)
		winningLines := checkWin(spin, multipliers)
		//this is the winning condition
		for i, multi := range winningLines {
			win := bet * uint(multi)
			balance += win
			if multi > 0 {
				fmt.Printf("You won $%d on line %d!\n", win, i+1)
			}
		}
		if bet == 0 || !continueGame() {
			break
		}

	}

	fmt.Printf("Thanks for playing! You left with $%d\n", balance)

}
