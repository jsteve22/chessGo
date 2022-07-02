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

func TestPerft1(t *testing.T) {
	// this test will test the accuracy of the
	// engine's current perft evaluation compared
	// to the evaluations on https://www.chessprogramming.org/Perft_Results

	// initiliaze chess board
	var cb board.ChessBoard
	cb.InitBoard()

	// determine depth for Perft
	// n := 5
	n := 3

	// use Fen representation from website
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -"

	// set board and check perft with depth n
	cb.FENSet(fen)
	x := cb.Perft(n)

	// This is the precalculated node count from chessprogramming.org
	// nodes := (uint64)(4865609) 
	nodes := (uint64)(8902) 

	if x != nodes {
		t.Errorf("Perft(%v) = %v; want %v\n",n,x,nodes)
	}
}

func TestPerft2(t *testing.T) {
	// this test will test the accuracy of the
	// engine's current perft evaluation compared
	// to the evaluations on https://www.chessprogramming.org/Perft_Results

	// initiliaze chess board
	var cb board.ChessBoard
	cb.InitBoard()

	// determine depth for Perft
	// n := 4
	n := 3

	// use Fen representation from website
	fen := "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - "

	// set board and check perft with depth n
	cb.FENSet(fen)
	x := cb.Perft(n)

	// This is the precalculated node count from chessprogramming.org
	// nodes := (uint64)(4085603) 
	nodes := (uint64)(97862)

	if x != nodes {
		t.Errorf("Perft(%v) = %v; want %v\n",n,x,nodes)
	}
}

func TestPerft3(t *testing.T) {
	// this test will test the accuracy of the
	// engine's current perft evaluation compared
	// to the evaluations on https://www.chessprogramming.org/Perft_Results

	// initiliaze chess board
	var cb board.ChessBoard
	cb.InitBoard()

	// determine depth for Perft
	n := 5

	// use Fen representation from website
	fen := "8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - -"

	// set board and check perft with depth n
	cb.FENSet(fen)
	x := cb.Perft(n)

	// This is the precalculated node count from chessprogramming.org
	nodes := (uint64)(674624) 

	if x != nodes {
		t.Errorf("Perft(%v) = %v; want %v\n",n,x,nodes)
	}
}

func TestPerft4(t *testing.T) {
	// this test will test the accuracy of the
	// engine's current perft evaluation compared
	// to the evaluations on https://www.chessprogramming.org/Perft_Results

	// initiliaze chess board
	var cb board.ChessBoard
	cb.InitBoard()

	// determine depth for Perft
	// n := 5
	n := 3

	// use Fen representation from website
	fen := "r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq -"

	// set board and check perft with depth n
	cb.FENSet(fen)
	x := cb.Perft(n)

	// This is the precalculated node count from chessprogramming.org
	// nodes := (uint64)(15833292) 
	nodes := (uint64)(9467)

	if x != nodes {
		t.Errorf("Perft(%v) = %v; want %v\n",n,x,nodes)
	}
}

func TestPerft5(t *testing.T) {
	// this test will test the accuracy of the
	// engine's current perft evaluation compared
	// to the evaluations on https://www.chessprogramming.org/Perft_Results

	// initiliaze chess board
	var cb board.ChessBoard
	cb.InitBoard()

	// determine depth for Perft
	// n := 5
	n := 3

	// use Fen representation from website
	fen := "rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ -"

	// set board and check perft with depth n
	cb.FENSet(fen)
	x := cb.Perft(n)

	// This is the precalculated node count from chessprogramming.org
	// nodes := (uint64)(89941194) 
	nodes := (uint64)(62379) 

	if x != nodes {
		t.Errorf("Perft(%v) = %v; want %v\n",n,x,nodes)
	}
}

func TestPerft6(t *testing.T) {
	// this test will test the accuracy of the
	// engine's current perft evaluation compared
	// to the evaluations on https://www.chessprogramming.org/Perft_Results

	// initiliaze chess board
	var cb board.ChessBoard
	cb.InitBoard()

	// determine depth for Perft
	// n := 5
	n := 3

	// use Fen representation from website
	fen := "r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - -"

	// set board and check perft with depth n
	cb.FENSet(fen)
	x := cb.Perft(n)

	// This is the precalculated node count from chessprogramming.org
	// nodes := (uint64)(164075551) 
	nodes := (uint64)(89890) 

	if x != nodes {
		t.Errorf("Perft(%v) = %v; want %v\n",n,x,nodes)
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
	n := 5

	fen1 := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -"
	// fen2 := "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - "
	// fen3 := "8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - -"
	// fen4 := "r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq -"
	// fen5 := "rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ -"
	// fen6 := "r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - -"

	for i := 0; i < b.N; i++ {
		// cb.FENSet("k7/pppppppp/8/8/8/8/PPPPPPPP/7K w - -")
		// cb.FENSet("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -")
		cb.FENSet(fen1)
		cb.Perft(n)
	}
}

func BenchmarkBishopPerft(b *testing.B) {
	// this benchmark will specifically test the
	// ability to detect the movement of bishops

	// create and init a chessboard
	var cb board.ChessBoard
	cb.InitBoard()
	n := 4

	fen := "2b1kb2/8/8/8/8/8/8/2B1KB2 w - -"	

	for i := 0; i < b.N; i++ {
		cb.FENSet(fen)
		cb.Perft(n)
	}
}

func BenchmarkRookPerft(b *testing.B) {
	// this benchmark will specifically test the
	// ability to detect the movement of rooks

	// create and init a chessboard
	var cb board.ChessBoard
	cb.InitBoard()
	n := 4

	fen := "1r2k1r1/8/8/8/8/8/8/R3K2R w - -"

	for i := 0; i < b.N; i++ {
		cb.FENSet(fen)
		cb.Perft(n)
	}
}

func BenchmarkQueenPerft(b *testing.B) {
	// this benchmark will specifically test the
	// ability to detect the movement of queens

	// create and init a chessboard
	var cb board.ChessBoard
	cb.InitBoard()
	n := 4

	fen := "2q1k3/8/8/8/8/8/8/3QK3 w - -"

	for i := 0; i < b.N; i++ {
		cb.FENSet(fen)
		cb.Perft(n)
	}
}

func BenchmarkKnightPerft(b *testing.B) {
	// this benchmark will specifically test the
	// ability to detect the movement of knights

	// create and init a chessboard
	var cb board.ChessBoard
	cb.InitBoard()
	n := 4

	fen := "1n2k1n1/8/8/8/8/8/8/1N2K1N1 w - -"

	for i := 0; i < b.N; i++ {
		cb.FENSet(fen)
		cb.Perft(n)
	}
}

func BenchmarkPawnPerft(b *testing.B) {
	// this benchmark will specifically test the
	// ability to detect the movement of pawns

	// create and init a chessboard
	var cb board.ChessBoard
	cb.InitBoard()
	n := 4

	fen := "4k3/pppppppp/8/8/8/8/PPPPPPPP/4K3 w - -"

	for i := 0; i < b.N; i++ {
		cb.FENSet(fen)
		cb.Perft(n)
	}
}