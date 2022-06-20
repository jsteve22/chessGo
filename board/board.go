package board

import (
	"fmt"
)

type Move struct {
	start int8
	end int8
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

	nextmove int
	turn uint
	revcnt uint

	castle [4]bool	// KQkq
	enpas int8
}

func Aboard() {
	fmt.Printf("board test")
}

func InitBoard() ChessBoard {
	var newCB ChessBoard
	
	newCB.white = initSide(0)
	newCB.black = initSide(1)	

	newCB.FENSet("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")

	newCB.castle[0] = true
	newCB.castle[1] = true
	newCB.castle[2] = true
	newCB.castle[3] = true

	newCB.nextmove = 0
	newCB.turn = 1
	newCB.revcnt = 0

	return newCB
}

func (cb *ChessBoard) UpdateNextMove() {
	cb.nextmove = cb.nextmove ^ 1
}

func (cb *ChessBoard) MakeMove(move string) {
	// this function will make a move on the board 
	sf := (int8)(move[0] - 'A')
	sr := (int8)(move[1] - '1')

	ef := (int8)(move[2] - 'A')
	er := (int8)(move[3] - '1')

	st := sf + (sr << 3)
	en := ef + (er << 3)

	if cb.board[en] != nil {
		cb.board[en].alive = false
	}

	cb.board[en] = cb.board[st]
	cb.board[st] = nil
	cb.board[en].pos = en
}

func (cb *ChessBoard) GenMoves() {
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

	for _, c := range str {
		// if c == '/' 
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

func (b *ChessBoard) PrintBoard() {
	var p byte
	var rank int = 7
	var file int = 0
	var i int
	printHorizontalLine()
	for {
		if rank < 0 {
			break
		}
		i = file + ( rank << 3 )
		if ( b.board[i] == nil ) {
			p = ' '
		} else {
			p = b.board[i].rep
		}
		fmt.Printf("| %v ", (string)(p)) 
		file++
		if file == 8 {
			fmt.Printf("|\n")
			printHorizontalLine()
			file = 0
			rank--
		}
	}
}

func printHorizontalLine() {
	fmt.Printf("---------------------------------\n")
}

func (cb *ChessBoard) PrintMoves() {
	for _,m := range cb.moves {
		sFile := m.start & 7
		sRank := (m.start &  56) >> 3
		eFile := m.end & 7
		eRank := (m.end &  56) >> 3

		fmt.Printf("[%v] (%v,%v) -> (%v,%v)\n",(string)(cb.board[m.start].rep),sFile,sRank,eFile,eRank)
	}
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