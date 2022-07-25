package board

func BishopMove(cb *ChessBoard, p Piece) {
	var pos int8
	var nMove Move
	color := cb.nextMove

	/*
		var rank int8
		var file int8
		file = p.pos & 7
		rank = (p.pos & 56) >> 3
	*/

	posMoves := make([]Move, 0, 8)
	cb.inCheck(color)

	boardDist := [4]int8{9, 7, -9, -7}
	minDistIndex := (int)(p.pos) << 3
	minDistIndex += 4
	/*
		cb.PrintBoard()
		fmt.Printf("curr pos: %v\n", p.pos)
		fmt.Printf("minDistIndex: %v\n", minDistIndex)
		fmt.Printf("minDist[0]: %v\n", cb.minDist[minDistIndex+0])
		fmt.Printf("minDist[1]: %v\n", cb.minDist[minDistIndex+1])
		fmt.Printf("minDist[2]: %v\n", cb.minDist[minDistIndex+2])
		fmt.Printf("minDist[3]: %v\n", cb.minDist[minDistIndex+3])
	*/

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

	/*
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
	*/

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

	/*

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
	*/
}
