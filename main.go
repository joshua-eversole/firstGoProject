package main

import (
	"fmt"
)

// you have to put the type returned in the function decleration

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
	symbolArr := GenerateSymbolArray(symbols)
	//we're starting the user with 200 dollars
	balance := uint(200)
	//GetName()
	bet := uint(0)
	for balance > 0 {
		bet = GetBet(balance)
		if bet == 0 { // Corrected this line
			break
		}
		balance -= bet

		spin := GetSpin(symbolArr, 3, 3)
		PrintSpin(spin)
		winningLines := checkWin(spin, multipliers)
		//this is the winning condition
		fmt.Printf("winning lines: %v\n", winningLines)
		for i, multi := range winningLines {
			win := bet * uint(multi)
			balance += win
			if multi > 0 {
				fmt.Printf("You won $%d on line %d!\n", win, i+1)
			}
		}
		//if bet == 0 || !continueGame() {
		if bet == 0 {
			break
		}

	}

	fmt.Printf("Thanks for playing! You left with $%d\n", balance)

}
