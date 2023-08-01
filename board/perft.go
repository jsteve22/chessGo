package board

func Perft(game Game, depth uint64) uint64 {

	if depth == 0 {
		return 1
	}

	if depth == 1 {
		moves := GenerateMoves(game)
		// if len(moves) == 0 {
		// 	PrintBoard(game)
		// }
		return uint64(len(moves))
	}

	count := uint64(0)

	moves := GenerateMoves(game)
	// count += uint64(len(moves))
	for _, move := range moves {
		newGame := MakeMove(game, move)
		count += Perft(newGame, depth-1)
	}

	return count
}