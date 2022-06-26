package board

func QueenMove(cb *ChessBoard, p Piece) {
	BishopMove(cb, p)
	RookMove(cb, p)
}

func QueenAttack(cb *ChessBoard, p Piece) {
	BishopAttack(cb, p)
	RookAttack(cb, p)
}