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
	cpuOption, err := tictactoe.StringOptionDialog("Play against CPU?", []string{"y", "n"})
	if err != nil {
		panic(err)
	}
	switch cpuOption {
	case "n", "N":
		cpuOpponent = false
	case "y", "Y", "":
		cpuOpponent = true
	default:
		fmt.Println("Yes or no. There is nothing in between.")
	}

	tagOption, err := tictactoe.StringOptionDialog("Choose your tag ", []string{"X", "O"})
	if err != nil {
		panic(err)
	}
	var playerOneTag, playerTwoTag tictactoe.TAG
	switch tagOption {
	case "O":
		playerOneTag = tictactoe.O
		playerTwoTag = tictactoe.X
	case "X":
		playerOneTag = tictactoe.X
		playerTwoTag = tictactoe.O
	}

	p1 := tictactoe.NewConsolePlayer(playerOneTag, "Player One")
	var p2 tictactoe.Player
	switch cpuOpponent {
	case true:
		p2 = tictactoe.NewCPUPlayer(playerTwoTag)
	case false:
		p2 = tictactoe.NewConsolePlayer(playerTwoTag, "Player Two")
	}
	ttt := tictactoe.NewTicTacToe(p1, p2)

	printIndexedBoard()

	err = ttt.Run()
	if err != nil {
		fmt.Printf("game failed nobody wins: %v", err)
		os.Exit(1)
	}

	switch ttt.Winner {
	case tictactoe.NO:
		fmt.Println("Even")
	case playerTwoTag:
		fmt.Printf("%s won!\n", p2.GetName())
	case playerOneTag:
		fmt.Printf("%s won!\n", p1.GetName())
	}
	ttt.Print()
}
