package tictactoe

import "fmt"

type HumanPlayer struct {
	GenericPlayer
}

func NewHumanPlayer(tag TAG) *HumanPlayer {
	return &HumanPlayer{GenericPlayer{tag: tag}}
}

func (h *HumanPlayer) RequestMove(board [3][3]TAG) (*Coordinate, error) {
	for {
		fmt.Println("Place tag at index: ")
		var index uint
		_, err := fmt.Scanf("%d", &index)
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
