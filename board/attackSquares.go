package board

func (cb *ChessBoard) checkAttacks(color uint8) {
	cb.attackSquares = make([]int8, 0)

	var pieces *[16]Piece

	if color == 0 {
		pieces = &cb.black
		//forward = 8
	} else {
		pieces = &cb.white
		//forward = -8
	}

	for _, p := range pieces {
		if !p.alive {
			continue
		}
		switch p.piece {
		case 0:
			KingAttack(cb, p)
		case 1:
			PawnAttack(cb, p)
		case 2:
			KnightAttack(cb, p)
		case 3:
			BishopAttack(cb, p)
		case 4:
			RookAttack(cb, p)
		case 5:
			QueenAttack(cb, p)
		}
		/*
			if p.piece == 0 {
				KingAttack(cb, p)
			} else if p.piece == 1 {
				PawnAttack(cb, p)
				//cb.moves = append(cb.moves, Move{p.pos,p.pos+(uint8)(forward)})
			} else if p.piece == 2 {
				KnightAttack(cb, p)
			} else if p.piece == 3 {
				BishopAttack(cb, p)
			} else if p.piece == 4 {
				RookAttack(cb, p)
			} else if p.piece == 5 {
				QueenAttack(cb, p)
				// bishopAttack(cb, p)
				// rookAttack(cb, p)
			}
		*/
	}
}

func (cb *ChessBoard) PinPieces(color uint8) {
	// this function will add a list of all the pinned pieces
	// to a king, so that if a piece is pinned it's moved can
	// be checked if valid
	cb.pinned = make([]*Piece, 0)

	var pieces *[16]Piece

	if color == 0 {
		pieces = &cb.white
	} else {
		pieces = &cb.black
	}

	// figure out the king
	var king Piece = pieces[0]

	var pinP *Piece
	pinP = nil

	for i := -1; i < 2; i++ {
		for j := -8; j < 16; j += 8 {
			file := king.pos & 7
			rank := king.pos & 56
			pinP = nil
			for {
				file += (int8)(i)
				rank += (int8)(j)
				if rank > 56 || rank < 0 || file > 7 || file < 0 {
					break
				}
				// fmt.Printf("king vision %v\n", file+rank)
				// loop through each of the directions the king can go to
				//
				// check if square has a piece
				if cb.board[rank+file] != nil {
					// opposite color on same file/rank/diagonal as king
					if cb.board[rank+file].color != color {
						if cb.board[rank+file].piece == 1 {
							continue
						}
						if pinP != nil {
							cb.pinned = append(cb.pinned, pinP)
							pinP.pinned = true
						}
						break
					}
					// same color piece
					if pinP != nil {
						// two same color pieces in a row
						break
					}
					pinP = cb.board[rank+file]
				}
			}
		}
	}
}