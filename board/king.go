package board

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