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

type Lane struct {
	Lane [3]Coordinate
}

var lanes [8]Lane

var GameOverErr error = errors.New("Game over")

func init() {
	lanes = initLanes()
}

func initLanes() [8]Lane {
	lanes := [8]Lane{}
	for i := 0; i < 3; i++ {
		lanes[i] = Lane{
			[3]Coordinate{
				{uint(i), 0},
				{uint(i), 1},
				{uint(i), 2},
			},
		}
		lanes[3+i] = Lane{
			[3]Coordinate{
				{0, uint(i)},
				{1, uint(i)},
				{2, uint(i)},
			},
		}
	}
	lanes[6] = Lane{
		[3]Coordinate{
			{0, 0},
			{1, 1},
			{2, 2},
		},
	}
	lanes[7] = Lane{
		[3]Coordinate{
			{0, 2},
			{1, 1},
			{2, 0},
		},
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
		for _, c := range l.Lane {
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
		fmt.Printf("%v", err)
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
		for _, c := range l.Lane {
			switch t.Board[c.X][c.Y] {
			case tag:
				myTag++
			case NO:
				laneCandidates = append(laneCandidates, &Coordinate{c.X, c.Y})
			default:
				otherTag++
			}
		}
		if len(laneCandidates) == 0 {
			// remove lane
			continue
		}
		// win or rescue
		if myTag == 2 {
			for i := 0; i < 10; i++ {
				laneCandidates = append(laneCandidates, laneCandidates[0])
			}
		}
		if otherTag == 2 {
			for i := 0; i < 5; i++ {
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
