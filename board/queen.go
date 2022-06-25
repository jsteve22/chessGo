package board

func queenMove(cb *ChessBoard, p Piece) {
	bishopMove(cb, p)
	rookMove(cb, p)
}

func queenAttack(cb *ChessBoard, p Piece) {
	bishopAttack(cb, p)
	rookAttack(cb, p)
}