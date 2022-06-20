package board

func pawnMove(cb *ChessBoard, p Piece) {
	var forward int8
	var unMoved bool
	if p.color == 0 {
		forward = 8
		unMoved = (p.pos >> 3) == 1
	} else {
		forward = -8
		unMoved = (p.pos >> 3) == 6
	}

	// if nothing is in front of pawn, add move forward once
	// if unmoved check if can push twice
	up := p.pos + forward
	if cb.board[up] == nil {
		cb.moves = append(cb.moves, Move{p.pos, up})
		if unMoved {
			if cb.board[p.pos+(2*forward)] == nil {
				cb.moves = append(cb.moves, Move{p.pos, p.pos + (2 * forward)})
			}
		}
	}

	// if not on left side, check if can take piece to the left
	if (p.pos & 7) != 0 {
		if cb.board[p.pos+forward-1] != nil {
			if cb.board[p.pos+forward-1].color != p.color {
				cb.moves = append(cb.moves, Move{p.pos, p.pos + forward - 1})
			}
		}
		if cb.enpas == p.pos+forward-1 {
			cb.moves = append(cb.moves, Move{p.pos, p.pos + forward - 1})
		}
	}

	// if not on right side, check if can take piece to the right
	if (p.pos & 7) != 7 {
		if cb.board[p.pos+forward+1] != nil {
			if cb.board[p.pos+forward+1].color != p.color {
				cb.moves = append(cb.moves, Move{p.pos, p.pos + forward + 1})
			}
		}
		if cb.enpas == p.pos+forward+1 {
			cb.moves = append(cb.moves, Move{p.pos, p.pos + forward + 1})
		}
	}

}

func knightMove(cb *ChessBoard, p Piece) {

	vlong := make([]int8, 0)
	vshort := make([]int8, 0)
	hlong := make([]int8, 0)
	hshort := make([]int8, 0)

	file := p.pos & 7
	rank := (p.pos & 56) >> 3

	if file < 6 {
		hlong = append(hlong, 2)
	}
	if file < 7 {
		hshort = append(hshort, 1)
	}
	if file > 1 {
		hlong = append(hlong, -2)
	}
	if file > 0 {
		hshort = append(hshort, -1)
	}

	if rank < 6 {
		vlong = append(vlong, 16)
	}
	if rank < 7 {
		vshort = append(vshort, 8)
	}
	if rank > 1 {
		vlong = append(vlong, -16)
	}
	if rank > 0 {
		vshort = append(vshort, -8)
	}

	for _, i := range hlong {
		// check with long going horizontal and short going vertical
		for _, j := range vshort {
			if cb.board[p.pos+i+j] == nil {
				cb.moves = append(cb.moves, Move{p.pos, p.pos + i + j})
			} else if cb.board[p.pos+i+j].color != p.color {
				cb.moves = append(cb.moves, Move{p.pos, p.pos + i + j})
			}
		}
	}

	for _, i := range vlong {
		// check with long going vertical and short going horizontal
		for _, j := range hshort {
			if cb.board[p.pos+i+j] == nil {
				cb.moves = append(cb.moves, Move{p.pos, p.pos + i + j})
			} else if cb.board[p.pos+i+j].color != p.color {
				cb.moves = append(cb.moves, Move{p.pos, p.pos + i + j})
			}
		}
	}

}

func bishopMove(cb *ChessBoard, p Piece) {
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
				cb.moves = append(cb.moves, Move{p.pos, pos})
			}
			break
		}
		cb.moves = append(cb.moves, Move{p.pos, pos})
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
				cb.moves = append(cb.moves, Move{p.pos, pos})
			}
			break
		}
		cb.moves = append(cb.moves, Move{p.pos, pos})
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
				cb.moves = append(cb.moves, Move{p.pos, pos})
			}
			break
		}
		cb.moves = append(cb.moves, Move{p.pos, pos})
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
				cb.moves = append(cb.moves, Move{p.pos, pos})
			}
			break
		}
		cb.moves = append(cb.moves, Move{p.pos, pos})
	}
}

func rookMove(cb *ChessBoard, p Piece) {
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
				cb.moves = append(cb.moves, Move{p.pos, pos})
			}
			break
		}
		cb.moves = append(cb.moves, Move{p.pos, pos})
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
				cb.moves = append(cb.moves, Move{p.pos, pos})
			}
			break
		}
		cb.moves = append(cb.moves, Move{p.pos, pos})
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
				cb.moves = append(cb.moves, Move{p.pos, pos})
			}
			break
		}
		cb.moves = append(cb.moves, Move{p.pos, pos})
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
				cb.moves = append(cb.moves, Move{p.pos, pos})
			}
			break
		}
	}
}

func kingMove(cb *ChessBoard, p Piece) {
	var file int8
	var rank int8
	var next int8

	file = p.pos & 7
	rank = (p.pos & 56) >> 3

	horz := make([]int8, 1)
	vert := make([]int8, 1)

	if rank < 7 {
		vert = append(vert, 1)
	}
	if rank > 0 {
		vert = append(vert, -1)
	}

	if file < 7 {
		horz = append(horz, 1)
	}
	if file > 0 {
		horz = append(horz, -1)
	}

	// check immediate squares
	for _, i := range vert {
		for _, j := range horz {
			next = p.pos + j + (i << 3)

			if cb.board[next] == nil {
				cb.moves = append(cb.moves, Move{p.pos, next})
			} else if cb.board[next].color != p.color {
				cb.moves = append(cb.moves, Move{p.pos, next})
			}
		}
	}

	// check castling rights
	if cb.castle[2*p.color] {
		if cb.board[p.pos+1] == nil && cb.board[p.pos+2] == nil {
			cb.moves = append(cb.moves, Move{p.pos, p.pos + 2})
		}
	}

	if cb.castle[(2*p.color)+1] {
		if cb.board[p.pos-1] == nil && cb.board[p.pos-2] == nil && cb.board[p.pos-3] == nil {
			cb.moves = append(cb.moves, Move{p.pos, p.pos - 2})
		}
	}

}