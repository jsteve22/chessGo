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

	// generate a row of pawns for each side, this will be a test
	// to see how fast Perft can generate moves at different depths
 	// b.FENSet("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq -")

	// b.FENSet("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P1Q1/2N5/PPPBBPpP/R3K2R w KQkq -")

	/*
		var fenboard, fenside, fencas, fenenpas string
		fmt.Printf("Enter FEN string: \n")
		fmt.Scanf("%s %s %s %s",&fenboard,&fenside,&fencas,&fenenpas)
		var FEN strings.Builder
		FEN.Grow(256)
		fmt.Fprintf(&FEN,"%s %s %s %s",fenboard,fenside,fencas,fenenpas)

		// fmt.Printf("%s\n",FEN.String())
		b.FENSet( FEN.String() )
	*/

	// b.FENSet("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R4K1R b kq -")
	// b.FENSet("8/2p5/3p4/KP5r/1R3pPk/8/4P3/8 b - g3")
	// b.PrintBoard()

	// FENrec := b.GenFEN()

	// fmt.Printf("%s\n",FENrec)

	// b.PrintBoard()
	
	n := 5
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