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
	err := ttt.CpuPlay(tictactoe.X)
	if err != tictactoe.GameOverErr {
		t.Log("Could not calc winning tag")
		t.FailNow()
	}
	if ttt.Board[0][0] != tictactoe.X {
		t.Log("Wrong winning coords calculated not 0,0")
		t.FailNow()
	}
}

func TestCounterWin(t *testing.T) {
	ttt := tictactoe.NewTicTacToe()
	ttt.Board[1][2] = tictactoe.X
	ttt.Board[2][1] = tictactoe.O
	ttt.Board[0][1] = tictactoe.X
	ttt.Board[2][2] = tictactoe.O
	err := ttt.CpuPlay(tictactoe.X)
	if err != nil {
		t.Log("Could not calc counter tag")
		t.FailNow()
	}
	if ttt.Board[2][0] != tictactoe.X {
		t.Logf("Wrong counter coords calculated not 2,0")
		t.FailNow()
	}
}
