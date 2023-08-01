package board_test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/jsteve22/chessGo/board"
)

// these tests will test the accuracy of the
// engine's current perft evaluation compared
// to the evaluations on https://www.chessprogramming.org/Perft_Results

func TestPerft1(t *testing.T) {
	// use Fen representation from website
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -"
	// This is the precalculated node count from chessprogramming.org
	nodes := []uint64{1, 20, 400, 8902, 197281, 4865609} // , 119060324} // , 3195901860}
	PerftTest(t, fen, nodes)
}
func TestPerft2(t *testing.T) {
	// use Fen representation from website
	fen := "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - "
	// This is the precalculated node count from chessprogramming.org
	nodes := []uint64{1, 48, 2039, 97862, 4085603} // , 193690690} // , 8031647685}
	PerftTest(t, fen, nodes)
}

func TestPerft3(t *testing.T) {
	// initiliaze chess board
	// use Fen representation from website
	fen := "8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - -"
	// This is the precalculated node count from chessprogramming.org
	nodes := []uint64{1, 14, 191, 2812, 43238, 674624}
	PerftTest(t, fen, nodes)
}

func TestPerft4(t *testing.T) {
	// use Fen representation from website
	fen := "r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq -"
	// This is the precalculated node count from chessprogramming.org
	nodes := []uint64{1, 6, 264, 9467, 422333, 15833292}
	PerftTest(t, fen, nodes)
}

func TestPerft5(t *testing.T) {
	// use Fen representation from website
	fen := "rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ -"
	// This is the precalculated node count from chessprogramming.org
	nodes := []uint64{1, 44, 1486, 62379, 2103487}
	PerftTest(t, fen, nodes)
}

func TestPerft6(t *testing.T) {
	// use Fen representation from website
	fen := "r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - -"
	// This is the precalculated node count from chessprogramming.org
	nodes := []uint64{1, 46, 2079, 89890, 3894594}
	PerftTest(t, fen, nodes)
}

func TestPerftPromotion(t *testing.T) {
	// use FEN from http://www.rocechess.ch/perft.html
	fen := "n1n5/PPPk4/8/8/8/8/4Kppp/5N1N b - - 0 1"
	// This is the precalculated node count from chessprogramming.org
	nodes := []uint64{1, 24, 496, 9483, 182838, 3605103, 71179139}
	PerftTest(t, fen, nodes)
}

func TestCompleteTest(t *testing.T) {
	readFile, err := os.Open("perftsuite.epd")

	if err != nil {
		t.Error(err)
	}
	
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lineSplit := strings.Split(line, ";")
		fen := lineSplit[0]
		nodes := []uint64{1}
		for i := 1; i < len(lineSplit); i++ {
			var node uint64
			fmt.Sscanf(lineSplit[i], "D%d %d", &i, &node)
			nodes = append(nodes, node)
			if (i == 3) {
				break
			}
		}

		PerftTest(t, fen, nodes)
	}

	readFile.Close()
}

func PerftTest(t *testing.T, fen string, nodes []uint64) {
	// initiliaze chess board
	game := board.LoadBoard(fen)
	for depth, node := range nodes {
		// set board and check perft with depth n
		moveCount := board.Perft(game, uint64(depth))

		if moveCount != node {
			t.Errorf("Perft(%v) = %v; want %v (%v difference)\n", depth, moveCount, node, int64(node-moveCount))
			break
		}
	}
}

/*
func BenchmarkPerftCalcultationsDepth_5(b *testing.B) {
	// this benchmark test will call Perft which
	// is a recursive function that continuously
	// generates moves. This will test how fast
	// move generation is for my chess program

	// create and init a chessboard
	var cb board.ChessBoard
	cb.InitBoard()
	n := 5

	// fen1 := "k7/8/8/8/3B4/8/8/7K w - -"
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
*/