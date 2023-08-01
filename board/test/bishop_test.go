package board_test

import (
	"testing"

	"github.com/jsteve22/chessGo/board"
)

func TestBishopMove1(t *testing.T) {
	// Init Board for testing
	game := board.LoadBoard("8/8/8/8/3B4/8/8/8 w - -")

	// get moves
	moves := board.Perft(game, 1)

	if moves != 13 {
		t.Errorf("len(moves) = %v; want 13\n",moves)
	}
}

func TestBishopMove2(t *testing.T) {
	// Init Board for testing
	game := board.LoadBoard("8/8/8/4p3/3B4/8/8/8 w - -")

	// get moves
	moves := board.Perft(game, 1)

	if moves != 10 {
		t.Errorf("len(moves) = %v; want 10\n",moves)
	}
}