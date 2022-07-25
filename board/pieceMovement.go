package board

func (cb *ChessBoard) makeMove(move Move) {
	// this function will make a move across the board and update
	// the move counter. Each move made should be reversible in
	// undoMove, even if it's a capture, a castle, or a promotion.

	// initialize the start and end of the given move
	st := move.start
	en := move.end

	// create appendMove which will be added to end of prevMove slice
	// use the appendMove to make the move across the board
	var appendMove Move
	appendMove.start = st
	appendMove.end = en
	appendMove.pieceMoved = cb.board[st]
	appendMove.color = cb.nextMove
	appendMove.prevEnpas = cb.enpas
	appendMove.prevCastle[0] = cb.castle[0]
	appendMove.prevCastle[1] = cb.castle[1]
	appendMove.prevCastle[2] = cb.castle[2]
	appendMove.prevCastle[3] = cb.castle[3]
	appendMove.fullmove = cb.fullmove
	appendMove.halfmove = cb.halfmove

	// increment clock
	cb.fullmove = cb.fullmove + 1
	cb.halfmove = cb.halfmove + 1

	// if the end position is not nil, a piece is captured
	if cb.board[en] != nil {
		// set piece to dead
		cb.board[en].alive = false
		// add captured piece to move
		appendMove.pieceCaptured = cb.board[en]
		appendMove.capPos = en
		// reset halfmove counter
		cb.halfmove = 0
	}

	// this is the naive movement, for simple movement
	cb.board[en] = appendMove.pieceMoved // cb.board[st]
	appendMove.pieceMoved.pos = en
	cb.board[st] = nil

	// check if move was enpassant take and reset enpas
	enpasMove := (appendMove.pieceMoved.piece == 1 && en == cb.enpas)
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
		// check which side the rook was on before updating castling
		// check if on queen side
		if appendMove.start == (int8)(56*appendMove.color) {
			cb.castle[(2*appendMove.color)+1] = false
		}
		// check if on king side
		if appendMove.start == (int8)(56*appendMove.color+7) {
			cb.castle[2*appendMove.color] = false
		}
	}

	// if start or end of a move is on A1,A8,H1, or H8 get rid of castling
	if st == 0 || en == 0 {
		cb.castle[1] = false
	}

	if st == 7 || en == 7 {
		cb.castle[0] = false
	}

	if st == 56 || en == 56 {
		cb.castle[3] = false
	}

	if st == 63 || en == 63 {
		cb.castle[2] = false
	}

	// check if piece moved was a pawn
	if cb.board[en].piece == 1 {
		// reset halfmove counter
		cb.halfmove = 0

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

	// change turns on the board
	cb.nextMove = 1 ^ cb.nextMove

	// add new move to slice
	cb.prevMoves = append(cb.prevMoves, appendMove)
}

func (cb *ChessBoard) undoMove(move Move) {
	// undo a given move on the board. This will
	// only undo the previous move that was on the
	// board. If provided a different move than the
	// immediate previous one, then could cause issues

	lastMove := cb.prevMoves[len(cb.prevMoves)-1]

	// reset board positions with the move
	cb.nextMove = lastMove.color
	cb.enpas = lastMove.prevEnpas
	cb.fullmove = lastMove.fullmove
	cb.halfmove = lastMove.halfmove
	cb.castle[0] = lastMove.prevCastle[0]
	cb.castle[1] = lastMove.prevCastle[1]
	cb.castle[2] = lastMove.prevCastle[2]
	cb.castle[3] = lastMove.prevCastle[3]

	if lastMove.promotion != 0 {
		// fmt.Printf("1 lastMove: %v\n", *lastMove.pieceMoved)
		lastMove.pieceMoved.piece = 1
		lastMove.pieceMoved.rep = (byte)('P' + (32 * lastMove.pieceMoved.color))
		// fmt.Printf("2 lastMove: %v\n", *lastMove.pieceMoved)
		// fmt.Printf("\n")
	}

	st := move.start
	en := move.end

	// reset piece on board to previous position
	cb.board[move.start] = lastMove.pieceMoved
	lastMove.pieceMoved.pos = st
	cb.board[move.end] = nil

	// bring pack a piece if captured
	if lastMove.pieceCaptured != nil {
		lastMove.pieceCaptured.alive = true
		cb.board[lastMove.capPos] = lastMove.pieceCaptured
	}

	// check if move was a castle
	// check if piece moved was a king
	if lastMove.pieceMoved.piece == 0 {
		stRank := (st & 56)
		stFile := st & 7
		enRank := (en & 56)
		enFile := en & 7

		// check if king is castling and to which side
		// if stRank == enRank && enFile - stFile == 2 {
		if lastMove.castle == 1 || lastMove.castle == 3 {
			// king is castling king side
			// move rook to F1/F8
			cb.board[en+1] = cb.board[st+1]
			cb.board[en+1].pos = en + 1
			cb.board[st+1] = nil
		} else if stRank == enRank && stFile-enFile == 2 {
			// king is castling queen side
			// move rook to D1/D8
			cb.board[stRank] = cb.board[st-1]
			cb.board[stRank].pos = stRank
			cb.board[st-1] = nil
		}
	}

	// delete last element of prevMoves
	cb.prevMoves[len(cb.prevMoves)-1] = Move{}
	cb.prevMoves = cb.prevMoves[:len(cb.prevMoves)-1]
}

func (cb *ChessBoard) GenMoves() {

	cb.inCheck(cb.nextMove)
	cb.PinPieces(cb.nextMove)

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
		switch p.piece {
		case 0:
			KingMove(cb, p)
		case 1:
			PawnMove(cb, p)
		case 2:
			KnightMove(cb, p)
		case 3:
			BishopMove(cb, p)
		case 4:
			RookMove(cb, p)
		case 5:
			QueenMove(cb, p)
		}
		/*
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
		*/
	}
}
