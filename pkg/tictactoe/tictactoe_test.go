package tictactoe_test

import (
	"testing"

	"markus.tictactoe/pkg/tictactoe"
)

func Test_Ctor(t *testing.T) {
	ttt := tictactoe.NewTicTacToe()
	if ttt.Board[0][0] != tictactoe.NO {
		t.FailNow()
	}
}

func TestTicTacToe_Tag(t *testing.T) {
	ttt := tictactoe.NewTicTacToe()
	err := ttt.Tag(0, 0, tictactoe.X)
	if err != nil {
		t.Log("Can not tag")
		t.FailNow()
	}
	err = ttt.Tag(0, 0, tictactoe.X)
	if err == nil {
		t.Log("Should not tag")
		t.FailNow()
	}
}

func TestWin(t *testing.T) {
	ttt := tictactoe.NewTicTacToe()
	ttt.Board[0][2] = tictactoe.X
	ttt.Board[2][1] = tictactoe.O
	ttt.Board[0][1] = tictactoe.X
	ttt.Board[2][2] = tictactoe.O
	c, err := ttt.CalcMove(tictactoe.X)
	if err != nil {
		t.Log("Could not calc wining tag")
		t.FailNow()
	}
	if c.X != 0 || c.Y != 0 {
		t.Logf("Wrong wining coords calculated %d,%d is not 0,0", c.X, c.Y)
		t.FailNow()
	}
}

func TestCounterWin(t *testing.T) {
	ttt := tictactoe.NewTicTacToe()
	ttt.Board[1][2] = tictactoe.X
	ttt.Board[2][1] = tictactoe.O
	ttt.Board[0][1] = tictactoe.X
	ttt.Board[2][2] = tictactoe.O
	c, err := ttt.CalcMove(tictactoe.X)
	if err != nil {
		t.Log("Could not calc counter tag")
		t.FailNow()
	}
	if c.X != 2 || c.Y != 0 {
		t.Logf("Wrong counter coords calculated %d,%d is not 2,0", c.X, c.Y)
		t.FailNow()
	}
}
