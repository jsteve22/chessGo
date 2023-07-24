package board

func RookGeneratePseudoLegalMoves(rook Piece, game Game) []Move {
	var moves []Move

	emptySquare := uint8(0)
	lowerMask := uint8(8 - 1)
	// upperMask := uint8(64 - 1 - lowerMask)

	rank := rook.pos >> 3
	file := rook.pos & lowerMask

	BOARD_SIZE := uint8(8)

	r := rank + 1
	f := file
	for r < BOARD_SIZE {
		index := (r << 3) + f
		if (game.board[index] != emptySquare) && (game.board[index]>>3)^rook.color == 0 {
			break
		}
		moves = append(moves, Move{start: rook.pos, end: index})
		r++
	}

	r = rank - 1
	f = file
	for r != 255 {
		index := (r << 3) + f
		if (game.board[index] != emptySquare) && (game.board[index]>>3)^rook.color == 0 {
			break
		}
		moves = append(moves, Move{start: rook.pos, end: index})
		r--
	}

	r = rank
	f = file + 1
	for f < BOARD_SIZE {
		index := (r << 3) + f
		if (game.board[index] != emptySquare) && (game.board[index]>>3)^rook.color == 0 {
			break
		}
		moves = append(moves, Move{start: rook.pos, end: index})
		f++
	}

	r = rank
	f = file - 1
	for f != 255 {
		index := (r << 3) + f
		if (game.board[index] != emptySquare) && (game.board[index]>>3)^rook.color == 0 {
			break
		}
		moves = append(moves, Move{start: rook.pos, end: index})
		f--
	}

	return moves
}

/*
func RookMove(cb *ChessBoard, p Piece) {
	var pos int8
	var nMove Move
	color := cb.nextMove

	posMoves := make([]Move, 0, 8)
	cb.inCheck(color)

	boardDist := [4]int8{1, 8, -1, -8}
	minDistIndex := (int)(p.pos) << 3

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

func RookAttack(cb *ChessBoard, p Piece) {
	// var rank int8
	// var file int8
	var pos int8

	boardDist := [4]int8{1, 8, -1, -8}
	minDistIndex := (int)(p.pos) << 3

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