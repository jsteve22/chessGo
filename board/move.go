package board

type Move struct {
	start         int8
	end           int8
	pieceMoved    *Piece // pointer to piece that moved
	pieceCaptured *Piece // pointer if a piece was taken
	capPos        int8   // position where the piece was captured
	castle        uint8  // 1=K; 2=Q; 3=k; 4=q
	color         uint8  // color of move
	promotion     int8   // knight=2; bishop=3; rook=4; queen5
}

func (m *Move) GetStart() int8 {
	// return start
	return m.start
}

func (m *Move) GetEnd() int8 {
	// return end
	return m.end
}