package board

import "fmt"

type Move struct {
	start uint8
	end   uint8
}

func (m *Move) GetStart() uint8 {
	return m.start
}

func (m *Move) GetEnd() uint8 {
	return m.end
}

func PrintMoves(game Game, moves []Move) {
	fmt.Printf("{")
	moveCutoff := 3
	for ind, move := range moves {
		printMove(game, move)
		fmt.Printf(", ")
		if (ind % moveCutoff == moveCutoff-1) && (ind != len(moves)-1) {
			fmt.Printf("\n ")
		}
	}
	fmt.Printf("}\n")
}

func printMove(game Game, move Move) {
	startNotation, err := BoardIndexToChessNotation(move.start)
	if err != nil {
		fmt.Printf("%e", err)
	}

	endNotation, err := BoardIndexToChessNotation(move.end)
	if err != nil {
		fmt.Printf("%e", err)
	}

	pieceRep := PieceRepresentation(game.board[move.start])

	fmt.Printf("(%s: %s -> %s)", pieceRep, startNotation, endNotation)
}