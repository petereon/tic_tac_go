package tic_tac_toe_test

import (
	"testing"

	"github.com/petereon/tic_tac_go/src/tic_tac_toe"
)

func TestCreateBoard(t *testing.T) {
	board := tic_tac_toe.CreateBoard()
	if len(board) != 9 {
		t.Error("Expected board to have 9 spaces")
	}
}
