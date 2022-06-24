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

	/*
	n := 4
	x := b.Perft(n)

	fmt.Printf("Perft(%v): %v\n",n,x)
	return
	*/

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