package tictactoe

import (
	"fmt"
)

type ConsolePlayer struct {
	GenericPlayer
}

func NewConsolePlayer(tag TAG, defaultName string) *ConsolePlayer {
	name, err := StringDialogDefault("Enter name", defaultName)
	if err != nil {
		panic(err.Error())
	}
	return &ConsolePlayer{GenericPlayer{tag: tag, name: name}}
}

func (h *ConsolePlayer) RequestMove(board [3][3]TAG) (*Coordinate, error) {
	for {
		var index uint
		index, err := UintDialog("Place tag at index")
		if err != nil {
			fmt.Printf("Please ... %v\n", err)
			continue
		}

		c, err := IndexToCoord(uint(index))
		if err != nil {
			fmt.Printf("Please ... %v\n", err)
			continue
		}

		return c, nil
	}
}
