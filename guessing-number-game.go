package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	var startGame string
	fmt.Println("Welcome to man versus computer.\n ",
		"I am skudo i will think of a number 5 time and if u guess it right three times you win\n",
		"\t\t\tHINT \n I will be guessing a between 1 to 10\n")
	fmt.Print("Shall we begin? yes / no:   ")
	fmt.Scanf("%s", &startGame)

	var guessed int
	var skudoGuess int
	skudoScores := 0
	playerScore := 0
	capStartGame := strings.ToUpper(startGame)

	for capStartGame == "YES" {
		//fmt.Println(capStartGame)   for testing of the function only. c
		flag := 0
		if flag < 5 {
			skudoGuess = rand.Intn(10) + 1
			fmt.Println(skudoGuess)
			flag++
			fmt.Println("Enter the number you think i guess")
			fmt.Scanf("%d", &guessed)
			if guessed == skudoGuess {
				playerScore++
			} else {
				skudoScores++
			}
			fmt.Printf(" player %d - skudo %d\n", playerScore, skudoScores)
			if playerScore >= 3 {
				fmt.Println("player won")
				capStartGame = "done"
			}

		} else {
			capStartGame = "no"
		}
	}

}
