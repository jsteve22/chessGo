package board

func pawnMove(cb *ChessBoard, p Piece) {
	// this function will calculate all of the moves a pawn
	// can make on the board. Any possible moves will be appended
	// to ChessBoard.moves array
	var forward int8
	var unMoved bool
	var nMove Move

	// determine which direction to move a pawn, up for white/down for black
	// also determine if current pawn has moved previously by determining
	// if it is still on it's starting rank. 1 for white and 6 for black
	if p.color == 0 {
		forward = 8
		unMoved = (p.pos >> 3) == 1
	} else {
		forward = -8
		unMoved = (p.pos >> 3) == 6
	}

	posMoves := make([]Move, 0)
	cb.inCheck()

	// if nothing is in front of pawn, add move forward once
	// if unmoved check if can push twice
	up := p.pos + forward
	if cb.board[up] == nil {
		// make move and check if king is still in check
		//cb.moves = append(cb.moves, Move{p.pos, up})
		nMove.start = p.pos
		nMove.end = up
		posMoves = append(posMoves, nMove)
		if unMoved {
			if cb.board[p.pos+(2*forward)] == nil {
				//cb.moves = append(cb.moves, Move{p.pos, p.pos + (2 * forward)})
				nMove.start = p.pos
				nMove.end = p.pos + (2 * forward)
				posMoves = append(posMoves, nMove)
			}
		}
	}

	// if not on left side, check if can take piece to the left
	if (p.pos & 7) != 0 {
		if cb.board[p.pos+forward-1] != nil {
			if cb.board[p.pos+forward-1].color != p.color {
				//cb.moves = append(cb.moves, Move{p.pos, p.pos + forward - 1})
				nMove.start = p.pos
				nMove.end = p.pos + forward - 1
				posMoves = append(posMoves, nMove)
			}
		}
		if cb.enpas == p.pos+forward-1 {
			//cb.moves = append(cb.moves, Move{p.pos, p.pos + forward - 1})
			nMove.start = p.pos
			nMove.end = p.pos + forward - 1
			posMoves = append(posMoves, nMove)
		}
	}

	// if not on right side, check if can take piece to the right
	if (p.pos & 7) != 7 {
		if cb.board[p.pos+forward+1] != nil {
			if cb.board[p.pos+forward+1].color != p.color {
				//cb.moves = append(cb.moves, Move{p.pos, p.pos + forward + 1})
				nMove.start = p.pos
				nMove.end = p.pos + forward + 1
				posMoves = append(posMoves, nMove)
			}
		}
		if cb.enpas == p.pos+forward+1 {
			//cb.moves = append(cb.moves, Move{p.pos, p.pos + forward + 1})
			nMove.start = p.pos
			nMove.end = p.pos + forward + 1
			posMoves = append(posMoves, nMove)
		}
	}

	// go through posMoves and check if any of the moves would stop
	// check and add those to cb.moves
	// only check if it will prevent check if king already in check
	if cb.check || true {
		var resetEnpas int8
		for _, m := range posMoves {
			resetEnpas = cb.enpas
			cb.makeMove(m)
			cb.inCheck()
			if !cb.check {
				cb.moves = append(cb.moves, m)
			}
			cb.enpas = resetEnpas
			cb.undoMove(m)
			cb.inCheck()
		}
	} else {
		cb.moves = append(cb.moves, posMoves...)
	}
}

func knightMove(cb *ChessBoard, p Piece) {

	vlong := make([]int8, 0)
	vshort := make([]int8, 0)
	hlong := make([]int8, 0)
	hshort := make([]int8, 0)

	file := p.pos & 7
	rank := (p.pos & 56) >> 3
	var nMove Move

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

	posMoves := make([]Move, 0)
	cb.inCheck()

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
	if cb.check || true {
		var resetEnpas int8
		for _, m := range posMoves {
			resetEnpas = cb.enpas
			cb.makeMove(m)
			cb.inCheck()
			if !cb.check {
				cb.moves = append(cb.moves, m)
			}
			cb.enpas = resetEnpas
			cb.undoMove(m)
			cb.inCheck()
		}
	} else {
		cb.moves = append(cb.moves, posMoves...)
	}
}

func bishopMove(cb *ChessBoard, p Piece) {
	var rank int8
	var file int8
	var pos int8
	var nMove Move

	file = p.pos & 7
	rank = (p.pos & 56) >> 3

	posMoves := make([]Move, 0)
	cb.inCheck()

	// check top right
	for {
		file++
		rank++
		if rank == 8 || file == 8 {
			break
		}
		pos = (rank << 3) + file
		// hit a piece
		if cb.board[pos] != nil {
			// other side's piece
			if cb.board[pos].color != p.color {
				nMove.start = p.pos
				nMove.end = pos
				posMoves = append(posMoves, nMove)
				//cb.moves = append(cb.moves, Move{p.pos, pos})
			}
			break
		}
		nMove.start = p.pos
		nMove.end = pos
		posMoves = append(posMoves, nMove)
		//cb.moves = append(cb.moves, Move{p.pos, pos})
	}

	file = p.pos & 7
	rank = (p.pos & 56) >> 3
	// check top left
	for {
		file--
		rank++
		if rank == 8 || file == -1 {
			break
		}
		pos = (rank << 3) + file
		// hit a piece
		if cb.board[pos] != nil {
			// other side's piece
			if cb.board[pos].color != p.color {
				nMove.start = p.pos
				nMove.end = pos
				posMoves = append(posMoves, nMove)
				//cb.moves = append(cb.moves, Move{p.pos, pos})
			}
			break
		}
		nMove.start = p.pos
		nMove.end = pos
		posMoves = append(posMoves, nMove)
		//cb.moves = append(cb.moves, Move{p.pos, pos})
	}

	file = p.pos & 7
	rank = (p.pos & 56) >> 3
	// check bottom left
	for {
		file--
		rank--
		if rank == -1 || file == -1 {
			break
		}
		pos = (rank << 3) + file
		// hit a piece
		if cb.board[pos] != nil {
			// other side's piece
			if cb.board[pos].color != p.color {
				nMove.start = p.pos
				nMove.end = pos
				posMoves = append(posMoves, nMove)
				//cb.moves = append(cb.moves, Move{p.pos, pos})
			}
			break
		}
		nMove.start = p.pos
		nMove.end = pos
		posMoves = append(posMoves, nMove)
		//cb.moves = append(cb.moves, Move{p.pos, pos})
	}

	file = p.pos & 7
	rank = (p.pos & 56) >> 3
	// check bottom right
	for {
		file++
		rank--
		if rank == -1 || file == 8 {
			break
		}
		pos = (rank << 3) + file
		// hit a piece
		if cb.board[pos] != nil {
			// other side's piece
			if cb.board[pos].color != p.color {
				nMove.start = p.pos
				nMove.end = pos
				posMoves = append(posMoves, nMove)
				//cb.moves = append(cb.moves, Move{p.pos, pos})
			}
			break
		}
		nMove.start = p.pos
		nMove.end = pos
		posMoves = append(posMoves, nMove)
		//cb.moves = append(cb.moves, Move{p.pos, pos})
	}

	// go through posMoves and check if any of the moves would stop
	// check and add those to cb.moves
	// only check if it will prevent check if king already in check
	if cb.check || true {
		var resetEnpas int8
		for _, m := range posMoves {
			resetEnpas = cb.enpas
			cb.makeMove(m)
			cb.inCheck()
			if !cb.check {
				cb.moves = append(cb.moves, m)
			}
			cb.enpas = resetEnpas
			cb.undoMove(m)
			cb.inCheck()
		}
	} else {
		cb.moves = append(cb.moves, posMoves...)
	}
}

func rookMove(cb *ChessBoard, p Piece) {
	var rank int8
	var file int8
	var pos int8
	var nMove Move

	posMoves := make([]Move, 0)
	cb.inCheck()

	file = p.pos & 7
	rank = (p.pos & 56) >> 3
	// check right
	for {
		file++
		if file == 8 {
			break
		}
		pos = (rank << 3) + file
		// hit a piece
		if cb.board[pos] != nil {
			// other side's piece
			if cb.board[pos].color != p.color {
				nMove.start = p.pos
				nMove.end = pos
				posMoves = append(posMoves, nMove)
				//cb.moves = append(cb.moves, Move{p.pos, pos})
			}
			break
		}
		nMove.start = p.pos
		nMove.end = pos
		posMoves = append(posMoves, nMove)
		//cb.moves = append(cb.moves, Move{p.pos, pos})
	}

	file = p.pos & 7
	rank = (p.pos & 56) >> 3
	// check left
	for {
		file--
		if file == -1 {
			break
		}
		pos = (rank << 3) + file
		// hit a piece
		if cb.board[pos] != nil {
			// other side's piece
			if cb.board[pos].color != p.color {
				nMove.start = p.pos
				nMove.end = pos
				posMoves = append(posMoves, nMove)
				//cb.moves = append(cb.moves, Move{p.pos, pos})
			}
			break
		}
		nMove.start = p.pos
		nMove.end = pos
		posMoves = append(posMoves, nMove)
		//cb.moves = append(cb.moves, Move{p.pos, pos})
	}

	file = p.pos & 7
	rank = (p.pos & 56) >> 3
	// check up
	for {
		rank++
		if rank == 8 {
			break
		}
		pos = (rank << 3) + file
		// hit a piece
		if cb.board[pos] != nil {
			// other side's piece
			if cb.board[pos].color != p.color {
				nMove.start = p.pos
				nMove.end = pos
				posMoves = append(posMoves, nMove)
				//cb.moves = append(cb.moves, Move{p.pos, pos})
			}
			break
		}
		nMove.start = p.pos
		nMove.end = pos
		posMoves = append(posMoves, nMove)
		//cb.moves = append(cb.moves, Move{p.pos, pos})
	}

	file = p.pos & 7
	rank = (p.pos & 56) >> 3
	// check down
	for {
		rank--
		if rank == -1 {
			break
		}
		pos = (rank << 3) + file
		// hit a piece
		if cb.board[pos] != nil {
			// other side's piece
			if cb.board[pos].color != p.color {
				nMove.start = p.pos
				nMove.end = pos
				posMoves = append(posMoves, nMove)
				//cb.moves = append(cb.moves, Move{p.pos, pos})
			}
			break
		}
		nMove.start = p.pos
		nMove.end = pos
		posMoves = append(posMoves, nMove)
		//cb.moves = append(cb.moves, Move{p.pos, pos})
	}

	// go through posMoves and check if any of the moves would stop
	// check and add those to cb.moves
	// only check if it will prevent check if king already in check
	if cb.check || true {
		var resetEnpas int8
		for _, m := range posMoves {
			resetEnpas = cb.enpas
			cb.makeMove(m)
			cb.inCheck()
			if !cb.check {
				cb.moves = append(cb.moves, m)
			}
			cb.enpas = resetEnpas
			cb.undoMove(m)
			cb.inCheck()
		}
	} else {
		cb.moves = append(cb.moves, posMoves...)
	}
}

func kingMove(cb *ChessBoard, p Piece) {
	var file int8
	var rank int8
	var next int8
	var nMove Move

	posMoves := make([]Move, 0)
	cb.inCheck()

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
			nMove.start = p.pos
			nMove.end = p.pos + 2
			posMoves = append(posMoves, nMove)
			//cb.moves = append(cb.moves, Move{p.pos, p.pos + 2})
		}
	}

	if cb.castle[(2*p.color)+1] && !cb.check {
		if cb.board[p.pos-1] == nil && cb.board[p.pos-2] == nil && cb.board[p.pos-3] == nil {
			nMove.start = p.pos
			nMove.end = p.pos - 2
			posMoves = append(posMoves, nMove)
			//cb.moves = append(cb.moves, Move{p.pos, p.pos - 2})
		}
	}

	// go through posMoves and check if any of the moves would stop
	// check and add those to cb.moves
	// only check if it will prevent check if king already in check
	if cb.check || true {
		var resetEnpas int8
		for _, m := range posMoves {
			resetEnpas = cb.enpas
			cb.makeMove(m)
			cb.inCheck()
			if !cb.check {
				cb.moves = append(cb.moves, m)
			}
			cb.enpas = resetEnpas
			cb.undoMove(m)
			cb.inCheck()
		}
	} else {
		cb.moves = append(cb.moves, posMoves...)
	}
}