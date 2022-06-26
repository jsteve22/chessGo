package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jsteve22/chessGo/board"
)

func main() {
	// this is a comment
	var b board.ChessBoard

	b.InitBoard()

	// b.FENSet("8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w KQkq -")
	// b.FENSet("r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq -")
	// b.FENSet("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/R3K2R")
	
	// b.FENSet("6rb/7P/8/8/8/8/k4q2/7K w - -")
	// b.FENSet("8/pp5P/kp6/1p6/1P6/KP5p/PP6/8 w - -")
	b.FENSet("7N/pp6/kp6/1p6/1P6/KP5p/PP6/8 w - -")

	n := 2
	x := b.Perft(n)

	fmt.Printf("Perft(%v): %v\n",n,x)
	return
	

	// b.FENSet("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQK2R")
	// b.FENSet("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
	b.FENSet("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R")

	var nextMove string
	badMove := false

	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		b.PrintBoard()
		b.PrintCastlePriv()

		if badMove {
			fmt.Printf("Bad move\n")
		}

		b.GenMoves()

		// check for checkmate
		if b.CheckMate() {
			fmt.Printf("GAME OVER\n")
			return
		}

		b.PrintPrevMoves()
		b.PrintMoves()
		// fmt.Printf("\n")
		// b.PrintPieces()

		fmt.Scan(&nextMove)

		if b.UserMove(nextMove) {
			b.UpdateNextMove()
			badMove = false
		} else {
			badMove = true
		}
	}
}