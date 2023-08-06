package board

import (
	"math/rand"
	"time"
)

func GenerateMoves(game Game) []Move {
	// var moves []Move
	// var pseudoMoves []Move
	moves := make([]Move, 0, 32)
	pseudoMoves := make([]Move, 0, 32)

	BLACK := uint8(8)
	COLOR_BLACK := uint8(1)
	PAWN := uint8(1)
	KNIGHT := uint8(2)
	BISHOP := uint8(3)
	ROOK := uint8(4)
	QUEEN := uint8(5)
	KING := uint8(6)

	if game.nextToPlay == COLOR_BLACK {
		PAWN += BLACK
		KNIGHT += BLACK
		BISHOP += BLACK
		ROOK += BLACK
		QUEEN += BLACK
		KING += BLACK
	}

	for index, piece := range game.board {
		switch piece {
		case PAWN:
			pseudoMoves = append(pseudoMoves, PawnGeneratePseudoLegalMoves(Piece{pos: uint8(index), color: game.nextToPlay}, game)...)
		case KNIGHT:
			pseudoMoves = append(pseudoMoves, KnightGeneratePseudoLegalMoves(Piece{pos: uint8(index), color: game.nextToPlay}, game)...)
		case BISHOP:
			pseudoMoves = append(pseudoMoves, BishopGeneratePseudoLegalMoves(Piece{pos: uint8(index), color: game.nextToPlay}, game)...)
		case ROOK:
			pseudoMoves = append(pseudoMoves, RookGeneratePseudoLegalMoves(Piece{pos: uint8(index), color: game.nextToPlay}, game)...)
		case QUEEN:
			pseudoMoves = append(pseudoMoves, QueenGeneratePseudoLegalMoves(Piece{pos: uint8(index), color: game.nextToPlay}, game)...)
		case KING:
			pseudoMoves = append(pseudoMoves, KingGeneratePseudoLegalMoves(Piece{pos: uint8(index), color: game.nextToPlay}, game)...)
		}
	}

	for _, pseudoMove := range pseudoMoves {
		nextGame := MakeMove(game, pseudoMove)
		nextGame.nextToPlay ^= 1
		if !IsKingInCheck(nextGame) {
			moves = append(moves, pseudoMove)
		}
	}

	// PrintMoves(game, moves)

	return moves
}

func GenerateAttacks(game Game, color uint8) uint64 {
	bitboard := uint64(0)

	COLOR_BLACK := uint8(1)
	BLACK := uint8(8)
	PAWN := uint8(1)
	KNIGHT := uint8(2)
	BISHOP := uint8(3)
	ROOK := uint8(4)
	QUEEN := uint8(5)
	KING := uint8(6)

	if color == COLOR_BLACK {
		PAWN += BLACK
		KNIGHT += BLACK
		BISHOP += BLACK
		ROOK += BLACK
		QUEEN += BLACK
		KING += BLACK
	}

	for index, square := range game.board {
		switch square {
		case PAWN:
			bitboard |= PawnGenerateAttackSquaresBitboard(Piece{pos: uint8(index), color: game.nextToPlay ^ 1})
		case KNIGHT:
			bitboard |= KnightGenerateAttackSquaresBitboard(uint8(index))
		case BISHOP:
			bitboard |= BishopGenerateAttackSquaresBitboard(uint8(index), game)
		case ROOK:
			bitboard |= RookGenerateAttackSquaresBitboard(uint8(index), game)
		case QUEEN:
			bitboard |= QueenGenerateAttackSquaresBitboard(uint8(index), game)
		case KING:
			bitboard |= KingGenerateAttackSquaresBitboard(uint8(index))
		}
	}

	// fmt.Printf("bitboard: %d\n", bitboard)
	return bitboard
}

func MakeMove(game Game, move Move) Game {
	// this function will make a move without question
	// even if the move is illegal... might change later

	// 1. update pieces for white/black
	// 2. check for enpassant
	// 3. make move on board
	// 4. update enpassant
	// 5. move rook if castling and update rook in pieces

	XOR := uint8(1)
	MASK := uint8(8 - 1)
	endRank := move.end >> 3
	endFile := move.end & MASK
	startRank := move.start >> 3
	startFile := move.start & MASK
	COLOR := game.board[move.start] >> 3
	PIECE := game.board[move.start] & MASK
	PAWN := uint8(1)
	KING := uint8(6)
	ROOK := uint8(4)
	emptySquare := uint8(0)

	// update clocks
	game.fullMoveClock += 1
	game.halfMoveClock += 1

	if PIECE == PAWN && move.end == uint8(game.enPassant) {
		deleteRank := endRank - 1 + ((COLOR ^ XOR) * 2)
		game.board[(deleteRank<<3)+endFile] = 0
	}

	// king side castle
	if PIECE == KING && startFile == 4 && endFile == 6 {
		game.castlingRights[COLOR*2+0] = false
		game.castlingRights[COLOR*2+1] = false
		game.board[(startRank<<3)+5] = game.board[(startRank<<3)+7]
		game.board[(startRank<<3)+7] = 0
	}

	// queen side castle
	if PIECE == KING && startFile == 4 && endFile == 2 {
		game.castlingRights[COLOR*2+0] = false
		game.castlingRights[COLOR*2+1] = false
		game.board[(startRank<<3)+3] = game.board[(startRank<<3)+0]
		game.board[(startRank<<3)+0] = 0
	}

	if PIECE == ROOK && startFile == 0 {
		game.castlingRights[COLOR*2+1] = false
	}

	if PIECE == ROOK && startFile == 7 {
		game.castlingRights[COLOR*2+0] = false
	}

	if PIECE == KING && startFile == 4 {
		game.castlingRights[COLOR*2+0] = false
		game.castlingRights[COLOR*2+1] = false
	}

	// update half move clock if piece is taken or a pawn is moved
	if PIECE == PAWN || game.board[move.end] != emptySquare {
		game.halfMoveClock = 0
	}

	// update board
	// pieceTaken := game.board[move.end] != 0
	game.board[move.end] = game.board[move.start]
	game.board[move.start] = 0

	// if pieceTaken {
	// 	PrintBoard(game)
	// }

	// update color to move
	game.nextToPlay ^= XOR

	game.enPassant = -1
	if PIECE == PAWN && (endRank-startRank == 2 || startRank-endRank == 2) {
		enPasRank := endRank - 1 + ((COLOR ^ XOR) * 2)
		game.enPassant = int8((enPasRank << 3) + endFile)
	}

	if PIECE == PAWN && move.promotion != 0 {
		game.board[move.end] = move.promotion
	}
	// if PIECE == 1 && endRank == 4 && startRank == 6 {
	// 	game.enPassant = int8((5 << 3) + endFile)
	// }
	// if PIECE == 1 && endRank == 3 && startRank == 1 {
	// 	game.enPassant = int8((2 << 3) + endFile)
	// }

	return game
}

func IsKingInCheck(game Game) bool {
	COLOR := game.nextToPlay
	XOR := uint8(1)
	// WHITE := uint8(0)
	KING := uint8(6 + (8 * COLOR))

	bitboardAttacks := GenerateAttacks(game, COLOR^XOR)

	// pieces := game.whitePieces
	// if COLOR == BLACK {
	// 	pieces = game.blackPieces
	// 	KING += 8
	// }

	kingBitBoard := uint64(0)

	for index, square := range game.board {
		if square == KING {
			kingBitBoard = uint64(1 << index)
			break
		}
	}

	return (bitboardAttacks & (kingBitBoard)) != 0
}

func GetWinner(game Game) int8 {
	// only call this function if there are no moves to make
	if !IsKingInCheck(game) {
		return -1
	}

	XOR := uint8(1)

	return int8(game.nextToPlay ^ XOR)
}

func ComputerMakeMove(game Game) Game {
	moves := GenerateMoves(game)

	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)

	index := generator.Intn(len(moves))

	game = MakeMove(game, moves[index])

	return game
}

/*
func (cb *ChessBoard) makeMove(move Move) {
	// this function will make a move across the board and update
	// the move counter. Each move made should be reversible in
	// undoMove, even if it's a capture, a castle, or a promotion.

	// initialize the start and end of the given move
	st := move.start
	en := move.end

	// create appendMove which will be added to end of prevMove slice
	// use the appendMove to make the move across the board
	var appendMove Move
	appendMove.start = st
	appendMove.end = en
	appendMove.pieceMoved = cb.board[st]
	appendMove.color = cb.nextMove
	appendMove.prevEnpas = cb.enpas
	appendMove.prevCastle[0] = cb.castle[0]
	appendMove.prevCastle[1] = cb.castle[1]
	appendMove.prevCastle[2] = cb.castle[2]
	appendMove.prevCastle[3] = cb.castle[3]
	appendMove.fullmove = cb.fullmove
	appendMove.halfmove = cb.halfmove

	// increment clock
	cb.fullmove = cb.fullmove + 1
	cb.halfmove = cb.halfmove + 1

	// if the end position is not nil, a piece is captured
	if cb.board[en] != nil {
		// set piece to dead
		cb.board[en].alive = false
		// add captured piece to move
		appendMove.pieceCaptured = cb.board[en]
		appendMove.capPos = en
		// reset halfmove counter
		cb.halfmove = 0
	}

	// this is the naive movement, for simple movement
	cb.board[en] = appendMove.pieceMoved // cb.board[st]
	appendMove.pieceMoved.pos = en
	cb.board[st] = nil

	// check if move was enpassant take and reset enpas
	enpasMove := (appendMove.pieceMoved.piece == 1 && en == cb.enpas)
	cb.enpas = -1

	// check if move was a castle
	// check if piece moved was a king
	if cb.board[en].piece == 0 {
		stRank := st & 56
		stFile := st & 7
		enRank := en & 56
		enFile := en & 7

		// changed castling permissions
		cb.castle[2*cb.board[en].color] = false
		cb.castle[(2*cb.board[en].color)+1] = false

		// check if king is castling and to which side
		if stRank == enRank && enFile-stFile == 2 {
			// king is castling king side
			// move rook to F1/F8
			appendMove.castle = (2 * cb.board[en].color) + 1
			// appendMove.pieceCaptured = cb.board[en+1]
			// appendMove.capPos = (stRank << 3) + 7
			cb.board[st+1] = cb.board[en+1]
			cb.board[st+1].pos = st + 1
			cb.board[en+1] = nil
		} else if stRank == enRank && stFile-enFile == 2 {
			// king is castling queen side
			// move rook to D1/D8
			appendMove.castle = (2 * cb.board[en].color) + 2
			// appendMove.pieceCaptured = cb.board[stRank << 3]
			// appendMove.capPos = stRank << 3
			cb.board[st-1] = cb.board[stRank]
			cb.board[st-1].pos = st - 1
			cb.board[stRank] = nil
		}
	}

	// check if piece moved was a rook
	if cb.board[en].piece == 4 {
		// check which side the rook was on before updating castling
		// check if on queen side
		if appendMove.start == (int8)(56*appendMove.color) {
			cb.castle[(2*appendMove.color)+1] = false
		}
		// check if on king side
		if appendMove.start == (int8)(56*appendMove.color+7) {
			cb.castle[2*appendMove.color] = false
		}
	}

	// if start or end of a move is on A1,A8,H1, or H8 get rid of castling
	if st == 0 || en == 0 {
		cb.castle[1] = false
	}

	if st == 7 || en == 7 {
		cb.castle[0] = false
	}

	if st == 56 || en == 56 {
		cb.castle[3] = false
	}

	if st == 63 || en == 63 {
		cb.castle[2] = false
	}

	// check if piece moved was a pawn
	if cb.board[en].piece == 1 {
		// reset halfmove counter
		cb.halfmove = 0

		stRank := (st & 56) >> 3
		// stFile := st & 7
		enRank := (en & 56) >> 3
		enFile := en & 7

		// check if the move was a double push
		if enRank-stRank == 2 || stRank-enRank == 2 {
			cb.enpas = st + (4 * (enRank - stRank))
			// cb.enpas = -1
		}

		// check if move was an enpas take
		if enpasMove {
			enpasPawnPos := (stRank << 3) + enFile
			// fmt.Printf("%v", enpasPawnPos )
			appendMove.pieceCaptured = cb.board[enpasPawnPos]
			appendMove.capPos = enpasPawnPos
			cb.board[enpasPawnPos].alive = false
			cb.board[enpasPawnPos] = nil
			// cb.enpas = -1
		}

		// promote pawn to new piece
		if move.promotion != 0 {
			appendMove.pieceMoved.piece = move.promotion
			appendMove.promotion = move.promotion
			switch move.promotion {
			case 2:
				appendMove.pieceMoved.rep = (rune)('N' + (appendMove.pieceMoved.color * 32))
			case 3:
				appendMove.pieceMoved.rep = (rune)('B' + (appendMove.pieceMoved.color * 32))
			case 4:
				appendMove.pieceMoved.rep = (rune)('R' + (appendMove.pieceMoved.color * 32))
			case 5:
				appendMove.pieceMoved.rep = (rune)('Q' + (appendMove.pieceMoved.color * 32))
			}
		}
	}

	// change turns on the board
	cb.nextMove = 1 ^ cb.nextMove

	// add new move to slice
	cb.prevMoves = append(cb.prevMoves, appendMove)
}

func (cb *ChessBoard) undoMove(move Move) {
	// undo a given move on the board. This will
	// only undo the previous move that was on the
	// board. If provided a different move than the
	// immediate previous one, then could cause issues

	lastMove := cb.prevMoves[len(cb.prevMoves)-1]

	// reset board positions with the move
	cb.nextMove = lastMove.color
	cb.enpas = lastMove.prevEnpas
	cb.fullmove = lastMove.fullmove
	cb.halfmove = lastMove.halfmove
	cb.castle[0] = lastMove.prevCastle[0]
	cb.castle[1] = lastMove.prevCastle[1]
	cb.castle[2] = lastMove.prevCastle[2]
	cb.castle[3] = lastMove.prevCastle[3]

	if lastMove.promotion != 0 {
		// fmt.Printf("1 lastMove: %v\n", *lastMove.pieceMoved)
		lastMove.pieceMoved.piece = 1
		lastMove.pieceMoved.rep = (rune)('P' + (32 * lastMove.pieceMoved.color))
		// fmt.Printf("2 lastMove: %v\n", *lastMove.pieceMoved)
		// fmt.Printf("\n")
	}

	st := move.start
	en := move.end

	// reset piece on board to previous position
	cb.board[move.start] = lastMove.pieceMoved
	lastMove.pieceMoved.pos = st
	cb.board[move.end] = nil

	// bring pack a piece if captured
	if lastMove.pieceCaptured != nil {
		lastMove.pieceCaptured.alive = true
		cb.board[lastMove.capPos] = lastMove.pieceCaptured
	}

	// check if move was a castle
	// check if piece moved was a king
	if lastMove.pieceMoved.piece == 0 {
		stRank := (st & 56)
		stFile := st & 7
		enRank := (en & 56)
		enFile := en & 7

		// check if king is castling and to which side
		// if stRank == enRank && enFile - stFile == 2 {
		if lastMove.castle == 1 || lastMove.castle == 3 {
			// king is castling king side
			// move rook to F1/F8
			cb.board[en+1] = cb.board[st+1]
			cb.board[en+1].pos = en + 1
			cb.board[st+1] = nil
		} else if stRank == enRank && stFile-enFile == 2 {
			// king is castling queen side
			// move rook to D1/D8
			cb.board[stRank] = cb.board[st-1]
			cb.board[stRank].pos = stRank
			cb.board[st-1] = nil
		}
	}

	// delete last element of prevMoves
	cb.prevMoves[len(cb.prevMoves)-1] = Move{}
	cb.prevMoves = cb.prevMoves[:len(cb.prevMoves)-1]
}

func (cb *ChessBoard) GenMoves() {

	cb.inCheck(cb.nextMove)
	cb.PinPieces(cb.nextMove)

	cb.moves = make([]Move, 0, 256)

	var pieces *[16]Piece
	//var forward int8

	if cb.nextMove == 0 {
		pieces = &cb.white
		//forward = 8
	} else {
		pieces = &cb.black
		//forward = -8
	}

	for _, p := range pieces {
		if !p.alive {
			continue
		}
		switch p.piece {
		case 0:
			KingMove(cb, p)
		case 1:
			PawnMove(cb, p)
		case 2:
			KnightMove(cb, p)
		case 3:
			BishopMove(cb, p)
		case 4:
			RookMove(cb, p)
		case 5:
			QueenMove(cb, p)
		}
	}
}

*/

// save for later I guess?
// get pieces using piece lists instead of board
// BLACK := uint8(8)
// COLOR_BLACK := uint8(1)
// PAWN := uint8(1)
// KNIGHT := uint8(2)
// BISHOP := uint8(3)
// ROOK := uint8(4)
// QUEEN := uint8(5)
// KING := uint8(6)
//
// pieces := game.whitePieces
// if game.nextToPlay == COLOR_BLACK {
// pieces = game.blackPieces
// PAWN += BLACK
// KNIGHT += BLACK
// BISHOP += BLACK
// ROOK += BLACK
// QUEEN += BLACK
// KING += BLACK
// }
//
// for _, piece := range pieces {
// switch piece.piece {
// case PAWN:
// pseudoMoves = append(pseudoMoves, PawnGeneratePseudoLegalMoves(piece, game)...)
// case KNIGHT:
// pseudoMoves = append(pseudoMoves, KnightGeneratePseudoLegalMoves(piece, game)...)
// case BISHOP:
// pseudoMoves = append(pseudoMoves, BishopGeneratePseudoLegalMoves(piece, game)...)
// case ROOK:
// pseudoMoves = append(pseudoMoves, RookGeneratePseudoLegalMoves(piece, game)...)
// case QUEEN:
// pseudoMoves = append(pseudoMoves, QueenGeneratePseudoLegalMoves(piece, game)...)
// case KING:
// pseudoMoves = append(pseudoMoves, KingGeneratePseudoLegalMoves(piece, game)...)
// }
// }