package main

import "github.com/jsteve22/chessGo/board"

func main() {
	// this is a comment
	var b = board.InitBoard()

	b.FENSet("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R")

	b.PrintBoard()
}