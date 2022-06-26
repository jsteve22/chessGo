package board

func (cb *ChessBoard) checkAttacks() {
	cb.attackSquares = make([]int8, 0)

	var pieces *[16]Piece

	if cb.nextMove == 0 {
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
	}
}
