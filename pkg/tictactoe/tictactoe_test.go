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
		t.Log("failed to calc winning tag")
		t.FailNow()
	}
	if ttt.Board[0][0] != tictactoe.X {
		t.Log("failed calculating winning coords at 0,0")
		t.FailNow()
	}
}

func TestWin2(t *testing.T) {
	ttt := tictactoe.NewTicTacToe()
	ttt.Board[0][0] = tictactoe.X
	ttt.Board[0][2] = tictactoe.X
	ttt.Board[1][1] = tictactoe.X
	ttt.Board[1][0] = tictactoe.O
	ttt.Board[2][2] = tictactoe.O
	err := ttt.CpuPlay(tictactoe.O)
	if err != nil {
		t.Logf("%v", ttt.Board)
		t.Log("failed to calculate valid move")
		t.FailNow()
	}
	err = ttt.CpuPlay(tictactoe.X)
	if err != tictactoe.GameOverErr {
		t.Log("failed calculating winning coords at 0,1")
		t.FailNow()
	}
}

func TestCPUvsCPU(t *testing.T) {
	ttt := tictactoe.NewTicTacToe()
	for ttt.Move < 10 {
		err := ttt.CpuPlay(tictactoe.O)
		if err == tictactoe.GameOverErr {
			if ttt.Winner != tictactoe.NO {
				t.Log("CPU match ended uneven")
				t.FailNow()
			}
			break
		}
		err = ttt.CpuPlay(tictactoe.X)
		if err == tictactoe.GameOverErr {
			if ttt.Winner != tictactoe.NO {
				t.Log("CPU match ended uneven")
				t.FailNow()
			}
			break
		}
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
