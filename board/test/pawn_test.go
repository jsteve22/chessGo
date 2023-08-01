package board_test

import (
	"testing"

	"github.com/jsteve22/chessGo/board"
)

func TestPawnMove1(t *testing.T) {
	// Init Board for testing
	game := board.LoadBoard("8/8/8/8/8/8/4P3/8 w - -")

	// get moves
	moves := board.Perft(game, 1)
	expected := uint64(2)
	if moves != expected {
		t.Errorf("len(moves) = %v; want %v\n",moves, expected)
	}
}

func TestPawnMove2(t *testing.T) {
	// Init Board for testing
	game := board.LoadBoard("8/4p3/8/8/8/8/8/8 b - -")

	// get moves
	moves := board.Perft(game, 1)
	expected := uint64(2)
	if moves != expected {
		t.Errorf("len(moves) = %v; want %v\n",moves, expected)
	}
}

func TestPawnMove3(t *testing.T) {
	// Init Board for testing
	game := board.LoadBoard("8/8/8/8/8/4P3/8/8 w - -")

	// get moves
	moves := board.Perft(game, 1)
	expected := uint64(1)
	if moves != expected {
		t.Errorf("len(moves) = %v; want %v\n",moves, expected)
	}
}

func TestPawnMove4(t *testing.T) {
	// Init Board for testing
	game := board.LoadBoard("8/8/8/8/8/4p3/8/8 b - -")

	// get moves
	moves := board.Perft(game, 1)
	expected := uint64(1)
	if moves != expected {
		t.Errorf("len(moves) = %v; want %v\n",moves, expected)
	}
}

func TestPawnMove5(t *testing.T) {
	// Init Board for testing
	game := board.LoadBoard("8/8/5p2/4P3/8/8/8/8 w - -")

	// get moves
	moves := board.Perft(game, 1)
	expected := uint64(2)
	if moves != expected {
		t.Errorf("len(moves) = %v; want %v\n",moves, expected)
	}
	
}

func TestPawnMove6(t *testing.T) {
	// Init Board for testing
	game := board.LoadBoard("8/8/8/4Pp2/8/8/8/8 w - -")

	// get moves
	moves := board.Perft(game, 1)
	expected := uint64(1)
	if moves != expected {
		t.Errorf("len(moves) = %v; want %v\n",moves, expected)
	}
	
}
 
func TestPawnMove7(t *testing.T) {
	// Init Board for testing
	game := board.LoadBoard("8/8/8/4Pp2/8/8/8/8 w - f6")

	// get moves
	moves := board.Perft(game, 1)

	// board.PrintGame(game)
	// allMoves := board.GenerateMoves(game)
	// board.PrintMoves(game, allMoves)

	expected := uint64(2)
	if moves != expected {
		t.Errorf("len(moves) = %v; want %v\n",moves, expected)
	}
	
}

func TestPawnMove8(t *testing.T) {
	// Init Board for testing
	game := board.LoadBoard("8/8/3ppp2/4P3/8/8/8/8 w - -")

	// get moves
	moves := board.Perft(game, 1)
	expected := uint64(2)
	if moves != expected {
		t.Errorf("len(moves) = %v; want %v\n",moves, expected)
	}
	
}

func TestPawnMove9(t *testing.T) {
	// Init Board for testing
	game := board.LoadBoard("8/8/4p3/4P3/8/8/8/8 w - -")

	// get moves
	moves := board.Perft(game, 1)
	expected := uint64(0)
	if moves != expected {
		t.Errorf("len(moves) = %v; want %v\n",moves, expected)
	}
	
}

func TestPawnMove10(t *testing.T) {
	// Init Board for testing
	game := board.LoadBoard("8/8/p6p/P7/8/8/8/8 w - -")

	// get moves
	moves := board.Perft(game, 1)
	expected := uint64(0)
	if moves != expected {
		t.Errorf("len(moves) = %v; want %v\n",moves, expected)
	}
	
}

func TestPawnMove11(t *testing.T) {
	// Init Board for testing
	game := board.LoadBoard("7k/8/8/KPp4r/8/8/8/8 w - C6")

	// get moves
	moves := board.Perft(game, 1)
	expected := uint64(4)
	if moves != expected {
		t.Errorf("len(moves) = %v; want %v\n",moves, expected)
	}
	
}

func TestPawnMove12(t *testing.T) {
	// Init Board for testing
	game := board.LoadBoard("K6k/8/8/1Pp4r/8/8/8/8 w - c6")

	// get moves
	moves := board.Perft(game, 1)
	expected := uint64(5)
	if moves != expected {
		t.Errorf("len(moves) = %v; want %v\n",moves, expected)
	}

}
