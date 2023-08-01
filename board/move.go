package board

import "fmt"

type Move struct {
	start uint8
	end   uint8
	castle bool
	promotion uint8
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

type JsonMove struct {
	Start uint8
	End   uint8
	Castle bool
	Promotion uint8
}

func GetJsonMove(move Move) JsonMove {
	return JsonMove{ Start: move.start, End: move.end, Castle: move.castle, Promotion: move.promotion }
}

func GetJsonMoves(moves []Move) []JsonMove {
	jsons := make([]JsonMove, 0, len(moves))
	for _, move := range moves {
		jsons = append(jsons, GetJsonMove(move))
	}
	return jsons
}