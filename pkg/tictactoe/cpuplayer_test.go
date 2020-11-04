package tictactoe_test

import (
	"testing"

	"markus.tictactoe/pkg/tictactoe"
)

func TestWin(t *testing.T) {
	cpuPlayer := tictactoe.NewCPUPlayer(tictactoe.X)
	cpuPlayer2 := tictactoe.NewCPUPlayer(tictactoe.O)
	ttt := tictactoe.NewTicTacToe(cpuPlayer, cpuPlayer2)

	ttt.Board[0][2] = tictactoe.X
	ttt.Board[2][1] = tictactoe.O
	ttt.Board[0][1] = tictactoe.X
	ttt.Board[2][2] = tictactoe.O
	c, err := cpuPlayer.RequestMove(ttt.Board)
	if err != nil {
		t.Logf("failed to calculate valid coordinates")
		t.FailNow()
	}
	if c.X != 0 || c.Y != 0 {
		t.Log("failed calculating winning coords at 0,0")
		t.FailNow()
	}
}

func TestWin2(t *testing.T) {
	cpuPlayer := tictactoe.NewCPUPlayer(tictactoe.O)
	cpuPlayer2 := tictactoe.NewCPUPlayer(tictactoe.X)
	ttt := tictactoe.NewTicTacToe(cpuPlayer, cpuPlayer2)
	ttt.Board[0][0] = tictactoe.X
	ttt.Board[0][2] = tictactoe.X
	ttt.Board[1][1] = tictactoe.X
	ttt.Board[1][0] = tictactoe.O
	ttt.Board[2][2] = tictactoe.O
	c, err := cpuPlayer.RequestMove(ttt.Board)
	if err != nil {
		t.Logf("%v", ttt.Board)
		t.Log("failed to calculate valid move")
		t.FailNow()
	}
	ttt.Board[c.X][c.Y] = tictactoe.O

	c, err = cpuPlayer2.RequestMove(ttt.Board)
	if err != nil {
		t.Log("failed calculating winning coords")
		t.FailNow()
	}
	if c.X != 0 || c.Y != 1 {
		t.Log("failed calculating right winning coords at 0,1")
		t.FailNow()
	}
}

func TestCPUvsCPU(t *testing.T) {
	cpuPlayer := tictactoe.NewCPUPlayer(tictactoe.O)
	cpuPlayer2 := tictactoe.NewCPUPlayer(tictactoe.X)
	ttt := tictactoe.NewTicTacToe(cpuPlayer, cpuPlayer2)
	err := ttt.Run()
	if err != nil {
		t.Logf("game failed: %v", err)
		t.FailNow()
	}
	if ttt.Winner != tictactoe.NO {
		t.Log("Game ended uneven")
		t.FailNow()
	}
}

func TestCounterWin(t *testing.T) {
	cpuPlayer := tictactoe.NewCPUPlayer(tictactoe.X)
	cpuPlayer2 := tictactoe.NewCPUPlayer(tictactoe.O)
	ttt := tictactoe.NewTicTacToe(cpuPlayer, cpuPlayer2)
	ttt.Board[1][2] = tictactoe.X
	ttt.Board[2][1] = tictactoe.O
	ttt.Board[0][1] = tictactoe.X
	ttt.Board[2][2] = tictactoe.O
	c, err := cpuPlayer.RequestMove(ttt.Board)
	if err != nil {
		t.Log("Could not calc counter tag")
		t.FailNow()
	}
	if c.X != 2 || c.Y != 0 {
		t.Logf("Wrong counter coords calculated not 2,0")
		t.FailNow()
	}
}
