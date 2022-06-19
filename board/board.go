package board

import (
	"fmt"
)

type Piece struct {
	alive bool
	pos uint8
	piece int 	// 0: king; 1: pawn; 2: knight; 3: bishop; 4: rook; 5: queen 
	moves []uint8
	color uint8
	rep byte
}

type ChessBoard struct {
	board [64]*Piece
	white [16]Piece
	black [16]Piece

	nextmove uint
	turn uint
	revcnt uint
}

func Aboard() {
	fmt.Printf("board test")
}

func InitBoard() ChessBoard {
	var newCB ChessBoard
	
	newCB.white = initSide(0)
	newCB.black = initSide(1)	

	newCB.FENSet("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")

	newCB.nextmove = 0
	newCB.turn = 1
	newCB.revcnt = 0

	return newCB
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
			addP.pos = (uint8)(pos)
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