package main

import (
	"fmt"

	"markus.tictactoe/pkg/tictactoe"
)

func printIndexedBoard() {
	for i := 0; i < 9; i += 3 {
		fmt.Printf("|%d|%d|%d|\n", i, i+1, i+2)
	}
}

func main() {
	ttt := tictactoe.NewTicTacToe()
	playerTag := tictactoe.NO
	var cpuTag tictactoe.TAG
	for playerTag == tictactoe.NO {
		var tag string

		fmt.Println("Choose your tag [X|O]: ")
		fmt.Scanf("%s", &tag)
		switch tag {
		case "O":
			playerTag = tictactoe.O
			cpuTag = tictactoe.X
		case "X":
			playerTag = tictactoe.X
			cpuTag = tictactoe.O
		}
	}

	printIndexedBoard()

	for {
		fmt.Println("Place tag at index: ")
		var index uint
		_, err := fmt.Scanf("%d", &index)
		if err != nil {
			fmt.Printf("Please ... %v\n", err)
			continue
		}

		c, err := tictactoe.IndexToCoord(uint(index))
		if err != nil {
			fmt.Printf("Please ... %v\n", err)
			continue
		}

		err = ttt.Tag(c.X, c.Y, playerTag)
		if err != nil {
			if err == tictactoe.GameOverErr {
				break
			}
			fmt.Printf("Dude ... %v\n", err)
			continue
		}
		ttt.Print()

		err = ttt.CpuPlay(cpuTag)
		if err != nil {
			if err == tictactoe.GameOverErr {
				break
			}
			fmt.Printf("Oops ... %v\n", err)
			break
		}

		ttt.Print()
	}
	switch ttt.Winner {
	case tictactoe.NO:
		fmt.Println("Even")
	case cpuTag:
		fmt.Println("You lost!")
	case playerTag:
		fmt.Println("You won!")
	}
	ttt.Print()
}
