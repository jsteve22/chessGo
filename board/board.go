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

func (cb *ChessBoard) CheckMate() bool {
	// this function will return checkmate if the current
	// side has no more moves to play

	return len(cb.moves) == 0
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

	// QPerft 
	if ( depth == 1 ) {
		return (uint64)(len( cb.moves ))
	}

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
