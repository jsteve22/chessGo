package board

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

func bishopAttack(cb *ChessBoard, p Piece) {
	var rank int8
	var file int8
	var pos int8

	file = p.pos & 7
	rank = (p.pos & 56) >> 3
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
				cb.attackSquares = append(cb.attackSquares, pos)
			}
			break
		}
		cb.attackSquares = append(cb.attackSquares, pos)
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
				cb.attackSquares = append(cb.attackSquares, pos)
			}
			break
		}
		cb.attackSquares = append(cb.attackSquares, pos)
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
				cb.attackSquares = append(cb.attackSquares, pos)
			}
			break
		}
		cb.attackSquares = append(cb.attackSquares, pos)
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
				cb.attackSquares = append(cb.attackSquares, pos)
			}
			break
		}
		cb.attackSquares = append(cb.attackSquares, pos)
	}
}
