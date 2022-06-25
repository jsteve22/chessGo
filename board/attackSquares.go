package board

func (cb *ChessBoard) checkAttacks() {
	cb.attackSquares = make([]int8, 0)

	var pieces *[16]Piece

	if cb.nextmove == 0 {
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
		if p.piece == 0 {
			kingAttack(cb, p)
		} else if p.piece == 1 {
			pawnAttack(cb, p)
			//cb.moves = append(cb.moves, Move{p.pos,p.pos+(uint8)(forward)})
		} else if p.piece == 2 {
			knightAttack(cb, p)
		} else if p.piece == 3 {
			bishopAttack(cb, p)
		} else if p.piece == 4 {
			rookAttack(cb, p)
		} else if p.piece == 5 {
			queenAttack(cb, p)
			// bishopAttack(cb, p)
			// rookAttack(cb, p)
		}
	}
}
