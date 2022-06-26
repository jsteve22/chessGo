package board

func (cb *ChessBoard) makeMove(move Move) {
	// this function will make a move and update prevMove
	st := move.start
	en := move.end

	// create appendMove which will be added to end of prevMove slice
	var appendMove Move
	appendMove.start = st
	appendMove.end = en
	appendMove.pieceMoved = cb.board[st]
	appendMove.color = cb.nextMove

	if cb.board[en] != nil {
		cb.board[en].alive = false
		appendMove.pieceCaptured = cb.board[en]
		appendMove.capPos = en
	}

	// this is the naive movement, for simple movement
	cb.board[en] = cb.board[st]
	cb.board[st] = nil
	cb.board[en].pos = en

	// check if move was enpassant take
	enpasMove := (cb.board[en].piece == 1 && en == cb.enpas)

	cb.enpas = -1

	// check if move was a castle
	// check if piece moved was a king
	if cb.board[en].piece == 0 {
		stRank := st & 56
		stFile := st & 7
		enRank := en & 56
		enFile := en & 7

		// changed castling permissions
		cb.castle[2*cb.board[en].color] = false
		cb.castle[(2*cb.board[en].color)+1] = false

		// check if king is castling and to which side
		if stRank == enRank && enFile-stFile == 2 {
			// king is castling king side
			// move rook to F1/F8
			appendMove.castle = (2 * cb.board[en].color) + 1
			// appendMove.pieceCaptured = cb.board[en+1]
			// appendMove.capPos = (stRank << 3) + 7
			cb.board[st+1] = cb.board[en+1]
			cb.board[st+1].pos = st + 1
			cb.board[en+1] = nil
		} else if stRank == enRank && stFile-enFile == 2 {
			// king is castling queen side
			// move rook to D1/D8
			appendMove.castle = (2 * cb.board[en].color) + 2
			// appendMove.pieceCaptured = cb.board[stRank << 3]
			// appendMove.capPos = stRank << 3
			cb.board[st-1] = cb.board[stRank]
			cb.board[st-1].pos = st - 1
			cb.board[stRank] = nil
		}
	}

	// check if piece moved was a rook
	if cb.board[en].piece == 4 {
		// changed castling permissions
		cb.castle[2*appendMove.color] = false
		cb.castle[(2*appendMove.color)+1] = false
	}

	// check if piece moved was a pawn
	if cb.board[en].piece == 1 {
		stRank := (st & 56) >> 3
		// stFile := st & 7
		enRank := (en & 56) >> 3
		enFile := en & 7

		// check if the move was a double push
		if enRank-stRank == 2 || stRank-enRank == 2 {
			cb.enpas = st + (4 * (enRank - stRank))
			// cb.enpas = -1
		}

		// check if move was an enpas take
		if enpasMove {
			enpasPawnPos := (stRank << 3) + enFile
			/*
				fmt.Printf("ENPAS MADNESS (%v%v)\n",(string)('A'+enFile),stRank+1)
				cb.PrintBoard()
				cb.PrintPieces()
				fmt.Printf("\n\n")
			*/
			// fmt.Printf("%v", enpasPawnPos )
			appendMove.pieceCaptured = cb.board[enpasPawnPos]
			appendMove.capPos = enpasPawnPos
			cb.board[enpasPawnPos].alive = false
			cb.board[enpasPawnPos] = nil
			// cb.enpas = -1
		}

		// promote pawn to new piece
		if move.promotion != 0 {
			appendMove.pieceMoved.piece = move.promotion
			appendMove.promotion = move.promotion
			switch move.promotion {
			case 2:
				appendMove.pieceMoved.rep = (byte)('N' + (appendMove.pieceMoved.color * 32))
			case 3:
				appendMove.pieceMoved.rep = (byte)('B' + (appendMove.pieceMoved.color * 32))
			case 4:
				appendMove.pieceMoved.rep = (byte)('R' + (appendMove.pieceMoved.color * 32))
			case 5:
				appendMove.pieceMoved.rep = (byte)('Q' + (appendMove.pieceMoved.color * 32))
			}
		}
	}

	// add new move to slice
	cb.prevMoves = append(cb.prevMoves, appendMove)
}

func (cb *ChessBoard) undoMove(move Move) {
	// undo a given move on the board. This will
	// only undo the previous move that was on the
	// board. If provided a different move than the
	// immediate previous one, then could cause issues

	var pieces *[16]Piece
	lastMove := cb.prevMoves[len(cb.prevMoves)-1]

	if cb.nextMove == 0 {
		pieces = &cb.black
	} else {
		pieces = &cb.white
		// pieces = &cb.black
	}

	if lastMove.promotion != 0 {
		lastMove.pieceMoved.piece = 1
		lastMove.pieceMoved.color = (byte)('P' + (32 * lastMove.pieceMoved.color))
	}

	st := move.start
	en := move.end

	// reset piece on board to previous position
	cb.board[move.start] = cb.board[move.end]
	cb.board[move.start].pos = move.start
	cb.board[move.end] = nil

	// loop through other side's pieces and bring
	// back a piece if it was taken on previous turn
	for i, p := range pieces {
		if p.pos == lastMove.capPos {
			// p.alive = true
			// cb.board[move.end] = &p
			pieces[i].alive = true
			// cb.board[lastMove.capPos] = lastMove.pieceCaptured
			cb.board[lastMove.capPos] = &pieces[i]
			cb.prevMoves[len(cb.prevMoves)-1] = Move{}
			cb.prevMoves = cb.prevMoves[:len(cb.prevMoves)-1]
			return
		}
	}

	// check if move was a castle
	// check if piece moved was a king
	if cb.board[st].piece == 0 {
		stRank := (st & 56)
		stFile := st & 7
		enRank := (en & 56)
		enFile := en & 7

		// check if last move was a castle move
		if lastMove.castle != 0 {
			if lastMove.castle < 3 { // white castle
				cb.castle[0] = true
				cb.castle[1] = true
			} else { // black castle
				cb.castle[2] = true
				cb.castle[3] = true
			}
		}

		// check if king is castling and to which side
		// if stRank == enRank && enFile - stFile == 2 {
		if lastMove.castle == 1 || lastMove.castle == 3 {
			// king is castling king side
			// move rook to F1/F8
			// lastMove.pieceCaptured.pos = lastMove.capPos
			// cb.board[ ( stRank << 3) + 7 ] = lastMove.pieceCaptured
			// cb.board[ ( stRank << 3) + 5 ] = nil
			cb.board[en+1] = cb.board[st+1]
			cb.board[en+1].pos = en + 1
			cb.board[st+1] = nil
		} else if stRank == enRank && stFile-enFile == 2 {
			// king is castling queen side
			// move rook to D1/D8
			// lastMove.pieceCaptured.pos = lastMove.capPos
			// cb.board[ stRank << 3] = lastMove.pieceCaptured
			// cb.board[ ( stRank << 3) + 4 ] = nil
			cb.board[stRank] = cb.board[st-1]
			cb.board[stRank].pos = stRank
			cb.board[st-1] = nil
		}
	}

	// check if piece moved was a pawn
	if cb.board[st].piece == 1 {
		stRank := (st & 56) >> 3
		// stFile := st & 7
		// enRank := ( en & 56 ) >> 3
		enFile := en & 7

		// check if move was an enpas take
		if en == cb.enpas {
			// find which piece was captured and bring back to life
			enpasPawnPos := (stRank << 3) + enFile
			for i, p := range pieces {
				if p.pos == (enpasPawnPos) {
					pieces[i].alive = true
					cb.board[enpasPawnPos] = &pieces[i]
					/*
						fmt.Printf("ENPAS UNDO (%v%v)\n",(string)('A'+enFile),stRank+1)
						cb.PrintBoard()
						cb.PrintPieces()
						fmt.Printf("\n\n")
					*/
				}
			}
		}
	}

	// delete last element of prevMoves
	cb.prevMoves[len(cb.prevMoves)-1] = Move{}
	cb.prevMoves = cb.prevMoves[:len(cb.prevMoves)-1]
}

func (cb *ChessBoard) GenMoves() {

	cb.inCheck()

	cb.moves = make([]Move, 0)

	var pieces *[16]Piece
	//var forward int8

	if cb.nextMove == 0 {
		pieces = &cb.white
		//forward = 8
	} else {
		pieces = &cb.black
		//forward = -8
	}

	for _, p := range pieces {
		if !p.alive {
			continue
		}
		if p.piece == 0 {
			KingMove(cb, p)
		} else if p.piece == 1 {
			PawnMove(cb, p)
			//cb.moves = append(cb.moves, Move{p.pos,p.pos+(uint8)(forward)})
		} else if p.piece == 2 {
			KnightMove(cb, p)
		} else if p.piece == 3 {
			BishopMove(cb, p)
		} else if p.piece == 4 {
			RookMove(cb, p)
		} else if p.piece == 5 {
			QueenMove(cb, p)
			// bishopMove(cb, p)
			// rookMove(cb, p)
		}
	}
}
