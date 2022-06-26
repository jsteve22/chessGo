package board

func PawnMove(cb *ChessBoard, p Piece) {
	// this function will calculate all of the moves a pawn
	// can make on the board. Any possible moves will be appended
	// to ChessBoard.moves array
	var forward int8
	var unMoved bool
	var promote bool
	var nMove Move

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

func PawnAttack(cb *ChessBoard, p Piece) {
	var forward int8
	if p.color == 0 {
		forward = 8
	} else {
		forward = -8
	}

	// if not on left side, check if can take piece to the left
	if (p.pos & 7) != 0 {
		if cb.board[p.pos+forward-1] != nil {
			if cb.board[p.pos+forward-1].color != p.color {
				cb.attackSquares = append(cb.attackSquares, p.pos+forward-1)
			}
		}
		if cb.enpas == p.pos+forward-1 {
			cb.attackSquares = append(cb.attackSquares, p.pos+forward-1)
		}
	}

	// if not on right side, check if can take piece to the right
	if (p.pos & 7) != 7 {
		if cb.board[p.pos+forward+1] != nil {
			if cb.board[p.pos+forward+1].color != p.color {
				cb.attackSquares = append(cb.attackSquares, p.pos+forward+1)
			}
		}
		if cb.enpas == p.pos+forward+1 {
			cb.attackSquares = append(cb.attackSquares, p.pos+forward+1)
		}
	}

}
