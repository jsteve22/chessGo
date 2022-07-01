package board

func RookMove(cb *ChessBoard, p Piece) {
	var rank int8
	var file int8
	var pos int8
	var nMove Move
	color := cb.nextMove

	posMoves := make([]Move, 0)
	cb.inCheck(color)

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

func RookAttack(cb *ChessBoard, p Piece) {
	var rank int8
	var file int8
	var pos int8

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
				cb.attackSquares = append(cb.attackSquares, pos)
			}
			break
		}
		cb.attackSquares = append(cb.attackSquares, pos)
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
				cb.attackSquares = append(cb.attackSquares, pos)
			}
			break
		}
		cb.attackSquares = append(cb.attackSquares, pos)
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
				cb.attackSquares = append(cb.attackSquares, pos)
			}
			break
		}
		cb.attackSquares = append(cb.attackSquares, pos)
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
				cb.attackSquares = append(cb.attackSquares, pos)
			}
			break
		}
	}
}
