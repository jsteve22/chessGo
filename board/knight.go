package board

func getKnightBitmaps() []uint64 {
	return []uint64{
		0x0000000000020400, 0x0000000000050800, 0x00000000000a1100, 0x0000000000142200,
		0x0000000000284400, 0x0000000000508800, 0x0000000000a01000, 0x0000000000402000,
		0x0000000002040004, 0x0000000005080008, 0x000000000a110011, 0x0000000014220022,
		0x0000000028440044, 0x0000000050880088, 0x00000000a0100010, 0x0000000040200020,
		0x0000000204000402, 0x0000000508000805, 0x0000000a1100110a, 0x0000001422002214,
		0x0000002844004428, 0x0000005088008850, 0x000000a0100010a0, 0x0000004020002040,
		0x0000020400040200, 0x0000050800080500, 0x00000a1100110a00, 0x0000142200221400,
		0x0000284400442800, 0x0000508800885000, 0x0000a0100010a000, 0x0000402000204000,
		0x0002040004020000, 0x0005080008050000, 0x000a1100110a0000, 0x0014220022140000,
		0x0028440044280000, 0x0050880088500000, 0x00a0100010a00000, 0x0040200020400000,
		0x0204000402000000, 0x0508000805000000, 0x0a1100110a000000, 0x1422002214000000,
		0x2844004428000000, 0x5088008850000000, 0xa0100010a0000000, 0x4020002040000000,
		0x0400040200000000, 0x0800080500000000, 0x1100110a00000000, 0x2200221400000000,
		0x4400442800000000, 0x8800885000000000, 0x100010a000000000, 0x2000204000000000,
		0x0004020000000000, 0x0008050000000000, 0x00110a0000000000, 0x0022140000000000,
		0x0044280000000000, 0x0088500000000000, 0x0010a00000000000, 0x0020400000000000,
	}
}

func KnightGeneratePseudoLegalMoves(knight Piece, game Game) []Move {
	// var moves []Move
	moves := make([]Move, 0, 8)

	KNIGHT_BITMAPS := getKnightBitmaps()

	bitmap := KNIGHT_BITMAPS[knight.pos]
	bitmap_pos := uint64(1)
	emptySquare := uint8(0)

	for i := uint8(0); i < 64; i++ {
		if (bitmap_pos&bitmap) > 0 && (game.board[i] == emptySquare || ((game.board[i]>>3)^knight.color) != 0) {
			moves = append(moves, Move{start: knight.pos, end: i})
		}

		bitmap_pos = bitmap_pos << 1
	}

	return moves
}

func KnightGenerateAttackSquaresBitboard(knight uint8) uint64 {
	KNIGHT_BITMAPS := getKnightBitmaps()
	return KNIGHT_BITMAPS[knight]
}

/*
func KnightMove(cb *ChessBoard, p Piece) {

	vlong := make([]int8, 0, 4)
	vshort := make([]int8, 0, 4)
	hlong := make([]int8, 0, 4)
	hshort := make([]int8, 0, 4)

	file := p.pos & 7
	rank := (p.pos & 56) >> 3
	var nMove Move
	color := cb.nextMove

	// go through pinned pieces and see if the piece is pinned to king
	pin := false
	for _, pinP := range cb.pinned {
		if pinP == cb.board[p.pos] {
			pin = true
			return
		}
	}

	if file < 6 {
		hlong = append(hlong, 2)
	}
	if file < 7 {
		hshort = append(hshort, 1)
	}
	if file > 1 {
		hlong = append(hlong, -2)
	}
	if file > 0 {
		hshort = append(hshort, -1)
	}

	if rank < 6 {
		vlong = append(vlong, 16)
	}
	if rank < 7 {
		vshort = append(vshort, 8)
	}
	if rank > 1 {
		vlong = append(vlong, -16)
	}
	if rank > 0 {
		vshort = append(vshort, -8)
	}

	posMoves := make([]Move, 0, 8)
	cb.inCheck(color)

	for _, i := range hlong {
		// check with long going horizontal and short going vertical
		for _, j := range vshort {
			if cb.board[p.pos+i+j] == nil {
				nMove.start = p.pos
				nMove.end = p.pos + i + j
				posMoves = append(posMoves, nMove)
				//cb.moves = append(cb.moves, Move{p.pos, p.pos + i + j})
			} else if cb.board[p.pos+i+j].color != p.color {
				nMove.start = p.pos
				nMove.end = p.pos + i + j
				posMoves = append(posMoves, nMove)
				//cb.moves = append(cb.moves, Move{p.pos, p.pos + i + j})
			}
		}
	}

	for _, i := range vlong {
		// check with long going vertical and short going horizontal
		for _, j := range hshort {
			if cb.board[p.pos+i+j] == nil {
				nMove.start = p.pos
				nMove.end = p.pos + i + j
				posMoves = append(posMoves, nMove)
				//cb.moves = append(cb.moves, Move{p.pos, p.pos + i + j})
			} else if cb.board[p.pos+i+j].color != p.color {
				nMove.start = p.pos
				nMove.end = p.pos + i + j
				posMoves = append(posMoves, nMove)
				//cb.moves = append(cb.moves, Move{p.pos, p.pos + i + j})
			}
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

func KnightAttack(cb *ChessBoard, p Piece) {

	vlong := make([]int8, 0)
	vshort := make([]int8, 0)
	hlong := make([]int8, 0)
	hshort := make([]int8, 0)

	file := p.pos & 7
	rank := (p.pos & 56) >> 3

	if file < 6 {
		hlong = append(hlong, 2)
	}
	if file < 7 {
		hshort = append(hshort, 1)
	}
	if file > 1 {
		hlong = append(hlong, -2)
	}
	if file > 0 {
		hshort = append(hshort, -1)
	}

	if rank < 6 {
		vlong = append(vlong, 16)
	}
	if rank < 7 {
		vshort = append(vshort, 8)
	}
	if rank > 1 {
		vlong = append(vlong, -16)
	}
	if rank > 0 {
		vshort = append(vshort, -8)
	}

	for _, i := range hlong {
		// check with long going horizontal and short going vertical
		for _, j := range vshort {
			cb.attackSquares = append(cb.attackSquares, p.pos+i+j)
		}
	}

	for _, i := range vlong {
		// check with long going vertical and short going horizontal
		for _, j := range hshort {
			cb.attackSquares = append(cb.attackSquares, p.pos+i+j)
		}
	}

}
*/