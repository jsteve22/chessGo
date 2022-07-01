package board_test

import (
	"testing"

	"github.com/jsteve22/chessGo/board"
)

func TestInitBoard(t *testing.T) {
	// this will test that the init function will
	// properly set-up the chessboard for a game
	var cb board.ChessBoard

	// Call InitBoard()
	cb.InitBoard()

	// check that the starting color is white
	if cb.CurrTurn() != 0 {
		t.Errorf("nextMove = %d; want 0\n",cb.CurrTurn())
	}
	
	// get each side and test both of them
	white := cb.GetSide(0)
	black := cb.GetSide(1)
	b := cb.GetBoard()
	var piecePos int8

	// loop through white side and to ensure it was set properly
	for i,p := range white {
		// check that the position of each piece matched the board
		piecePos = p.GetPos()
		if b[piecePos] != &white[i] {
			t.Errorf("%p != %p; want equal\n",b[piecePos],&white[i])
		}

		// check that each piece was alive
		if !p.GetAlive() {
			t.Errorf("%v.alive = %v; want true\n",(string)(p.GetRep()),p.GetAlive())
		}
	}
	
	// loop through black side and to ensure it was set properly
	for i,p := range black {
		// check that the position of each piece matched the board
		piecePos = p.GetPos()
		if b[piecePos] != &black[i] {
			t.Errorf("%p != %p; want equal\n",b[piecePos],&black[i])
		}

		// check that each piece was alive
		if !p.GetAlive() {
			t.Errorf("%v.alive = %v; want true\n",(string)(p.GetRep()),p.GetAlive())
		}
	}

	// loop through middle of board and ensure that nonpiece squares 
	// are nil
	for i := 16; i < 48; i++ {
		if b[i] != nil {
			t.Errorf("board[%d] = %v; want board[%d] = nil\n",i,b[i],i)
		}
	}

	// make sure both sides have castling rights
	castle := cb.GetCastle()
	for i := 0; i < 4; i++ {
		if !castle[i] {
			t.Errorf("Castle[%d] = %v; want true\n",i,castle[i])
		}
	}

	// make sure enpassant is -1
	enpas := cb.GetEnpas()
	if enpas != -1 {
		t.Errorf("Enpas = %v; want -1\n",enpas)
	}
}

func TestFENSet1(t *testing.T) {
	// this will test FENSet by setting up weird configurations

	var cb board.ChessBoard

	// initialize board to initilize pieces
	cb.InitBoard()

	// creating an empty board with no castling rights and no enpassant
	// still white's turn to play
	cb.FENSet("8/8/8/8/8/8/8/8 w - -")

	// check board is empty
	b := cb.GetBoard()
	for i := 0; i < 64; i++ {
		if b[i] != nil {
			t.Errorf("board[%d] = %v; want nil\n",i,b[i])
		}
	}

	// check all pieces are dead
	white := cb.GetSide(0)
	black := cb.GetSide(1)
	for i := 0; i < 16; i++ {
		if white[i].GetAlive() {
			t.Errorf("white[%d].alive = %v; want false\n",i,white[i].GetAlive())
		}
		if black[i].GetAlive() {
			t.Errorf("black[%d].alive = %v; want false\n",i,black[i].GetAlive())
		}
	}

	// check turn is white
	nextTurn := cb.GetNextMove()
	if nextTurn != 0 {
		t.Errorf("nextTurn = %v; want 0\n",nextTurn)
	}

	// check castling rights are all false
	castle := cb.GetCastle()
	for i := 0; i < 4; i++ {
		if castle[i] {
			t.Errorf("Castle[%d] = %v; want false\n",i,castle[i])
		}
	}

	// check enpassant is -1
	enpas := cb.GetEnpas()
	if enpas != -1 {
		t.Errorf("Enpas = %v; want -1\n",enpas)
	}
}

func TestFENSet2(t *testing.T) {
	// this will test FENSet by setting up weird configurations

	var cb board.ChessBoard

	// initialize board to initilize pieces
	cb.InitBoard()

	// creating an empty board with no castling rights and no enpassant
	// still white's turn to play
	cb.FENSet("8/8/8/8/8/8/8/8 b kq a3")

	// check board is empty
	b := cb.GetBoard()
	for i := 0; i < 64; i++ {
		if b[i] != nil {
			t.Errorf("board[%d] = %v; want nil\n",i,b[i])
		}
	}

	// check all pieces are dead
	white := cb.GetSide(0)
	black := cb.GetSide(1)
	for i := 0; i < 16; i++ {
		if white[i].GetAlive() {
			t.Errorf("white[%d].alive = %v; want false\n",i,white[i].GetAlive())
		}
		if black[i].GetAlive() {
			t.Errorf("black[%d].alive = %v; want false\n",i,black[i].GetAlive())
		}
	}

	// check turn is black
	nextTurn := cb.GetNextMove()
	if nextTurn != 1 {
		t.Errorf("nextTurn = %v; want 0\n",nextTurn)
	}

	// check castling rights for white are false
	castle := cb.GetCastle()
	for i := 0; i < 4; i++ {
		if i < 2 {
			if castle[i] {
				t.Errorf("Castle[%d] = %v; want false\n",i,castle[i])
			}
		} else {
			if !castle[i] {
				t.Errorf("Castle[%d] = %v; want true\n",i,castle[i])
			}
		}
	}

	// check enpassant is 16
	enpas := cb.GetEnpas()
	if enpas != 16 {
		t.Errorf("Enpas = %v; want 16\n",enpas)
	}
}

func TestFENSet3(t *testing.T) {
	// this will test FENSet by setting up weird configurations

	var cb board.ChessBoard

	// initialize board to initilize pieces
	cb.InitBoard()

	// creating an empty board with no castling rights and no enpassant
	// still white's turn to play
	cb.FENSet("8/8/8/8/8/8/8/8 b kq A3")

	// check board is empty
	b := cb.GetBoard()
	for i := 0; i < 64; i++ {
		if b[i] != nil {
			t.Errorf("board[%d] = %v; want nil\n",i,b[i])
		}
	}

	// check all pieces are dead
	white := cb.GetSide(0)
	black := cb.GetSide(1)
	for i := 0; i < 16; i++ {
		if white[i].GetAlive() {
			t.Errorf("white[%d].alive = %v; want false\n",i,white[i].GetAlive())
		}
		if black[i].GetAlive() {
			t.Errorf("black[%d].alive = %v; want false\n",i,black[i].GetAlive())
		}
	}

	// check turn is black
	nextTurn := cb.GetNextMove()
	if nextTurn != 1 {
		t.Errorf("nextTurn = %v; want 0\n",nextTurn)
	}

	// check castling rights for white are false
	castle := cb.GetCastle()
	for i := 0; i < 4; i++ {
		if i < 2 {
			if castle[i] {
				t.Errorf("Castle[%d] = %v; want false\n",i,castle[i])
			}
		} else {
			if !castle[i] {
				t.Errorf("Castle[%d] = %v; want true\n",i,castle[i])
			}
		}
	}

	// check enpassant is 16
	enpas := cb.GetEnpas()
	if enpas != 16 {
		t.Errorf("Enpas = %v; want 16\n",enpas)
	}
}

func BenchmarkPerftCalcultationsDepth_5(b *testing.B) {
	// this benchmark test will call Perft which
	// is a recursive function that continuously 
	// generates moves. This will test how fast 
	// move generation is for my chess program

	// create and init a chessboard
	var cb board.ChessBoard
	cb.InitBoard()
	n := 2

	for i := 0; i < b.N; i++ {
		// cb.FENSet("k7/pppppppp/8/8/8/8/PPPPPPPP/7K w - -")
		cb.FENSet("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -")
		cb.Perft(n)
	}
}