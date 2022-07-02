package board_test

import (
	"testing"

	"github.com/jsteve22/chessGo/board"
)

func TestBishopMove1(t *testing.T) {
	// check bishop moves

	// Init Board for testing
	var cb board.ChessBoard
	cb.InitBoard()

	// set board with one bishop in the middle
	cb.FENSet("8/8/8/8/3B4/8/8/8 w - -")

	// get access to board
	b := cb.GetBoard()

	// get piece
	bishop := b[27]

	// generate moves for piece
	board.BishopMove(&cb, *bishop)

	// get moves
	moves := cb.GetMoves()

	if len( *moves ) != 13 {
		t.Errorf("len(moves) = %v; want 13\n",len( *moves ))
	}
}

func TestBishopMove2(t *testing.T) {
	// check bishop moves

	// Init Board for testing
	var cb board.ChessBoard
	cb.InitBoard()

	// set board with one bishop in the middle
	cb.FENSet("8/8/8/4p3/3B4/8/8/8 w - -")

	// get access to board
	b := cb.GetBoard()

	// get piece
	bishop := b[27]

	// generate moves for piece
	board.BishopMove(&cb, *bishop)

	// get moves
	moves := cb.GetMoves()

	if len( *moves ) != 10 {
		t.Errorf("len(moves) = %v; want 10\n",len( *moves ))
	}
}
