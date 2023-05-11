package tic_tac_toe

import "fmt"

type Board []string

func CreateBoard() Board {
	return Board{" ", " ", " ", " ", " ", " ", " ", " ", " "}
}

func (board Board) PlaceMark(position int, mark string) (Board, error) {
	if board[position] != " " {
		return board, fmt.Errorf("position already taken")
	}
	board[position] = mark
	return board, nil
}
