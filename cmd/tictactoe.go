package main

import (
	"fmt"

	"markus.tictactoe/pkg/tictactoe"
)

func main() {
	ttt := tictactoe.NewTicTacToe()
	playerTag := tictactoe.NO
	var cpuTag tictactoe.TAG
	for playerTag == tictactoe.NO {
		var tag string

		fmt.Println("Choose tag: ")
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

	for {
		ttt.Print()
		fmt.Println("Tag: ")
		var x uint
		var y uint
		i, err := fmt.Scanf("%d,%d", &x, &y)
		if err != nil {
			fmt.Printf("Please ... %v\n", err)
			continue
		}
		if i != 2 {
			fmt.Println("Come on 2 dimensions")
			continue
		}
		if x < 0 || x > 2 || y < 0 || y > 2 {
			fmt.Println("Aren't you a programmer? The very first digit is 0")
			continue
		}
		err = ttt.Tag(x, y, playerTag)
		if err != nil {
			if err == tictactoe.GameOverErr {
				ttt.Print()
				fmt.Println("You won!")
				break
			}
			fmt.Printf("Dude ... %v\n", err)
			continue
		}
		ttt.Print()

		err = ttt.CpuPlay(cpuTag)
		if err != nil {
			if err == tictactoe.GameOverErr {
				ttt.Print()
				fmt.Println("You lost!")
				break
			}
			fmt.Printf("Oops ... %v", err)
			break
		}
	}
}
