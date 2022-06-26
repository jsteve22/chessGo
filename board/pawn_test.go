package board_test

import (
	"testing"

	"github.com/jsteve22/chessGo/board"
)

func TestPawnMove1(t *testing.T) {
	// this will test if the board generates the right moves for pawns

	var cb board.ChessBoard

	// Init Board for testing
	cb.InitBoard()

	// set board with one pawn on it
	cb.FENSet("8/8/8/8/8/8/4P3/8 w - -")

	// get the board to access certain pieces
	b := cb.GetBoard()

	// get piece at E2
	pawn := b[12]
	// generate moves from the piece
	board.PawnMove(&cb, *pawn)

	// get moves
	moves := cb.GetMoves()

	if len( *moves ) != 2 {
		t.Errorf("len(moves) = %v; want 2\n",len( *moves ))
	}

	for _,m := range *moves {
		if m.GetStart() != 12 {
			t.Errorf("m.start = %v; want 12(E2)\n",m.GetStart())
		}
		if m.GetEnd() != 20 && m.GetEnd() != 28 {
			t.Errorf("m.end = %v; want 20(E3) or 28(E4)\n",m.GetEnd())
		}
	}
}

func TestPawnMove2(t *testing.T) {
	// this will test if the board generates the right moves for pawns

	var cb board.ChessBoard

	// Init Board for testing
	cb.InitBoard()

	// set board with one pawn on it
	cb.FENSet("8/4p3/8/8/8/8/8/8 b - -")

	// get the board to access certain pieces
	b := cb.GetBoard()

	// get piece at E7
	pawn := b[52] 
	// generate moves from the piece
	board.PawnMove(&cb, *pawn)

	// get moves
	moves := cb.GetMoves()

	if len( *moves ) != 2 {
		t.Errorf("len(moves) = %v; want 2\n",len( *moves ))
	}

	for _,m := range *moves {
		if m.GetStart() != 52 {
			t.Errorf("m.start = %v; want 52(E7)\n",m.GetStart())
		}
		if m.GetEnd() != 44 && m.GetEnd() != 36 {
			t.Errorf("m.end = %v; want 44(E6) or 36(E5)\n",m.GetEnd())
		}
	}

}

func TestPawnMove3(t *testing.T) {
	// this will test if the board generates the right moves for pawns

	var cb board.ChessBoard

	// Init Board for testing
	cb.InitBoard()

	// set board with one pawn on it
	cb.FENSet("8/8/8/8/8/4P3/8/8 w - -")

	// get the board to access certain pieces
	b := cb.GetBoard()

	// get piece at E3
	pawn := b[20] 
	// generate moves from the piece
	board.PawnMove(&cb, *pawn)

	// get moves
	moves := cb.GetMoves()

	if len( *moves ) != 1 {
		t.Errorf("len(moves) = %v; want 1\n",len( *moves ))
	}

	for _,m := range *moves {
		if m.GetStart() != 20 {
			t.Errorf("m.start = %v; want 20(E3)\n",m.GetStart())
		}
		if m.GetEnd() != 28 {
			t.Errorf("m.end = %v; want 28(E4)\n",m.GetEnd())
		}
	}
}

func TestPawnMove4(t *testing.T) {
	// this will test if the board generates the right moves for pawns

	var cb board.ChessBoard

	// Init Board for testing
	cb.InitBoard()

	// set board with one pawn on it
	cb.FENSet("8/8/8/8/8/4p3/8/8 b - -")

	// get the board to access certain pieces
	b := cb.GetBoard()

	// get piece at E3
	pawn := b[20] 
	// generate moves from the piece
	board.PawnMove(&cb, *pawn)

	// get moves
	moves := cb.GetMoves()

	if len( *moves ) != 1 {
		t.Errorf("len(moves) = %v; want 1\n",len( *moves ))
	}

	for _,m := range *moves {
		if m.GetStart() != 20 {
			t.Errorf("m.start = %v; want 20(E3)\n",m.GetStart())
		}
		if m.GetEnd() != 12 {
			t.Errorf("m.end = %v; want 12(E2)\n",m.GetEnd())
		}
	}
}

func TestPawnMove5(t *testing.T) {
	// this will test if the board generates the right moves for pawns

	var cb board.ChessBoard

	// Init Board for testing
	cb.InitBoard()

	// set board with two pawns on it
	cb.FENSet("8/8/5p2/4P3/8/8/8/8 w - -")

	// get the board to access certain pieces
	b := cb.GetBoard()

	// get piece at E5
	pawn := b[36] 
	// generate moves from the piece
	board.PawnMove(&cb, *pawn)

	// get moves
	moves := cb.GetMoves()

	if len( *moves ) != 2 {
		t.Errorf("len(moves) = %v; want 2\n",len( *moves ))
	}

	for _,m := range *moves {
		if m.GetStart() != 36 {
			t.Errorf("m.start = %v; want 36(E5)\n",m.GetStart())
		}
		if m.GetEnd() != 44 && m.GetEnd() != 45 {
			t.Errorf("m.end = %v; want 44(E6) or 45(F6)\n",m.GetEnd())
		}
	}
}

func TestPawnMove6(t *testing.T) {
	// this will test if the board generates the right moves for pawns

	var cb board.ChessBoard

	// Init Board for testing
	cb.InitBoard()

	// set board with two pawns on it
	cb.FENSet("8/8/8/4Pp2/8/8/8/8 w - -")

	// get the board to access certain pieces
	b := cb.GetBoard()

	// get piece at E5
	pawn := b[36] 
	// generate moves from the piece
	board.PawnMove(&cb, *pawn)

	// get moves
	moves := cb.GetMoves()

	if len( *moves ) != 1 {
		t.Errorf("len(moves) = %v; want 1\n",len( *moves ))
	}

	for _,m := range *moves {
		if m.GetStart() != 36 {
			t.Errorf("m.start = %v; want 36(E5)\n",m.GetStart())
		}
		if m.GetEnd() != 44 {
			t.Errorf("m.end = %v; want 44(E6)\n",m.GetEnd())
		}
	}
}
 
func TestPawnMove7(t *testing.T) {
	// this will test if the board generates the right moves for pawns

	var cb board.ChessBoard

	// Init Board for testing
	cb.InitBoard()

	// set board with two pawns on it
	cb.FENSet("8/8/8/4Pp2/8/8/8/8 w - f6")

	// get the board to access certain pieces
	b := cb.GetBoard()

	// get piece at E5
	pawn := b[36] 
	// generate moves from the piece
	board.PawnMove(&cb, *pawn)

	// get moves
	moves := cb.GetMoves()

	if len( *moves ) != 2 {
		t.Errorf("len(moves) = %v; want 2\n",len( *moves ))
	}

	for _,m := range *moves {
		if m.GetStart() != 36 {
			t.Errorf("m.start = %v; want 36(E5)\n",m.GetStart())
		}
		if m.GetEnd() != 44 && m.GetEnd() != 45 {
			t.Errorf("m.end = %v; want 44(E6) or 45(F6)\n",m.GetEnd())
		}
	}
}

func TestPawnMove8(t *testing.T) {
	// this will test if the board generates the right moves for pawns

	var cb board.ChessBoard

	// Init Board for testing
	cb.InitBoard()

	// set board with four pawns on it
	cb.FENSet("8/8/3ppp2/4P3/8/8/8/8 w - -")

	// get the board to access certain pieces
	b := cb.GetBoard()

	// get piece at E5
	pawn := b[36] 
	// generate moves from the piece
	board.PawnMove(&cb, *pawn)

	// get moves
	moves := cb.GetMoves()

	if len( *moves ) != 2 {
		t.Errorf("len(moves) = %v; want 2\n",len( *moves ))
	}

	for _,m := range *moves {
		if m.GetStart() != 36 {
			t.Errorf("m.start = %v; want 36(E5)\n",m.GetStart())
		}
		if m.GetEnd() != 43 && m.GetEnd() != 45 {
			t.Errorf("m.end = %v; want 43(D6) or 45(F6)\n",m.GetEnd())
		}
	}
}

func TestPawnMove9(t *testing.T) {
	// this will test if the board generates the right moves for pawns

	var cb board.ChessBoard

	// Init Board for testing
	cb.InitBoard()

	// set board with four pawns on it
	cb.FENSet("8/8/4p3/4P3/8/8/8/8 w - -")

	// get the board to access certain pieces
	b := cb.GetBoard()

	// get piece at E5
	pawn := b[36] 
	// generate moves from the piece
	board.PawnMove(&cb, *pawn)

	// get moves
	moves := cb.GetMoves()

	if len( *moves ) > 0 {
		t.Errorf("moves = %v; want nothing\n",*moves)
	}
}

func TestPawnMove10(t *testing.T) {
	// this will test if the board generates the right moves for pawns

	var cb board.ChessBoard

	// Init Board for testing
	cb.InitBoard()

	// set board with four pawns on it
	cb.FENSet("8/8/p6p/P7/8/8/8/8 w - -")

	// get the board to access certain pieces
	b := cb.GetBoard()

	// get piece at A5
	pawn := b[32] 
	// generate moves from the piece
	board.PawnMove(&cb, *pawn)

	// get moves
	moves := cb.GetMoves()

	if len( *moves ) != 0 {
		t.Errorf("moves = %v; want nothing\n",*moves)
	}
}

func TestPawnMove11(t *testing.T) {
	// this will test if the board generates the right moves for pawns

	var cb board.ChessBoard

	// Init Board for testing
	cb.InitBoard()

	// set board so that taking enpassant will lead to discovered check
	// only move is to push pawn foward once
	cb.FENSet("7k/8/8/KPp4r/8/8/8/8 w - C6")

	// get the board to access certain pieces
	b := cb.GetBoard()

	// get piece at B5
	pawn := b[33] 
	// generate moves from the piece
	board.PawnMove(&cb, *pawn)

	// get moves
	moves := cb.GetMoves()

	if len( *moves ) != 1 {
		t.Errorf("len(moves) = %v; want 1\n",len( *moves ))
	}

	for _,m := range *moves {
		if m.GetStart() != 33 {
			t.Errorf("m.start = %v; want 33(B5)\n",m.GetStart())
		}
		if m.GetEnd() != 41 {
			t.Errorf("m.end = %v; want 41(B6)\n",m.GetEnd())
		}
	}
}

func TestPawnMove12(t *testing.T) {
	// this will test if the board generates the right moves for pawns

	var cb board.ChessBoard

	// Init Board for testing
	cb.InitBoard()

	// pawn can push once or take enpassant
	cb.FENSet("K6k/8/8/1Pp4r/8/8/8/8 w - c6")

	// get the board to access certain pieces
	b := cb.GetBoard()

	// get piece at B5
	pawn := b[33] 
	// generate moves from the piece
	board.PawnMove(&cb, *pawn)

	// get moves
	moves := cb.GetMoves()

	if len( *moves ) != 2 {
		t.Errorf("len(moves) = %v; want 2\nmoves = %v\n",len( *moves ),*moves)
	}

	for _,m := range *moves {
		if m.GetStart() != 33 {
			t.Errorf("m.start = %v; want 33(B5)\n",m.GetStart())
		}
		if m.GetEnd() != 41 && m.GetEnd() != 42 {
			t.Errorf("m.end = %v; want 41(B6) or 42(C6)\n",m.GetEnd())
		}
	}
}
