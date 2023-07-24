package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jsteve22/chessGo/board"
)

func main() {
	// this is a comment


	// currentGame := board.LoadBoard("8/2p5/3p4/KP5r/1R3pPk/8/4P3/8 w KQkq - 23 13")
	currentGame := board.LoadBoard("rnbqkbnr/pp1ppppp/8/2p5/4P3/7P/PPPP1PP1/RNBQKBNR w KQkq c6 0 2")
	// currentGame := board.NewGame()

	board.PrintGame(currentGame)

	fmt.Printf("\n")
	board.GenerateMoves(currentGame)

	rand.Seed(time.Now().Unix())

	// generate a row of pawns for each side, this will be a test
	// to see how fast Perft can generate moves at different depths
 	// b.FENSet("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq -")
	// b.FENSet("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P1Q1/2N5/PPPBBPpP/R3K2R w KQkq -")

	// b.FENSet("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R4K1R b kq -")
	// b.FENSet("8/2p5/3p4/KP5r/1R3pPk/8/4P3/8 b - g3")
	// b.PrintBoard()

	// FENrec := b.GenFEN()

	// fmt.Printf("%s\n",FENrec)

	
	// b.FENSet("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQK2R")
	// b.FENSet("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
	// b.FENSet("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R")
	// b.FENSet("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

}