package tictactoe

import (
	"errors"
	"fmt"
)

type TicTacToe struct {
	Board  [3][3]TAG
	Move   uint
	Winner TAG
}

type Coordinate struct {
	X uint
	Y uint
}

var lanes [][]*Coordinate

var GameOverErr error = errors.New("Game over")

func init() {
	lanes = initLanes()
}

func initLanes() [][]*Coordinate {
	// init all coordinates
	coords := []*Coordinate{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			coords = append(coords, &Coordinate{X: uint(i), Y: uint(j)})
		}
	}

	lanes := [][]*Coordinate{
		// horizontal
		coords[0:3],
		coords[3:6],
		coords[6:],

		// vertical
		{coords[0], coords[3], coords[6]},
		{coords[1], coords[4], coords[7]},
		{coords[2], coords[5], coords[8]},

		// diagonal
		{coords[0], coords[4], coords[8]},
		{coords[2], coords[4], coords[6]},
	}

	return lanes
}

type TAG uint

const (
	NO TAG = iota
	O
	X
)

func NewTicTacToe() *TicTacToe {
	return &TicTacToe{
		Board:  [3][3]TAG{},
		Move:   0,
		Winner: NO,
	}
}

func (t *TicTacToe) Tag(x uint, y uint, tag TAG) error {
	if t.Winner != NO {
		return GameOverErr
	}
	if t.Board[x][y] != NO {
		return errors.New("Already set")
	}
	t.Board[x][y] = tag
	err := t.checkGameStatus(tag)
	if err != nil {
		return err
	}
	t.Move++
	return nil
}

func (t *TicTacToe) checkGameStatus(tag TAG) error {
	for _, l := range lanes {
		myTag := 0
		for _, c := range l {
			switch t.Board[c.X][c.Y] {
			case tag:
				myTag++
			}
		}
		if myTag == 3 {
			t.Winner = tag
			return GameOverErr
		}
	}
	return nil
}

func (t *TicTacToe) CpuPlay(cpuTag TAG) error {
	c, err := t.calcMove(cpuTag)
	if err != nil {
		return err
	}
	err = t.Tag(c.X, c.Y, cpuTag)
	if err != nil {
		if err == GameOverErr {
			return err
		}
		fmt.Printf("Oops ... %v", err)
	}
	return nil
}

func (t *TicTacToe) calcMove(tag TAG) (*Coordinate, error) {
	candidates := map[*Coordinate]int{}
	for _, l := range lanes {
		myTag := 0
		otherTag := 0
		laneCandidates := []*Coordinate{}
		for _, c := range l {
			switch t.Board[c.X][c.Y] {
			case tag:
				myTag++
			case NO:
				laneCandidates = append(laneCandidates, c)
			default:
				otherTag++
			}
		}
		if len(laneCandidates) == 0 {
			// remove lane
			continue
		}
		// weight win and rescue
		if myTag == 2 {
			for i := 0; i < 16; i++ {
				laneCandidates = append(laneCandidates, laneCandidates[0])
			}
		}
		if otherTag == 2 {
			for i := 0; i < 8; i++ {
				laneCandidates = append(laneCandidates, laneCandidates[0])
			}
		}
		for _, lc := range laneCandidates {
			candidates[lc] = candidates[lc] + 1
		}
	}
	var result *Coordinate
	count := 0
	for k, v := range candidates {
		if v > count {
			result = k
			count = v
		}
	}
	if result == nil {
		return nil, GameOverErr
	}
	return result, nil
}

func (t *TicTacToe) Print() {
	fmt.Printf("Move %d\n", t.Move)
	for _, e := range t.Board {
		for _, e2 := range e {
			fmt.Print("|")
			if e2 == NO {
				fmt.Print(" ")
			}
			if e2 == X {
				fmt.Print("X")
			}
			if e2 == O {
				fmt.Print("O")
			}
		}
		fmt.Println("|")
	}
}
