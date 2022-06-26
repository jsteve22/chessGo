package board

import "fmt"

func (b *ChessBoard) PrintBoard() {
	var p byte
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