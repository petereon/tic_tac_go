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

func TestPlaceMark(t *testing.T) {
	board := tic_tac_toe.CreateBoard()
	board, err := board.PlaceMark(0, "X")
	if err != nil {
		t.Error("Expected no error")
	}
	if board[0] != "X" {
		t.Error("Expected board to have X at position 0")
	}
	for _, v := range board[1:] {
		if v != " " {
			t.Error("Expected board to have empty spaces")
		}
	}
}

func TestPlaceMarkFails(t *testing.T) {
	board := tic_tac_toe.CreateBoard()
	board, _ = board.PlaceMark(0, "X")
	board, err := board.PlaceMark(0, "O")
	if err == nil {
		t.Error("Expected error")
	}
	if board[0] == "O" {
		t.Error("Expected board to have X at position 0")
	}
}
