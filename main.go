package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jsteve22/chessGo/board"
)

func main() {
	// this is a comment
	var b = board.InitBoard()

	b.FENSet("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQK2R")


	var nextMove string

	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		b.PrintBoard()

		b.GenMoves()

		fmt.Scan(&nextMove)

		b.MakeMove(nextMove)
	}
}