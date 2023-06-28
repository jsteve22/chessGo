package board

type Piece struct {
	pos   uint8
	piece uint8 // pawn=1; knight=2; bishop=3; rook=4; queen=5; king=6
	color uint8
	rep   string
}

func (p *Piece) GetRep() string {
	// return rep attribute
	return p.rep
}

func (p *Piece) GetPos() uint8 {
	// return pos attribute
	return p.pos
}