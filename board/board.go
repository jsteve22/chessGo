package board

import (
	"fmt"
)

type Move struct {
	start int8
	end int8
	pieceMoved *Piece			// pointer to piece that moved
	pieceCaptured *Piece	// pointer if a piece was taken
	capPos int8						// position where the piece was captured
	castle uint8					// 1=K; 2=Q; 3=k; 4=q
	color uint8 					// color of move
}

type Piece struct {
	alive bool
	pos int8
	piece int 	// king=0; pawn=1; knight=2; bishop=3; rook=4; queen=5 
	color uint8
	rep byte
}

type ChessBoard struct {
	board [64]*Piece
	white [16]Piece
	black [16]Piece

	moves []Move
	prevMoves []Move
	attackSquares []int8

	nextmove uint8
	turn uint
	revcnt uint
	check bool

	castle [4]bool	// KQkq
	enpas int8
}

func Aboard() {
	fmt.Printf("board test")
}

func (cb *ChessBoard) InitBoard() {
	
	cb.white = initSide(0)
	cb.black = initSide(1)	

	cb.castle[0] = true
	cb.castle[1] = true
	cb.castle[2] = true
	cb.castle[3] = true
	cb.enpas = -1

	cb.prevMoves = make([]Move, 0)
	cb.nextmove = 0
	cb.turn = 1
	cb.revcnt = 0

	cb.FENSet("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")

}

func (cb *ChessBoard) UpdateNextMove() {
	cb.nextmove = cb.nextmove ^ 1
	// cb.enpas = -1
	// fmt.Printf("lol\n")

	// go through each side and check if a pawn is on the back rank
	// so that it can promote to a new piece
	var rank int8
	// go through the white pieces
	for i,p := range cb.white {
		rank = ( p.pos & 56 ) >> 3
		if rank == 7 && p.piece == 1 && p.alive {
			pawnPromote( &cb.white[i] )
		}
	}

	// go through the black pieces
	for i,p := range cb.black {
		rank = ( p.pos & 56 ) >> 3
		if rank == 0 && p.piece == 1 && p.alive {
			pawnPromote( &cb.black[i] )
		}
	}
}

func pawnPromote(p *Piece) {
	// this function will update a pawn to a different piece
	// this will also update the rep of the piece

	// for now this will only upgrade pawns to queens

	// change piece
	p.piece = 5

	if p.color == 0 {
		p.rep = 'Q'
	} else if p.color == 1 {
		p.rep = 'q'
	}
}

func (cb *ChessBoard) UserMove(move string) bool {
	// this function will make a move on the board 
	sf := (int8)(move[0] - 'A')
	sr := (int8)(move[1] - '1')

	ef := (int8)(move[2] - 'A')
	er := (int8)(move[3] - '1')

	st := sf + (sr << 3)
	en := ef + (er << 3)

	for _,m := range cb.moves {
		var nMove Move
		nMove.start = st
		nMove.end = en
		if nMove.start == m.start && nMove.end == m.end {
			cb.permanentMove( nMove )
			return true
		}
	}
	return false
	/*
		// cb.makeMove( Move{st, en} )
		cb.permanentMove( Move{st, en} )
		return true
	*/
}

func (cb *ChessBoard) permanentMove (move Move) {
	// this function will make a move that won't be 
	// undone. This move will permanently delete pieces
	// if taken 

	st := move.start
	en := move.end 

	if cb.board[en] != nil {
		// this will move the piece completely off board
		cb.board[en].pos = -1

	}

	// this will check if piece moved is pawn
	if cb.board[st].piece == 1 {
		// this will check if the move was an enpas take
		if en == cb.enpas {
			stRank := ( st & 56 ) >> 3
			enFile := en & 7 
			enpasPawnPos := (stRank << 3) + enFile
			cb.board[ enpasPawnPos ].pos = -1;
		}
	}

	// call make move to perform the move since the piece
	// has already been removed from board
	cb.makeMove( move )
}

func (cb *ChessBoard) makeMove(move Move) {
	// this function will make a move and update prevMove
	st := move.start
	en := move.end

	// create appendMove which will be added to end of prevMove slice
	var appendMove Move
	appendMove.start = st
	appendMove.end = en
	appendMove.pieceMoved = cb.board[st]
	appendMove.color = cb.nextmove

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
	enpasMove := ( cb.board[en].piece == 1 && en == cb.enpas )

	cb.enpas = -1

	// check if move was a castle
	// check if piece moved was a king
	if cb.board[en].piece == 0 {
		stRank := ( st & 56 ) >> 3
		stFile := st & 7 
		enRank := ( en & 56 ) >> 3
		enFile := en & 7 

		// changed castling permissions
		cb.castle[2*cb.board[en].color] = false
		cb.castle[(2*cb.board[en].color)+1] = false

		// check if king is castling and to which side
		if stRank == enRank && enFile - stFile == 2 {
			// king is castling king side
			// move rook to F1/F8
			cb.board[st+1] = cb.board[en+1]
			cb.board[st+1].pos = st+1
			cb.board[en+1] = nil
			appendMove.castle = (2*cb.board[en].color) + 1
		} else if stRank == enRank && stFile - enFile == 2 {
			// king is castling queen side
			// move rook to D1/D8
			cb.board[st-1] = cb.board[stRank]
			cb.board[st-1].pos = st-1
			cb.board[stRank] = nil
			appendMove.castle = (2*cb.board[en].color) + 2
		}
	}

	// check if piece moved was a rook
	if cb.board[en].piece == 4 {
		// changed castling permissions
		cb.castle[2*cb.board[en].color] = false
		cb.castle[(2*cb.board[en].color)+1] = false
	}

	// check if piece moved was a pawn
	if cb.board[en].piece == 1 {
		stRank := ( st & 56 ) >> 3
		// stFile := st & 7 
		enRank := ( en & 56 ) >> 3
		enFile := en & 7 

		// check if the move was a double push
		if enRank - stRank == 2 || stRank - enRank == 2 {
			cb.enpas = st + (4 * (enRank - stRank) )
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
			appendMove.pieceCaptured = cb.board[ enpasPawnPos ]
			appendMove.capPos = enpasPawnPos
			cb.board[ enpasPawnPos ].alive = false
			cb.board[ enpasPawnPos ] = nil
			// cb.enpas = -1
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
	lastMove := cb.prevMoves[ len(cb.prevMoves)-1 ]

	if ( cb.nextmove == 0 ) {
		pieces = &cb.black
	} else {
		pieces = &cb.white
		// pieces = &cb.black
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
		stRank := ( st & 56 ) >> 3
		stFile := st & 7 
		enRank := ( en & 56 ) >> 3
		enFile := en & 7 
		// check if king is castling and to which side
		if stRank == enRank && enFile - stFile == 2 {
			// king is castling king side
			// move rook to F1/F8
			cb.board[en+1] = cb.board[st+1]
			cb.board[en+1].pos = en+1
			cb.board[st+1] = nil
		} else if stRank == enRank && stFile - enFile == 2 {
			// king is castling queen side
			// move rook to D1/D8
			cb.board[stRank] = cb.board[st-1]
			cb.board[stRank].pos = stRank
			cb.board[st-1] = nil
		}
	}

	// check if piece moved was a pawn
	if cb.board[st].piece == 1 {
		stRank := ( st & 56 ) >> 3
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

func (cb *ChessBoard) CheckMate() bool {
	// this function will return checkmate if the current
	// side has no more moves to play

	return len(cb.moves) == 0
}

func (cb *ChessBoard) checkAttacks() {
	cb.attackSquares = make([]int8, 0)

	var pieces *[16]Piece

	if ( cb.nextmove == 0 ) {
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
			kingAttack( cb, p )
		} else if p.piece == 1 {
			pawnAttack( cb, p )
			//cb.moves = append(cb.moves, Move{p.pos,p.pos+(uint8)(forward)})
		} else if p.piece == 2 {
			knightAttack( cb, p )
		} else if p.piece == 3 {
			bishopAttack( cb, p )
		} else if p.piece == 4 {
			rookAttack( cb, p )
		} else if p.piece == 5 {
			bishopAttack( cb, p )
			rookAttack( cb, p )
		}
	}
}

func (cb *ChessBoard) inCheck() {
	// this function will check if the king of the current side is in check
	cb.check = false

	cb.checkAttacks()

	var pieces *[16]Piece

	if ( cb.nextmove == 0 ) {
		pieces = &cb.white
		//forward = 8
	} else {
		pieces = &cb.black
		//forward = -8
	}

	king := pieces[0]

	for _,attacked := range cb.attackSquares {
		if attacked == king.pos {
			cb.check = true
		}
	}
}

func (cb *ChessBoard) GenMoves() {

	cb.inCheck()

	cb.moves = make([]Move, 0)

	var pieces *[16]Piece
	//var forward int8

	if ( cb.nextmove == 0 ) {
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
			kingMove( cb, p )
		} else if p.piece == 1 {
			pawnMove( cb, p )
			//cb.moves = append(cb.moves, Move{p.pos,p.pos+(uint8)(forward)})
		} else if p.piece == 2 {
			knightMove( cb, p )
		} else if p.piece == 3 {
			bishopMove( cb, p )
		} else if p.piece == 4 {
			rookMove( cb, p )
		} else if p.piece == 5 {
			bishopMove( cb, p )
			rookMove( cb, p )
		}
	}
}

func (cb *ChessBoard) FENSet(str string) {
	// first kill everything on each side
	for i := 0; i < 16; i++ {
		cb.white[i].alive = false
		cb.black[i].alive = false
	}

	// and clear the board
	for i := 0; i < 64; i++ {
		cb.board[i] = nil
	}

	// set starting rank and file
	rank := 7
	file := 0
	var pos int
	var emp int

	var addP *Piece

	// loop through the string
	for _, c := range str {
		// if c == '/', starting filling next rank
		if c == '/' {
			rank--
			file = 0
			continue
		}

		emp = (int)(c - '0')

		if emp > 8 || emp < 1 {
			// calculate the position in board
			pos = file + ( rank << 3 )

			// find piece to add to the board
			// check both sides to see if the
			// piece can be found first
			for i := 0; i < 16; i++ {
				if cb.white[i].rep == (byte)(c) && !cb.white[i].alive {
					cb.white[i].alive = true
					addP = &cb.white[i]
					break
				}
				if cb.black[i].rep == (byte)(c) && !cb.black[i].alive {
					cb.black[i].alive = true
					addP = &cb.black[i]
					break
				}
			}

			cb.board[pos] = addP
			addP.pos = (int8)(pos)
			file++
		} else {
			// increment the file by the numeric value in string
			file += emp
		}
		
	}

}

func (cb *ChessBoard) Perft(depth int) uint64 {
	// this function will perform perft function to go to
	// all moves in chess board

	var nodes uint64
	var resetEnpas int8

	if ( depth == 0 ) {
		return 1
	}

	cb.GenMoves()
	cpyMoves := make([]Move, len( cb.moves ) )
	copy( cpyMoves, cb.moves )

	for _,m := range cpyMoves {
		resetEnpas = cb.enpas
		cb.makeMove(m)
		cb.nextmove = 1 ^ cb.nextmove
		nodes += cb.Perft( depth - 1 )
		cb.nextmove = 1 ^ cb.nextmove
		cb.enpas = resetEnpas
		cb.undoMove(m)
	}
	return nodes
}

func initSide(color uint8) [16]Piece {
	var side [16]Piece
	
	name := [6]byte{ 'k', 'q', 'r', 'b', 'n', 'p' }
	if color == 0 {
		name = [6]byte{ 'K', 'Q', 'R', 'B', 'N', 'P' }
	}

	side[0] = Piece{ alive: true, piece: 0, color: color, rep: name[0] } // king
	side[1] = Piece{ alive: true, piece: 5, color: color, rep: name[1] } // queen
	side[2] = Piece{ alive: true, piece: 4, color: color, rep: name[2] } // rook
	side[3] = Piece{ alive: true, piece: 4, color: color, rep: name[2] } // rook
	side[4] = Piece{ alive: true, piece: 3, color: color, rep: name[3] } // bishop
	side[5] = Piece{ alive: true, piece: 3, color: color, rep: name[3] } // bishop
	side[6] = Piece{ alive: true, piece: 2, color: color, rep: name[4] } // knight
	side[7] = Piece{ alive: true, piece: 2, color: color, rep: name[4] } // knight
	side[8] = Piece{ alive: true, piece: 1, color: color, rep: name[5] } // pawn
	side[9] = Piece{ alive: true, piece: 1, color: color, rep: name[5] } // pawn
	side[10] = Piece{ alive: true, piece: 1, color: color, rep: name[5] } // pawn
	side[11] = Piece{ alive: true, piece: 1, color: color, rep: name[5] } // pawn
	side[12] = Piece{ alive: true, piece: 1, color: color, rep: name[5] } // pawn
	side[13] = Piece{ alive: true, piece: 1, color: color, rep: name[5] } // pawn
	side[14] = Piece{ alive: true, piece: 1, color: color, rep: name[5] } // pawn
	side[15] = Piece{ alive: true, piece: 1, color: color, rep: name[5] } // pawn

	return side
}

/*
func pawnMove(cb *ChessBoard, p Piece ) {
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
	up := p.pos+forward
	if cb.board[up] == nil {	
		cb.moves = append(cb.moves, Move{p.pos,up})
		if unMoved {
			if cb.board[p.pos+(2*forward)] == nil {
				cb.moves = append(cb.moves, Move{p.pos,p.pos+(2*forward)})
			}
		}
	}

	// if not on left side, check if can take piece to the left
	if (p.pos & 7) != 0 {
		if cb.board[p.pos + forward - 1] != nil {
			if cb.board[p.pos + forward - 1].color != p.color {
				cb.moves = append(cb.moves, Move{p.pos,p.pos + forward - 1}) 
			}
		}
	}

	// if not on right side, check if can take piece to the right
	if (p.pos & 7) != 7 {
		if cb.board[p.pos + forward + 1] != nil {
			if cb.board[p.pos + forward + 1].color != p.color {
				cb.moves = append(cb.moves, Move{p.pos,p.pos + forward + 1}) 
			}
		}
	}
}
*/