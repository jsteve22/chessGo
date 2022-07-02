package board

type Piece struct {
	alive  bool
	pos    int8
	piece  int8 // king=0; pawn=1; knight=2; bishop=3; rook=4; queen=5
	color  uint8
	rep    byte
	pinned bool
}

func (p *Piece) GetAlive() bool {
	// return the alive attribute
	return p.alive
}

func (p *Piece) GetRep() byte {
	// return rep attribute
	return p.rep
}

func (p *Piece) GetPos() int8 {
	// return pos attribute
	return p.pos
}