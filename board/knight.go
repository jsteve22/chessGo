package board

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
