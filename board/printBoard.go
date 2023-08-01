package board

import "fmt"


func PrintGame(game Game) {
	PrintBoard(game)
	fmt.Printf("Turn to move: %c\n", game.nextToPlay)
	PrintCastlingRights(game)
	PrintEnPassant(game)
	PrintHalfMoveClock(game)
	PrintFullMoveClock(game)
	PrintColorPieces(game)
}

func PrintBoard(game Game) {
	fmt.Printf("Printing Board\n")
	for ind := 0; ind < 7; ind++ {
		fmt.Printf("------")
	}
	fmt.Printf("\n|")
	for ind := 0; ind < 64; ind++ {
		pieceRep := PieceRepresentation(game.board[ind])
		fmt.Printf(" %s |", pieceRep)
		if ind%8 == 7 && ind != 63 {
			fmt.Printf("\n")
			for jnd := 0; jnd < 7; jnd++ {
				fmt.Printf("------")
			}
			fmt.Printf("\n|")
		}
	}
	fmt.Printf("\n")
	for ind := 0; ind < 7; ind++ {
		fmt.Printf("------")
	}
	fmt.Printf("\n")
}

func PieceRepresentation(piece uint8) string {
	switch piece {
	case 1:
		return "WP"
	case 2:
		return "WN"
	case 3:
		return "WB"
	case 4:
		return "WR"
	case 5:
		return "WQ"
	case 6:
		return "WK"
	case 9:
		return "BP"
	case 10:
		return "BN"
	case 11:
		return "BB"
	case 12:
		return "BR"
	case 13:
		return "BQ"
	case 14:
		return "BK"
	}
	return "  "
}


func PrintCastlingRights(game Game) {
	WK, WQ, BK, BQ := game.castlingRights[0], game.castlingRights[1], game.castlingRights[2], game.castlingRights[3] 

	fmt.Printf("Castling Rights: ")

	if ((!WK) && (!WQ) && (!BK) && (!BQ)) {
		fmt.Printf("-\n")
		return
	}

	if (WK) {
		fmt.Printf("WK ")
	}

	if (WQ) {
		fmt.Printf("WQ ")
	}

	if (BK) {
		fmt.Printf("BK ")
	}
	
	if (BQ) {
		fmt.Printf("BQ ")
	}

	fmt.Printf("\n")
}

func PrintEnPassant(game Game) {
	if (game.enPassant == -1) {
		fmt.Printf("En Passant: -\n");
		return
	}
	chessNotation, err := BoardIndexToChessNotation(uint8(game.enPassant))
	if (err != nil) {
		fmt.Printf("%e\n", err)
		fmt.Printf("enPassant: %d\n", game.enPassant)
	}
	fmt.Printf("En Passant: %s\n", chessNotation)
}

func PrintHalfMoveClock(game Game) {
	fmt.Printf("Half Move Clock: %d\n", game.halfMoveClock)
}

func PrintFullMoveClock(game Game) {
	fmt.Printf("Full Move Clock: %d\n", game.fullMoveClock)
}

func PrintColorPieces(game Game) {
	emptySquare := uint8(0)
	WHITE := uint8(0)
	BLACK := uint8(0)
	fmt.Printf("White Side: ")
	for _, square := range game.board {
		if (square != emptySquare) && (square>>3)==WHITE {
			pieceRep := PieceRepresentation(square)
			fmt.Printf("%s ", pieceRep )
		}
	}
	fmt.Printf("\n")
	fmt.Printf("            ")
	for index, square := range game.board {
		if (square != emptySquare) && (square>>3)==WHITE {
			chessNotation, _ := BoardIndexToChessNotation( uint8(index) )
			fmt.Printf("%s ", chessNotation )
		}
	}
	fmt.Printf("\n")
	fmt.Printf("\n")

	fmt.Printf("Black Side: ")
	for _, square := range game.board {
		if (square != emptySquare) && (square>>3)==BLACK {
			pieceRep := PieceRepresentation(square)
			fmt.Printf("%s ", pieceRep )
		}
	}
	fmt.Printf("\n")
	fmt.Printf("            ")
	for index, square := range game.board {
		if (square != emptySquare) && (square>>3)==BLACK {
			chessNotation, _ := BoardIndexToChessNotation( uint8(index) )
			fmt.Printf("%s ", chessNotation )
		}
	}
	fmt.Printf("\n")
}

/*
import "fmt"

func (b *ChessBoard) PrintBoard() {
	var p rune
	var rank int = 7
	var file int = 0
	var i int
	printHorizontalLine()
	for {
		if rank < 0 {
			break
		}
		i = file + (rank << 3)
		if b.board[i] == nil {
			p = ' '
		} else {
			p = b.board[i].rep
		}
		fmt.Printf("| %v ", (string)(p))
		file++
		if file == 8 {
			fmt.Printf("|\n")
			printHorizontalLine()
			file = 0
			rank--
		}
	}
}

func printHorizontalLine() {
	fmt.Printf("---------------------------------\n")
}

func (cb *ChessBoard) PrintMoves() {
	fmt.Printf("enpas = (%v,%v)\n", (cb.enpas&56)>>3, cb.enpas&7)
	fmt.Printf("%v\n", cb.nextMove)
	for _, m := range cb.moves {
		sFile := m.start & 7
		sRank := (m.start & 56) >> 3
		eFile := m.end & 7
		eRank := (m.end & 56) >> 3

		fmt.Printf("[%v] (%v%v) -> (%v%v)\n", (string)(cb.board[m.start].rep), (string)('A'+sFile), sRank+1, (string)('A'+eFile), eRank+1)
	}
}

func (cb *ChessBoard) PrintPieces() {
	// this function will print all of the pieces in the game
	for _, p := range cb.white {
		file := p.pos & 7
		rank := (p.pos & 56) >> 3

		fmt.Printf("[%v] %v%v \tAlive: %v\n", (string)(p.rep), (string)('A'+file), rank+1, p.alive)
	}
	fmt.Printf("\n")

	for _, p := range cb.black {
		file := p.pos & 7
		rank := (p.pos & 56) >> 3

		fmt.Printf("[%v] %v%v \tAlive: %v\n", (string)(p.rep), (string)('A'+file), rank+1, p.alive)
	}

}

func (cb *ChessBoard) PrintPrevMoves() {
	var sRank, sFile, eRank, eFile int8
	var capRep string
	for i, m := range cb.prevMoves {
		sRank = (m.start & 56) >> 3
		sFile = m.start & 7
		eRank = (m.end & 56) >> 3
		eFile = m.end & 7

		if m.pieceCaptured != nil {
			capRep = (string)(m.pieceCaptured.rep)
		} else {
			capRep = ""
		}

		fmt.Printf("Move %v: (%v) [%v] (%v%v) -> (%v%v) [%v]\n", i+1, m.color, (string)(m.pieceMoved.rep), (string)('A'+sFile), sRank+1, (string)('A'+eFile), eRank+1, capRep)
	}
}

func (cb *ChessBoard) PrintCastlePriv() {
	// this function will print the castling privileges of the board
	if cb.castle[0] {
		fmt.Printf("K")
	} else {
		fmt.Printf("-")
	}

	if cb.castle[1] {
		fmt.Printf("Q")
	} else {
		fmt.Printf("-")
	}

	if cb.castle[2] {
		fmt.Printf("k")
	} else {
		fmt.Printf("-")
	}

	if cb.castle[3] {
		fmt.Printf("q")
	} else {
		fmt.Printf("-")
	}

	fmt.Printf("\n")
}
*/