package board

func PawnGeneratePseudoLegalMoves(pawn Piece, game Game) []Move {
	var moves []Move

	lowerMask := uint8(8 - 1)
	// upperMask := uint8(64 - 1 - lowerMask)

	rank := pawn.pos >> 3
	file := pawn.pos & lowerMask

	leftSideBoard := uint8(0)
	rightSideBoard := uint8(7)
	emptySquare := uint8(0)

	COLOR_BLACK := uint8(1)

	nextRank := uint8(rank - 1)
	enPasRank := uint8(6)
	doubleRank := uint8(rank - 2)
	if pawn.color == COLOR_BLACK {
		nextRank = uint8(rank + 1)
		enPasRank = uint8(1)
		doubleRank = uint8(rank + 2)
	}
	var doubleForwardSquare uint8
	var doubleForwardPiece uint8

	leftSquare := (nextRank << 3) + file - 1
	leftSquarePiece := game.board[leftSquare]
	if file == leftSideBoard { // check if the pawn is at the edge of the board
		goto skipleft
	}
	if (leftSquarePiece == emptySquare) && (leftSquare != uint8(game.enPassant)) { // check if there is a piece to take or if en passant is on the square
		goto skipleft
	}
	if (leftSquarePiece != emptySquare) && (leftSquarePiece>>3^pawn.color) != 0 { // check if piece is other color
		goto skipleft
	}

	moves = append(moves, Move{start: pawn.pos, end: leftSquare})
skipleft:

	rightSquare := (nextRank << 3) + file + 1
	rightSquarePiece := game.board[rightSquare]
	if file == rightSideBoard { // check if the pawn is at the edge of the board
		goto skipright
	}
	if (rightSquarePiece == emptySquare) && (rightSquare != uint8(game.enPassant)) { // check if there is a piece to take or if en passant is on the square
		goto skipright
	}
	if (rightSquarePiece != emptySquare) && (rightSquarePiece>>3^pawn.color) != 0 { // check if piece is other color
		goto skipright
	}

	moves = append(moves, Move{start: pawn.pos, end: rightSquare})
skipright:

	forwardSquare := (nextRank << 3) + file
	forwardSquarePiece := game.board[forwardSquare]
	if forwardSquarePiece != emptySquare {
		goto skipforward
	}

	// add pawn march forward move
	moves = append(moves, Move{start: pawn.pos, end: forwardSquare})

	if rank != enPasRank {
		goto skipenpas
	}

	doubleForwardSquare = (doubleRank << 3) + file
	doubleForwardPiece = game.board[doubleForwardSquare]
	if doubleForwardPiece != emptySquare {
		goto skipenpas
	}

	moves = append(moves, Move{start: pawn.pos, end: doubleForwardSquare})
skipenpas:
skipforward:

	return moves
}

/*
func PawnMove(cb *ChessBoard, p Piece) {
	// this function will calculate all of the moves a pawn
	// can make on the board. Any possible moves will be appended
	// to ChessBoard.moves array
	var forward int8
	var unMoved bool
	var promote bool
	var nMove Move
	color := cb.nextMove

	// determine which direction to move a pawn, up for white/down for black
	// also determine if current pawn has moved previously by determining
	// if it is still on it's starting rank. 1 for white and 6 for black
	if p.color == 0 {
		forward = 8
		unMoved = (p.pos >> 3) == 1
		promote = (p.pos >> 3) == 6
	} else {
		forward = -8
		unMoved = (p.pos >> 3) == 6
		promote = (p.pos >> 3) == 1
	}

	posMoves := make([]Move, 0, 4)
	cb.inCheck(color)

	// if nothing is in front of pawn, add move forward once
	// if unmoved check if can push twice
	up := p.pos + forward
	if cb.board[up] == nil {
		// make move and check if king is still in check
		//cb.moves = append(cb.moves, Move{p.pos, up})
		nMove.start = p.pos
		nMove.end = up
		if promote {
			for i := 2; i < 6; i++ {
				nMove.promotion = (int8)(i)
				posMoves = append(posMoves, nMove)
			}
		} else {
			posMoves = append(posMoves, nMove)
		}
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
				if promote {
					for i := 2; i < 6; i++ {
						nMove.promotion = (int8)(i)
						posMoves = append(posMoves, nMove)
					}
				} else {
					posMoves = append(posMoves, nMove)
				}
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
				if promote {
					for i := 2; i < 6; i++ {
						nMove.promotion = (int8)(i)
						posMoves = append(posMoves, nMove)
					}
				} else {
					posMoves = append(posMoves, nMove)
				}
			}
		}
		if cb.enpas == p.pos+forward+1 {
			//cb.moves = append(cb.moves, Move{p.pos, p.pos + forward + 1})
			nMove.start = p.pos
			nMove.end = p.pos + forward + 1
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

func PawnAttack(cb *ChessBoard, p Piece) {
	var forward int8
	if p.color == 0 {
		forward = 8
	} else {
		forward = -8
	}

	// if not on left side, check if can take piece to the left
	if (p.pos & 7) != 0 {
		cb.attackSquares = append(cb.attackSquares, p.pos+forward-1)
	}

	// if not on right side, check if can take piece to the right
	if (p.pos & 7) != 7 {
		cb.attackSquares = append(cb.attackSquares, p.pos+forward+1)
	}

}

*/