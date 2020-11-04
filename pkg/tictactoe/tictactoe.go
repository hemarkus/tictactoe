package tictactoe

import (
	"errors"
	"fmt"
)

type TicTacToe struct {
	Board     [3][3]TAG
	Move      uint
	Winner    TAG
	PlayerOne Player
	PlayerTwo Player
	GameOver  bool
}

type Coordinate struct {
	X uint
	Y uint
}

var lanes [][]*Coordinate
var coords []*Coordinate

var GameOverErr error = errors.New("Game over")

func init() {
	lanes = initLanes()
}

func initLanes() [][]*Coordinate {
	// init all coordinates
	coords = []*Coordinate{}
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

func NewTicTacToe(p1 Player, p2 Player) *TicTacToe {
	return &TicTacToe{
		Board:     [3][3]TAG{},
		Move:      0,
		Winner:    NO,
		PlayerOne: p1,
		PlayerTwo: p2,
		GameOver:  false,
	}
}

func (t *TicTacToe) Run() error {
	t.Print()
	for !t.GameOver {
		for i, p := range []Player{t.PlayerOne, t.PlayerTwo} {
			fmt.Printf("Player %d\n", i+1)
			c, err := p.RequestMove(t.Board)
			if err != nil {
				if err == GameOverErr {
					t.GameOver = true
					break
				}
				return err
			}
			err = t.tag(c, p.GetTag())
			if err != nil {
				if err == GameOverErr {
					t.win(p)
					break
				}
				return err
			}
			t.Print()
		}
	}
	return nil
}

func (t *TicTacToe) tag(coordinate *Coordinate, tag TAG) error {
	if t.Winner != NO {
		return GameOverErr
	}
	if t.Board[coordinate.X][coordinate.Y] != NO {
		return errors.New("Already set")
	}
	t.Board[coordinate.X][coordinate.Y] = tag
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
			return GameOverErr
		}
	}
	return nil
}

func (t *TicTacToe) win(p Player) {
	t.GameOver = true
	t.Winner = p.GetTag()
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
