package board

func getKingBitmaps() []uint64 {
	return []uint64{
		0x0000000000000302, 0x0000000000000705, 0x0000000000000e0a, 0x0000000000001c14,
		0x0000000000003828, 0x0000000000007050, 0x000000000000e0a0, 0x000000000000c040,
		0x0000000000030203, 0x0000000000070507, 0x00000000000e0a0e, 0x00000000001c141c,
		0x0000000000382838, 0x0000000000705070, 0x0000000000e0a0e0, 0x0000000000c040c0,
		0x0000000003020300, 0x0000000007050700, 0x000000000e0a0e00, 0x000000001c141c00,
		0x0000000038283800, 0x0000000070507000, 0x00000000e0a0e000, 0x00000000c040c000,
		0x0000000302030000, 0x0000000705070000, 0x0000000e0a0e0000, 0x0000001c141c0000,
		0x0000003828380000, 0x0000007050700000, 0x000000e0a0e00000, 0x000000c040c00000,
		0x0000030203000000, 0x0000070507000000, 0x00000e0a0e000000, 0x00001c141c000000,
		0x0000382838000000, 0x0000705070000000, 0x0000e0a0e0000000, 0x0000c040c0000000,
		0x0003020300000000, 0x0007050700000000, 0x000e0a0e00000000, 0x001c141c00000000,
		0x0038283800000000, 0x0070507000000000, 0x00e0a0e000000000, 0x00c040c000000000,
		0x0302030000000000, 0x0705070000000000, 0x0e0a0e0000000000, 0x1c141c0000000000,
		0x3828380000000000, 0x7050700000000000, 0xe0a0e00000000000, 0xc040c00000000000,
		0x0203000000000000, 0x0507000000000000, 0x0a0e000000000000, 0x141c000000000000,
		0x2838000000000000, 0x5070000000000000, 0xa0e0000000000000, 0x40c0000000000000,
	}
}

func KingGeneratePseudoLegalMoves(king Piece, game Game) []Move {
	var moves []Move

	OTHER_COLOR := uint8(king.color ^ 1)

	KING_BITMAPS := getKingBitmaps()

	bitmap := KING_BITMAPS[king.pos]

	// get only squares that are not targeted
	otherColorAttacks := GenerateAttacks(game, OTHER_COLOR)
	bitmap = (bitmap ^ otherColorAttacks) & bitmap

	bitmap_pos := uint64(1)
	emptySquare := uint8(0)

	var potentialMoves []uint8

	// get potential moves for king
	for i := 0; i < 64; i++ {
		if (bitmap_pos & bitmap) > 0 {
			potentialMoves = append(potentialMoves, uint8(i))
		}

		bitmap_pos = bitmap_pos << 1
	}

	// get king moves with squares that are not occupied with same color
	for _, potMove := range potentialMoves {
		if game.board[potMove] != emptySquare && (game.board[potMove]>>3)^king.color == 0 {
			continue
		}
		moves = append(moves, Move{start: king.pos, end: potMove})
	}

	// check castling rights
	king_side := game.castlingRights[(king.color*2)+0]
	queen_side := game.castlingRights[(king.color*2)+1]

	if king_side {
		empty_king_side := (game.board[king.pos+1] == emptySquare) && (game.board[king.pos+2] == emptySquare)
		king_side_not_attacked := ((1<<(king.pos+1))|(1<<(king.pos+2)))&otherColorAttacks == 0
		if empty_king_side && king_side_not_attacked {
			moves = append(moves, Move{start: king.pos, end: king.pos + 2, castle: true})
		}
	}

	if queen_side {
		empty_queen_side := (game.board[king.pos-1] == emptySquare) && (game.board[king.pos-2] == emptySquare) && (game.board[king.pos-3] == emptySquare)
		queen_side_not_attacked := ((1<<(king.pos-1))|(1<<(king.pos-2)))&otherColorAttacks == 0
		if empty_queen_side && queen_side_not_attacked {
			moves = append(moves, Move{start: king.pos, end: king.pos - 2, castle: true})
		}
	}

	return moves
}

func KingGenerateAttackSquaresBitboard(king Piece) uint64 {
	KING_BITMAPS := getKingBitmaps()
	return KING_BITMAPS[king.pos]
}

/*
func KingMove(cb *ChessBoard, p Piece) {
	var file int8
	var rank int8
	var next int8
	var nMove Move
	color := cb.nextMove
	var skip bool

	posMoves := make([]Move, 0, 8)
	cb.inCheck(color)

	file = p.pos & 7
	rank = (p.pos & 56) >> 3

	horz := make([]int8, 1)
	vert := make([]int8, 1)

	if rank < 7 {
		vert = append(vert, 1)
	}
	if rank > 0 {
		vert = append(vert, -1)
	}

	if file < 7 {
		horz = append(horz, 1)
	}
	if file > 0 {
		horz = append(horz, -1)
	}

	// check immediate squares
	for _, i := range vert {
		for _, j := range horz {
			next = p.pos + j + (i << 3)
			skip = false
			for _, as := range cb.attackSquares {
				if as == next {
					skip = true
					break
				}
			}
			if skip {
				continue
			}

			if cb.board[next] == nil {
				nMove.start = p.pos
				nMove.end = next
				posMoves = append(posMoves, nMove)
				//cb.moves = append(cb.moves, Move{p.pos, next})
			} else if cb.board[next].color != p.color {
				nMove.start = p.pos
				nMove.end = next
				posMoves = append(posMoves, nMove)
				//cb.moves = append(cb.moves, Move{p.pos, next})
			}
		}
	}

	// check castling rights
	if cb.castle[2*p.color] && !cb.check {
		if cb.board[p.pos+1] == nil && cb.board[p.pos+2] == nil {
			skip = false
			for _, as := range cb.attackSquares {
				// fmt.Printf("%v\n", as)
				if as == p.pos+1 || as == p.pos+2 {
					skip = true
					break
				}
			}
			if !skip {
				nMove.start = p.pos
				nMove.end = p.pos + 2
				posMoves = append(posMoves, nMove)
				//cb.moves = append(cb.moves, Move{p.pos, p.pos + 2})
			}
		}
	}

	if cb.castle[(2*p.color)+1] && !cb.check {
		if cb.board[p.pos-1] == nil && cb.board[p.pos-2] == nil && cb.board[p.pos-3] == nil {
			skip = false
			for _, as := range cb.attackSquares {
				if as == p.pos-1 || as == p.pos-2 {
					skip = true
					break
				}
			}
			if !skip {
				nMove.start = p.pos
				nMove.end = p.pos - 2
				posMoves = append(posMoves, nMove)
				//cb.moves = append(cb.moves, Move{p.pos, p.pos - 2})
			}
		}
	}

	// go through posMoves and check if any of the moves would stop
	// check and add those to cb.moves
	// only check if it will prevent check if king already in check
	if cb.check || true {
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

func KingAttack(cb *ChessBoard, p Piece) {
	var file int8
	var rank int8
	var next int8

	file = p.pos & 7
	rank = (p.pos & 56) >> 3

	horz := make([]int8, 1)
	vert := make([]int8, 1)

	if rank < 7 {
		vert = append(vert, 1)
	}
	if rank > 0 {
		vert = append(vert, -1)
	}

	if file < 7 {
		horz = append(horz, 1)
	}
	if file > 0 {
		horz = append(horz, -1)
	}

	// check immediate squares
	for _, i := range vert {
		for _, j := range horz {
			next = p.pos + j + (i << 3)

			cb.attackSquares = append(cb.attackSquares, next)
			if cb.board[next] == nil {
				cb.attackSquares = append(cb.attackSquares, next)
			} else if cb.board[next].color != p.color {
				cb.attackSquares = append(cb.attackSquares, next)
			}
		}
	}

}
*/