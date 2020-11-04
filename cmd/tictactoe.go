package main

import (
	"fmt"
	"os"

	"markus.tictactoe/pkg/tictactoe"
)

func printIndexedBoard() {
	for i := 0; i < 9; i += 3 {
		fmt.Printf("|%d|%d|%d|\n", i, i+1, i+2)
	}
}

func main() {
	cpuOpponent := true
	chosen := false
	var answer string
	for !chosen {
		fmt.Println("Play against CPU (Y/n)?")
		fmt.Scanf("%s", &answer)
		switch answer {
		case "n", "N":
			cpuOpponent = false
			chosen = true
			break
		case "y", "Y", "":
			chosen = true
			break
		default:
			fmt.Println("Yes or no. There is nothing in between.")
		}
	}

	playerOneTag := tictactoe.NO
	var playerTwoTag tictactoe.TAG
	for playerOneTag == tictactoe.NO {
		var tag string

		fmt.Println("Choose your tag [X|O]: ")
		fmt.Scanf("%s", &tag)
		switch tag {
		case "O":
			playerOneTag = tictactoe.O
			playerTwoTag = tictactoe.X
		case "X":
			playerOneTag = tictactoe.X
			playerTwoTag = tictactoe.O
		}
	}

	p1 := tictactoe.NewHumanPlayer(playerOneTag)
	var p2 tictactoe.Player
	switch cpuOpponent {
	case true:
		p2 = tictactoe.NewCPUPlayer(playerTwoTag)
	case false:
		p2 = tictactoe.NewHumanPlayer(playerTwoTag)
	}
	ttt := tictactoe.NewTicTacToe(p1, p2)

	printIndexedBoard()

	err := ttt.Run()
	if err != nil {
		fmt.Println("Game failed nobody wins.")
		os.Exit(1)
	}

	switch ttt.Winner {
	case tictactoe.NO:
		fmt.Println("Even")
	case playerTwoTag:
		fmt.Println("Player 2 won!")
	case playerOneTag:
		fmt.Println("Player 1 won!")
	}
	ttt.Print()
}
