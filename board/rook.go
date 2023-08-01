package board

func getRookBitmaps() []uint64 {
	return []uint64{
		0x01010101010101fe, 0x02020202020202fd, 0x04040404040404fb, 0x08080808080808f7,
		0x10101010101010ef, 0x20202020202020df, 0x40404040404040bf, 0x808080808080807f,
		0x010101010101fe01, 0x020202020202fd02, 0x040404040404fb04, 0x080808080808f708,
		0x101010101010ef10, 0x202020202020df20, 0x404040404040bf40, 0x8080808080807f80,
		0x0101010101fe0101, 0x0202020202fd0202, 0x0404040404fb0404, 0x0808080808f70808,
		0x1010101010ef1010, 0x2020202020df2020, 0x4040404040bf4040, 0x80808080807f8080,
		0x01010101fe010101, 0x02020202fd020202, 0x04040404fb040404, 0x08080808f7080808,
		0x10101010ef101010, 0x20202020df202020, 0x40404040bf404040, 0x808080807f808080,
		0x010101fe01010101, 0x020202fd02020202, 0x040404fb04040404, 0x080808f708080808,
		0x101010ef10101010, 0x202020df20202020, 0x404040bf40404040, 0x8080807f80808080,
		0x0101fe0101010101, 0x0202fd0202020202, 0x0404fb0404040404, 0x0808f70808080808,
		0x1010ef1010101010, 0x2020df2020202020, 0x4040bf4040404040, 0x80807f8080808080,
		0x01fe010101010101, 0x02fd020202020202, 0x04fb040404040404, 0x08f7080808080808,
		0x10ef101010101010, 0x20df202020202020, 0x40bf404040404040, 0x807f808080808080,
		0xfe01010101010101, 0xfd02020202020202, 0xfb04040404040404, 0xf708080808080808,
		0xef10101010101010, 0xdf20202020202020, 0xbf40404040404040, 0x7f80808080808080,
	}
}

func RookGeneratePseudoLegalMoves(rook Piece, game Game) []Move {
	var moves []Move

	emptySquare := uint8(0)
	lowerMask := uint8(8 - 1)
	// upperMask := uint8(64 - 1 - lowerMask)

	rank := rook.pos >> 3
	file := rook.pos & lowerMask

	BOARD_SIZE := uint8(8)

	r := rank + 1
	f := file
	for r < BOARD_SIZE {
		index := (r << 3) + f
		if game.board[index] != emptySquare {
			if (game.board[index]>>3)^rook.color == 0 {
				break
			}
			moves = append(moves, Move{start: rook.pos, end: index})
			break
		}
		moves = append(moves, Move{start: rook.pos, end: index})
		r++
	}

	r = rank - 1
	f = file
	for r != 255 {
		index := (r << 3) + f
		if game.board[index] != emptySquare {
			if (game.board[index]>>3)^rook.color == 0 {
				break
			}
			moves = append(moves, Move{start: rook.pos, end: index})
			break
		}
		moves = append(moves, Move{start: rook.pos, end: index})
		r--
	}

	r = rank
	f = file + 1
	for f < BOARD_SIZE {
		index := (r << 3) + f
		if game.board[index] != emptySquare {
			if (game.board[index]>>3)^rook.color == 0 {
				break
			}
			moves = append(moves, Move{start: rook.pos, end: index})
			break
		}
		moves = append(moves, Move{start: rook.pos, end: index})
		f++
	}

	r = rank
	f = file - 1
	for f != 255 {
		index := (r << 3) + f
		if game.board[index] != emptySquare {
			if (game.board[index]>>3)^rook.color == 0 {
				break
			}
			moves = append(moves, Move{start: rook.pos, end: index})
			break
		}
		moves = append(moves, Move{start: rook.pos, end: index})
		f--
	}

	return moves
}

func RookGenerateAttackSquaresBitboard(rook uint8, game Game) uint64 {
	bitboard := uint64(0)

	emptySquare := uint8(0)
	lowerMask := uint8(8 - 1)
	// upperMask := uint8(64 - 1 - lowerMask)

	rank := rook >> 3
	file := rook & lowerMask

	BOARD_SIZE := uint8(8)

	r := rank + 1
	f := file
	for r < BOARD_SIZE {
		index := (r << 3) + f
		bitboard |= (1 << index)
		if game.board[index] != emptySquare {
			break
		}
		r++
	}

	r = rank - 1
	f = file
	for r != 255 {
		index := (r << 3) + f
		bitboard |= (1 << index)
		if game.board[index] != emptySquare {
			break
		}
		r--
	}

	r = rank
	f = file + 1
	for f < BOARD_SIZE {
		index := (r << 3) + f
		bitboard |= (1 << index)
		if game.board[index] != emptySquare {
			break
		}
		f++
	}

	r = rank
	f = file - 1
	for f != 255 {
		index := (r << 3) + f
		bitboard |= (1 << index)
		if game.board[index] != emptySquare {
			break
		}
		f--
	}

	return bitboard
}

/*
func RookMove(cb *ChessBoard, p Piece) {
	var pos int8
	var nMove Move
	color := cb.nextMove

	posMoves := make([]Move, 0, 8)
	cb.inCheck(color)

	boardDist := [4]int8{1, 8, -1, -8}
	minDistIndex := (int)(p.pos) << 3

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
		for _, m := range posMoves {
			cb.makeMove(m)
			cb.inCheck(color)
			if !cb.check {
				cb.moves = append(cb.moves, m)
			}
			cb.undoMove(m)
			cb.inCheck(color)
		}
	} else {
		cb.moves = append(cb.moves, posMoves...)
	}
}

func RookAttack(cb *ChessBoard, p Piece) {
	// var rank int8
	// var file int8
	var pos int8

	boardDist := [4]int8{1, 8, -1, -8}
	minDistIndex := (int)(p.pos) << 3

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