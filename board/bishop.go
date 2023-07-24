package board

func BishopGeneratePseudoLegalMoves(bishop Piece, game Game) []Move {
	var moves []Move

	emptySquare := uint8(0)
	lowerMask := uint8(8 - 1)
	// upperMask := uint8(64 - 1 - lowerMask)

	rank := bishop.pos >> 3
	file := bishop.pos & lowerMask

	BOARD_SIZE := uint8(8)

	r := rank + 1
	f := file + 1
	for (r < BOARD_SIZE) && (f < BOARD_SIZE) {
		index := (r << 3) + f
		if game.board[index] != emptySquare {
			if (game.board[index]>>3)^bishop.color == 0 {
				break
			}
			moves = append(moves, Move{start: bishop.pos, end: index})
			break
		}
		moves = append(moves, Move{start: bishop.pos, end: index})
		r++
		f++
	}

	r = rank - 1
	f = file + 1
	for (r != 255) && (f < BOARD_SIZE) {
		index := (r << 3) + f
		if game.board[index] != emptySquare {
			if (game.board[index]>>3)^bishop.color == 0 {
				break
			}
			moves = append(moves, Move{start: bishop.pos, end: index})
			break
		}
		moves = append(moves, Move{start: bishop.pos, end: index})
		r--
		f++
	}

	r = rank - 1
	f = file - 1
	for (r != 255) && (f != 255) {
		index := (r << 3) + f
		if game.board[index] != emptySquare {
			if (game.board[index]>>3)^bishop.color == 0 {
				break
			}
			moves = append(moves, Move{start: bishop.pos, end: index})
			break
		}
		moves = append(moves, Move{start: bishop.pos, end: index})
		r--
		f--
	}

	r = rank + 1
	f = file - 1
	for (r < BOARD_SIZE) && (f != 255) {
		index := (r << 3) + f
		if game.board[index] != emptySquare {
			if (game.board[index]>>3)^bishop.color == 0 {
				break
			}
			moves = append(moves, Move{start: bishop.pos, end: index})
			break
		}
		moves = append(moves, Move{start: bishop.pos, end: index})
		r++
		f--
	}

	return moves
}

func BishopGenerateAttackSquaresBitboard(bishop Piece, game Game) uint64 {
	bitboard := uint64(0)

	emptySquare := uint8(0)
	lowerMask := uint8(8 - 1)
	// upperMask := uint8(64 - 1 - lowerMask)

	rank := bishop.pos >> 3
	file := bishop.pos & lowerMask

	BOARD_SIZE := uint8(8)

	r := rank + 1
	f := file + 1
	for (r < BOARD_SIZE) && (f < BOARD_SIZE) {
		index := (r << 3) + f
		bitboard |= (1 << index)
		if game.board[index] != emptySquare {
			break
		}
		r++
		f++
	}

	r = rank - 1
	f = file + 1
	for (r != 255) && (f < BOARD_SIZE) {
		index := (r << 3) + f
		bitboard |= (1 << index)
		if game.board[index] != emptySquare {
			break
		}
		r--
		f++
	}

	r = rank - 1
	f = file - 1
	for (r != 255) && (f != 255) {
		index := (r << 3) + f
		bitboard |= (1 << index)
		if game.board[index] != emptySquare {
			break
		}
		r--
		f--
	}

	r = rank + 1
	f = file - 1
	for (r < BOARD_SIZE) && (f != 255) {
		index := (r << 3) + f
		bitboard |= (1 << index)
		if game.board[index] != emptySquare {
			break
		}
		r++
		f--
	}

	return bitboard
}

/*
func BishopMove(cb *ChessBoard, p Piece) {
	var pos int8
	var nMove Move
	color := cb.nextMove


	posMoves := make([]Move, 0, 8)
	cb.inCheck(color)

	boardDist := [4]int8{9, 7, -9, -7}
	minDistIndex := (int)(p.pos) << 3
	minDistIndex += 4

	// loop through the min distances of the square
	for i := 0; i < 4; i++ {
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

func BishopAttack(cb *ChessBoard, p Piece) {
	// var rank int8
	// var file int8
	var pos int8

	boardDist := [4]int8{9, 7, -9, -7}
	minDistIndex := (int)(p.pos) << 3
	minDistIndex += 4

	// loop through the min distances of the square
	for i := 0; i < 4; i++ {
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