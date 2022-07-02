package board

import (
	"fmt"
	"strings"
)

type ChessBoard struct {
	board [64]*Piece
	white [16]Piece
	black [16]Piece

	moves         []Move
	prevMoves     []Move
	attackSquares []int8
	pinned        []*Piece

	nextMove uint8
	turn     uint
	revcnt   uint
	check    bool

	castle [4]bool // KQkq
	enpas  int8

	minDist [512]int8
}

func (cb *ChessBoard) InitBoard() {

	cb.white = initSide(0)
	cb.black = initSide(1)

	cb.calcMinDist()

	/*
		cb.castle[0] = true
		cb.castle[1] = true
		cb.castle[2] = true
		cb.castle[3] = true
		cb.enpas = -1
	*/

	cb.prevMoves = make([]Move, 0)
	cb.nextMove = 0
	cb.turn = 1
	cb.revcnt = 0

	cb.FENSet("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -")

}

func (cb *ChessBoard) UpdateNextMove() {
	cb.nextMove = cb.nextMove ^ 1
	// cb.enpas = -1
	// fmt.Printf("lol\n")

	// go through each side and check if a pawn is on the back rank
	// so that it can promote to a new piece
	var rank int8
	// go through the white pieces
	for i, p := range cb.white {
		rank = (p.pos & 56) >> 3
		if rank == 7 && p.piece == 1 && p.alive {
			pawnPromote(&cb.white[i])
		}
	}

	// go through the black pieces
	for i, p := range cb.black {
		rank = (p.pos & 56) >> 3
		if rank == 0 && p.piece == 1 && p.alive {
			pawnPromote(&cb.black[i])
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

	for _, m := range cb.moves {
		var nMove Move
		nMove.start = st
		nMove.end = en
		if nMove.start == m.start && nMove.end == m.end {
			cb.permanentMove(nMove)
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

func (cb *ChessBoard) permanentMove(move Move) {
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
			stRank := (st & 56) >> 3
			enFile := en & 7
			enpasPawnPos := (stRank << 3) + enFile
			cb.board[enpasPawnPos].pos = -1
		}
	}

	// call make move to perform the move since the piece
	// has already been removed from board
	cb.makeMove(move)
}

func (cb *ChessBoard) CheckMate() bool {
	// this function will return checkmate if the current
	// side has no more moves to play

	return len(cb.moves) == 0
}

func (cb *ChessBoard) inCheck(color uint8) {
	// this function will check if the king of the current side is in check
	cb.check = false

	cb.checkAttacks(color)

	var pieces *[16]Piece

	if color == 0 {
		pieces = &cb.white
		//forward = 8
	} else {
		pieces = &cb.black
		//forward = -8
	}

	king := pieces[0]

	for _, attacked := range cb.attackSquares {
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

	// set enpas to -1
	cb.enpas = -1
	enpasCnt := 0

	// set castling privileges to false
	cb.castle[0] = false
	cb.castle[1] = false
	cb.castle[2] = false
	cb.castle[3] = false

	// set starting rank and file
	rank := 7
	file := 0
	var pos int
	var emp int

	var addP *Piece

	fenPart := 0 // know which part of FEN string we're reading in
	// 0=board; 1=turn; 2=castle; 3=enpas; 4=halfmove; 5=turn

	// loop through the string
	for _, c := range str {
		if c == ' ' {
			fenPart++
			continue
		}
		switch fenPart {
		case 0: // fill board
			// if c == '/', starting filling next rank
			if c == '/' {
				rank--
				file = 0
				continue
			}

			emp = (int)(c - '0')

			if emp > 8 || emp < 1 {
				// calculate the position in board
				pos = file + (rank << 3)

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
		case 1: // update turn
			if c == 'w' {
				cb.nextMove = 0
			} else {
				cb.nextMove = 1
			}
		case 2: // castling
			switch c {
			case 'K':
				cb.castle[0] = true
			case 'Q':
				cb.castle[1] = true
			case 'k':
				cb.castle[2] = true
			case 'q':
				cb.castle[3] = true
			default:
				continue
			}
		case 3: // enpas
			if c == '-' {
				continue
			}
			if enpasCnt == 0 {
				cb.enpas = 0
				enpasFile := (int8)(c - 'a')
				if enpasFile < 0 {
					enpasFile += 32
				}
				cb.enpas += enpasFile
				enpasCnt++
			} else if enpasCnt == 1 {
				enpasRank := (int8)(c - '1')
				enpasRank = enpasRank << 3
				cb.enpas += enpasRank
				enpasCnt++
			}
		default:
			continue
		}
	}

}

func (cb *ChessBoard) GenFEN() string {
	// this function will generate a FEN string
	var FEN strings.Builder
	FEN.Grow(256)

	empSquare := 0

	for i := 7; i >= 0; i-- {
		for j := 0; j < 8; j++ {
			if cb.board[(i<<3)+j] != nil {
				if empSquare != 0 {
					fmt.Fprintf(&FEN,"%d",empSquare)
				}
				fmt.Fprintf(&FEN,"%s",(string)(cb.board[(i<<3)+j].rep))
				empSquare = -1
			}
			empSquare++
		}
		if empSquare != 0 {
			fmt.Fprintf(&FEN,"%d",empSquare)
		}
		empSquare = 0
		if ( i > 0) {
			fmt.Fprintf(&FEN,"/")
		}
	}

	fmt.Fprintf(&FEN," ")

	if cb.nextMove == 0 {
		fmt.Fprintf(&FEN,"w ")
	} else {
		fmt.Fprintf(&FEN,"b ")
	}

	if !(cb.castle[0] || cb.castle[1] || cb.castle[2] || cb.castle[3]) {
		fmt.Fprintf(&FEN,"- ")
	} else {
		if cb.castle[0] {
			fmt.Fprintf(&FEN,"K")
		}
		if cb.castle[1] {
			fmt.Fprintf(&FEN,"Q")
		}
		if cb.castle[2] {
			fmt.Fprintf(&FEN,"k")
		}
		if cb.castle[3] {
			fmt.Fprintf(&FEN,"q")
		}
		fmt.Fprintf(&FEN," ")
	}

	if cb.enpas == -1 {
		fmt.Fprintf(&FEN,"-")
	} else {
		fmt.Fprintf(&FEN,"%v%d",(string)((cb.enpas&7)+'A'),(cb.enpas&56)>>3+1)
	}

	// fmt.Printf("%s\n",FEN.String())

	return FEN.String()
}

func (cb *ChessBoard) Perft(depth int) uint64 {
	// this function will perform perft function to go to
	// all moves in chess board

	var nodes uint64

	if depth == 0 {
		return 1
	}

	nodes = 0
	// fmt.Printf("nextMove before GenMoves(): %v\n",cb.nextMove)
	// cb.PrintBoard()
	cb.GenMoves()
	// cb.PrintMoves()
	// fmt.Printf("nextMove after GenMoves(): %v\n",cb.nextMove)
	// fmt.Printf("attack squares: %v\n",cb.attackSquares)

	cpyMoves := make([]Move, len(cb.moves))
	copy(cpyMoves, cb.moves)

	// fmt.Printf("depth(%v): %v\n",depth,cpyMoves)
	// cb.PrintBoard()
	// cb.PrintPieces()

	// QPerft
	if depth == 1 {
		return (uint64)(len(cb.moves))
	}

	for _, m := range cpyMoves {
		cb.makeMove(m)
		// fen := cb.GenFEN()
		// fmt.Printf("%s\n",fen)
		nodes += cb.Perft(depth - 1)
		cb.undoMove(m)
	}
	return nodes
}

func (cb *ChessBoard) calcMinDist() {
	// this function will calculate the minimum distance to each
	// side from every square on the board
	//
	// the minimum distances will be stored in a single array
	// with the offset as such:
	// 6 bits for square position - 3 bits for direction
	// square position: 0b000000 -> A1, 0b111111 -> H8
	// direction: counter-clockwise rotation
	// 000 -> right, 011 -> down, 100 -> up-right, 111 -> down-right

	var file int8
	var rank int8

	// loop through all possible min distances for each square
	for i := 0; i < 512; i += 8 {
		file = (int8)(i >> 3) & 7
		rank = (int8)(i >> 6) & 7

		// calculate distances for single directions
		// i.e. rank or file movement

		// right
		cb.minDist[i] = 7 - file
		// up
		cb.minDist[i+1] = 7 - rank
		// left 
		cb.minDist[i+2] = file
		// down
		cb.minDist[i+3] = rank

		// up-right
		if 7 - file > 7 - rank {
			cb.minDist[i+4] = 7 - rank
		} else {
			cb.minDist[i+4] = 7 - file
		}
		// up-left
		if 7 - rank > file {
			cb.minDist[i+5] = file
		} else {
			cb.minDist[i+5] = 7 - rank
		}
		// down-left 
		if file > rank {
			cb.minDist[i+6] = rank
		} else {
			cb.minDist[i+6] = file
		}
		// down-right
		if rank > 7 - file {
			cb.minDist[i+7] = 7 - file
		} else {
			cb.minDist[i+7] = rank
		}
	}
}

func initSide(color uint8) [16]Piece {
	var side [16]Piece

	name := [6]byte{'k', 'q', 'r', 'b', 'n', 'p'}
	if color == 0 {
		name = [6]byte{'K', 'Q', 'R', 'B', 'N', 'P'}
	}

	side[0] = Piece{alive: true, piece: 0, color: color, rep: name[0]}  // king
	side[1] = Piece{alive: true, piece: 5, color: color, rep: name[1]}  // queen
	side[2] = Piece{alive: true, piece: 4, color: color, rep: name[2]}  // rook
	side[3] = Piece{alive: true, piece: 4, color: color, rep: name[2]}  // rook
	side[4] = Piece{alive: true, piece: 3, color: color, rep: name[3]}  // bishop
	side[5] = Piece{alive: true, piece: 3, color: color, rep: name[3]}  // bishop
	side[6] = Piece{alive: true, piece: 2, color: color, rep: name[4]}  // knight
	side[7] = Piece{alive: true, piece: 2, color: color, rep: name[4]}  // knight
	side[8] = Piece{alive: true, piece: 1, color: color, rep: name[5]}  // pawn
	side[9] = Piece{alive: true, piece: 1, color: color, rep: name[5]}  // pawn
	side[10] = Piece{alive: true, piece: 1, color: color, rep: name[5]} // pawn
	side[11] = Piece{alive: true, piece: 1, color: color, rep: name[5]} // pawn
	side[12] = Piece{alive: true, piece: 1, color: color, rep: name[5]} // pawn
	side[13] = Piece{alive: true, piece: 1, color: color, rep: name[5]} // pawn
	side[14] = Piece{alive: true, piece: 1, color: color, rep: name[5]} // pawn
	side[15] = Piece{alive: true, piece: 1, color: color, rep: name[5]} // pawn

	return side
}

func (cb *ChessBoard) CurrTurn() uint8 {
	// this will return which side will make next move
	// return 0 for white and return 1 for black
	return cb.nextMove
}

func (cb *ChessBoard) GetSide(color uint8) *[16]Piece {
	// this function will return a pointer to the color pieces
	// of the board
	if color == 0 {
		return &cb.white
	} else {
		return &cb.black
	}
}

func (cb *ChessBoard) GetBoard() *[64]*Piece {
	// this function will return the board
	return &cb.board
}

func (cb *ChessBoard) GetMoves() *[]Move {
	// return moves
	return &cb.moves
}

func (cb *ChessBoard) GetPrevMoves() *[]Move {
	// return prevMoves
	return &cb.prevMoves
}

func (cb *ChessBoard) GetAttackSquares() *[]int8 {
	// return attackSquares
	return &cb.attackSquares
}

func (cb *ChessBoard) GetCastle() *[4]bool {
	// return castle
	return &cb.castle
}

func (cb *ChessBoard) GetEnpas() int8 {
	// return enpas
	return cb.enpas
}

func (cb *ChessBoard) GetNextMove() uint8 {
	// return nextMove
	return cb.nextMove
}

func (cb *ChessBoard) GetPinned() *[]*Piece {
	// return pinned
	return &cb.pinned
}