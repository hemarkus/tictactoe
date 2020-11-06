package tictactoe

import (
	"errors"
)

func IndexToCoord(index uint) (*Coordinate, error) {
	if index >= uint(len(coords)) {
		return nil, errors.New("learn counting first")
	}
	coord := coords[index]
	return coord, nil
}
