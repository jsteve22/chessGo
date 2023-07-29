package board

func getBishopBitmaps() []uint64 {
	return []uint64{
		0x8040201008040200, 0x0080402010080500, 0x0000804020110a00, 0x0000008041221400,
		0x0000000182442800, 0x0000010204885000, 0x000102040810a000, 0x0102040810204000,
		0x4020100804020002, 0x8040201008050005, 0x00804020110a000a, 0x0000804122140014,
		0x0000018244280028, 0x0001020488500050, 0x0102040810a000a0, 0x0204081020400040,
		0x2010080402000204, 0x4020100805000508, 0x804020110a000a11, 0x0080412214001422,
		0x0001824428002844, 0x0102048850005088, 0x02040810a000a010, 0x0408102040004020,
		0x1008040200020408, 0x2010080500050810, 0x4020110a000a1120, 0x8041221400142241,
		0x0182442800284482, 0x0204885000508804, 0x040810a000a01008, 0x0810204000402010,
		0x0804020002040810, 0x1008050005081020, 0x20110a000a112040, 0x4122140014224180,
		0x8244280028448201, 0x0488500050880402, 0x0810a000a0100804, 0x1020400040201008,
		0x0402000204081020, 0x0805000508102040, 0x110a000a11204080, 0x2214001422418000,
		0x4428002844820100, 0x8850005088040201, 0x10a000a010080402, 0x2040004020100804,
		0x0200020408102040, 0x0500050810204080, 0x0a000a1120408000, 0x1400142241800000,
		0x2800284482010000, 0x5000508804020100, 0xa000a01008040201, 0x4000402010080402,
		0x0002040810204080, 0x0005081020408000, 0x000a112040800000, 0x0014224180000000,
		0x0028448201000000, 0x0050880402010000, 0x00a0100804020100, 0x0040201008040201,
	}
}

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