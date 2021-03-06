package board

func QueenMove(cb *ChessBoard, p Piece) {
	/*
		BishopMove(cb, p)
		RookMove(cb, p)
	*/
	var pos int8
	var nMove Move
	color := cb.nextMove

	posMoves := make([]Move, 0)

	boardDist := [8]int8{1, 8, -1, -8, 9, 7, -9, -7}
	minDistIndex := (int)(p.pos) << 3

	// loop through the min distances of the square
	for i := 0; i < 8; i++ {
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

func QueenAttack(cb *ChessBoard, p Piece) {
	/*
		BishopAttack(cb, p)
		RookAttack(cb, p)
	*/
	var pos int8

	boardDist := [8]int8{1, 8, -1, -8, 9, 7, -9, -7}
	minDistIndex := (int)(p.pos) << 3

	// loop through the min distances of the square
	for i := 0; i < 8; i++ {
		pos = p.pos
		for j := 0; j < (int)(cb.minDist[minDistIndex+i]); j++ {

			// pos = p.pos + (boardDist[i] * (int8)(j+1))
			pos += boardDist[i]
			// fmt.Printf("next square: %v\n",pos)

			if cb.board[pos] != nil {
				if cb.board[pos].color != p.color {
					cb.attackSquares = append(cb.attackSquares, pos)
				}
				break
			}
			cb.attackSquares = append(cb.attackSquares, pos)
		}
	}
}