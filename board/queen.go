package board

func getQueenBitmaps() []uint64 {
	return []uint64{
		0x81412111090503fe, 0x02824222120a07fd, 0x0404844424150efb, 0x08080888492a1cf7,
		0x10101011925438ef, 0x2020212224a870df, 0x404142444850e0bf, 0x8182848890a0c07f,
		0x412111090503fe03, 0x824222120a07fd07, 0x04844424150efb0e, 0x080888492a1cf71c,
		0x101011925438ef38, 0x20212224a870df70, 0x4142444850e0bfe0, 0x82848890a0c07fc0,
		0x2111090503fe0305, 0x4222120a07fd070a, 0x844424150efb0e15, 0x0888492a1cf71c2a,
		0x1011925438ef3854, 0x212224a870df70a8, 0x42444850e0bfe050, 0x848890a0c07fc0a0,
		0x11090503fe030509, 0x22120a07fd070a12, 0x4424150efb0e1524, 0x88492a1cf71c2a49,
		0x11925438ef385492, 0x2224a870df70a824, 0x444850e0bfe05048, 0x8890a0c07fc0a090,
		0x090503fe03050911, 0x120a07fd070a1222, 0x24150efb0e152444, 0x492a1cf71c2a4988,
		0x925438ef38549211, 0x24a870df70a82422, 0x4850e0bfe0504844, 0x90a0c07fc0a09088,
		0x0503fe0305091121, 0x0a07fd070a122242, 0x150efb0e15244484, 0x2a1cf71c2a498808,
		0x5438ef3854921110, 0xa870df70a8242221, 0x50e0bfe050484442, 0xa0c07fc0a0908884,
		0x03fe030509112141, 0x07fd070a12224282, 0x0efb0e1524448404, 0x1cf71c2a49880808,
		0x38ef385492111010, 0x70df70a824222120, 0xe0bfe05048444241, 0xc07fc0a090888482,
		0xfe03050911214181, 0xfd070a1222428202, 0xfb0e152444840404, 0xf71c2a4988080808,
		0xef38549211101010, 0xdf70a82422212020, 0xbfe0504844424140, 0x7fc0a09088848281,
	}
}

func QueenGeneratePseudoLegalMoves(queen Piece, game Game) []Move {
	moves := BishopGeneratePseudoLegalMoves(queen, game)
	moves = append(moves, RookGeneratePseudoLegalMoves(queen, game)...)
	return moves
}

func QueenGenerateAttackSquaresBitboard(queen uint8, game Game) uint64 {
	bitboard := uint64(0)

	bitboard |= BishopGenerateAttackSquaresBitboard(queen, game)
	bitboard |= RookGenerateAttackSquaresBitboard(queen, game)

	return bitboard
}

/*
func QueenMove(cb *ChessBoard, p Piece) {
	var pos int8
	var nMove Move
	color := cb.nextMove

	posMoves := make([]Move, 0, 32)

	boardDist := [8]int8{1, 8, -1, -8, 9, 7, -9, -7}
	minDistIndex := (int)(p.pos) << 3

	// loop through the min distances of the square
	for i := 0; i < 8; i++ {
		pos = p.pos
		for j := 0; j < (int)(cb.minDist[minDistIndex+i]); j++ {

			// pos = p.pos + (boardDist[i] * (int8)(j+1))
			pos += boardDist[i]
			// fmt.Printf("next square: %v\n",pos)

			if cb.board[pos] != nil {
				if cb.board[pos].color != p.color {
					nMove.start = p.pos
					nMove.end = pos
					posMoves = append(posMoves, nMove)
				}
				break
			}
			nMove.start = p.pos
			nMove.end = pos
			posMoves = append(posMoves, nMove)
		}
	}

	// go through pinned pieces and see if the piece is pinned to king
	pin := false
	for _, pinP := range cb.pinned {
		if pinP == cb.board[p.pos] {
			pin = true
			break
		}
	}

	// go through posMoves and check if any of the moves would stop
	// check and add those to cb.moves
	// only check if it will prevent check if king already in check
	if cb.check || pin {
		var resetEnpas int8
		for _, m := range posMoves {
			resetEnpas = cb.enpas
			cb.makeMove(m)
			cb.inCheck(color)
			if !cb.check {
				cb.moves = append(cb.moves, m)
			}
			cb.enpas = resetEnpas
			cb.undoMove(m)
			cb.inCheck(color)
		}
	} else {
		cb.moves = append(cb.moves, posMoves...)
	}
}

func QueenAttack(cb *ChessBoard, p Piece) {
	var pos int8

	boardDist := [8]int8{1, 8, -1, -8, 9, 7, -9, -7}
	minDistIndex := (int)(p.pos) << 3

	// loop through the min distances of the square
	for i := 0; i < 8; i++ {
		pos = p.pos
		for j := 0; j < (int)(cb.minDist[minDistIndex+i]); j++ {

			// pos = p.pos + (boardDist[i] * (int8)(j+1))
			pos += boardDist[i]
			// fmt.Printf("next square: %v\n",pos)

			if cb.board[pos] != nil {
				cb.attackSquares = append(cb.attackSquares, pos)
				break
			}
			cb.attackSquares = append(cb.attackSquares, pos)
		}
	}
}
*/