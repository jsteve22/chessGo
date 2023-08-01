package board_test

import (
	"testing"

	"github.com/jsteve22/chessGo/board"
)

func BenchmarkPerft(b *testing.B) {
	// this benchmark will specifically test the
	// ability to detect the movement of pawns

	// create and init a chessboard
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -"
	game := board.LoadBoard(fen)

	// This is the precalculated node count from chessprogramming.org
	// nodes := []uint64{1, 20, 400, 8902, 197281} // , 4865609, 119060324} // , 3195901860}
	node := uint64(197281)
	depth := uint64(4)

	for i := 0; i < b.N; i++ {
		// set board and check perft with depth n
		moveCount := board.Perft(game, depth)

		if moveCount != node {
			b.Errorf("Perft(%v) = %v; want %v (%v difference)\n", depth, moveCount, node, int64(node-moveCount))
			break
		}
	}
}