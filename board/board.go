package board

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	board [64]uint8
	halfMoveClock uint
	fullMoveClock uint
	nextToPlay uint8
	castlingRights [4]bool // [ white king side, white queen side, black king side, black queen side ]
	enPassant int8
}

/*

index 0  => A8
index 7  => H8
...
index 56 => A1
index 63 => H1

| ------------------------------------- |
| 00 | 01 | 02 | 03 | 04 | 05 | 06 | 07 |
| ------------------------------------- |
| 08 | 09 | 10 | 11 | 12 | 13 | 14 | 15 |
| ------------------------------------- |
| 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 |
| ------------------------------------- |
| 24 | 25 | 26 | 27 | 28 | 29 | 30 | 31 |
| ------------------------------------- |
| 32 | 33 | 34 | 35 | 36 | 37 | 38 | 39 |
| ------------------------------------- |
| 40 | 41 | 42 | 43 | 44 | 45 | 46 | 47 |
| ------------------------------------- |
| 48 | 49 | 50 | 51 | 52 | 53 | 54 | 55 |
| ------------------------------------- |
| 56 | 57 | 58 | 59 | 60 | 61 | 62 | 63 |
| ------------------------------------- | 
*/

func NewGame() Game {
	return LoadBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
}

func LoadBoard(FEN string) Game {
	var game Game
	var token byte
	var err error

	reader := strings.NewReader(FEN)

	FENBoard(reader, &(game.board))

	// Load pieces into the board
	// Load which side has the chance to play
	token, err = reader.ReadByte()
	if (err != nil) {
		return game
	}
	game.nextToPlay = uint8(0)
	if (token == byte('b')) {
		game.nextToPlay = uint8(1);
	}
	token, err = reader.ReadByte()
	if (err != nil) {
		return game
	}

	if (token != ' ') {
		return game
	}

	// castling rights
	FENCastlingRights(reader, &(game.castlingRights))

	// En passant target square
	game.enPassant, err = FENEnPassant(reader)
	if (err != nil) {
		return game
	}

	token, err = reader.ReadByte()
	if (err != nil) {
		return game
	}
	if (token != ' ') {
		return game
	}

	// Halfmove Clock
	game.halfMoveClock, err = FENClock(reader)
	if (err != nil) {
		return game
	}

	// Fullmove Clock
	game.fullMoveClock, err = FENClock(reader)
	if (err != nil) {
		return game
	}

	return game
}

func FENBoard(reader *strings.Reader, board *[64]uint8) error {
	board_index := 0

	for {
		token, err := reader.ReadByte()

		if (err != nil)  {
			return err
		}

		if (token == ' ') {
			break
		}

		if (token == '/') {
			continue
		}

		squares_to_skip, err := strconv.Atoi(string(token))

		if err != nil {
			board[board_index] = FENByteToInt(token) 
			board_index++
			continue
		}
		board_index += squares_to_skip
	}
	return nil
}

func FENCastlingRights(reader *strings.Reader, castlingRights *[4]bool) error {
	// castling rights
	for {
		token, err := reader.ReadByte()

		if (err != nil) {
			return err
		}

		if (token == ' ') {
			return nil
		}

		if (token == '-') {
			reader.ReadByte()
			return nil
		}

		if (token == 'K') {
			castlingRights[0] = true
		}
		if (token == 'Q') {
			castlingRights[1] = true
		}
		if (token == 'k') {
			castlingRights[2] = true
		}
		if (token == 'q') {
			castlingRights[3] = true
		}
	}
}

func FENEnPassant(reader *strings.Reader) (int8, error) {

	token, err := reader.ReadByte()
	
	if (err != nil) {
		return -1, err
	}
	
	if (token == '-') {
		return -1, err
	}

	column := int8(token) - int8('a')

	token, err = reader.ReadByte()
	if (err != nil) {
		return -1, err
	}

	row, err := strconv.Atoi(string(token))
	if (err != nil) {
		return -1, err
	}
	row = 8 - row

	return (column + int8(row*8)), nil
}

func FENByteToInt(token byte) uint8 {
	switch (token) {
		case 'P':
			return 1
		case 'N':
			return 2
		case 'B':
			return 3
		case 'R':
			return 4
		case 'Q':
			return 5
		case 'K':
			return 6
		case 'p':
			return 9
		case 'n':
			return 10
		case 'b':
			return 11
		case 'r':
			return 12
		case 'q':
			return 13
		case 'k':
			return 14
	}
	return 0
}

func FENClock(reader *strings.Reader) (uint, error)  {
	clockStr := ""

	for {
		token, err := reader.ReadByte()
		if (err != nil || token == ' ') {
			break
		}
		clockStr = clockStr + string(token)
	}

	var clock int
	fmt.Sscanf(clockStr, "%d", &clock)

	if (clock < 0) {
		return 0, errors.New("negative clock value")
	}

	return uint(clock), nil
}

// func LoadPiecesFromBoard(game *Game) error {
// 	game.whitePieceCount = 0
// 	game.blackPieceCount = 0

// 	// go through the board and load in the pieces
// 	// into the black and white pieces array
// 	for index := 0; index < 64; index++ {
// 		pieceVal := game.board[index]

// 		if (pieceVal > 0 && (pieceVal&8) == 0) {
// 			whitePiecesIndex := game.whitePieceCount
// 			game.whitePieces[ whitePiecesIndex ], _ = GeneratePiece( pieceVal, uint8(index) )
// 			game.whitePieceCount++
// 		} else if (pieceVal > 0) {
// 			blackPiecesIndex := game.blackPieceCount
// 			game.blackPieces[ blackPiecesIndex ], _ = GeneratePiece( pieceVal, uint8(index) )
// 			game.blackPieceCount++
// 		}
// 	}

// 	return nil
// }

func GeneratePiece( pieceVal uint8, boardIndex uint8 ) (Piece, error) { 
	var piece Piece

	// chessNotation, _ := BoardIndexToChessNotation(boardIndex)
	// fmt.Printf("%s: %d\n", chessNotation, pieceVal)

	piece.piece = pieceVal
	piece.color = (pieceVal & 8) >> 3
	piece.rep   = PieceRepresentation(pieceVal)
	piece.pos   = boardIndex

	return piece, nil
}

func BoardIndexToChessNotation(index uint8) (string, error) {

	if (index >= 64) {
		return "", errors.New("index error: board index >= 64 for converting chess notation")
	}

	file := 8 - (index / 8)
	rank := byte(uint8('a') + (index % 8))

	notation := fmt.Sprintf("%c%d", rank, file)

	return notation, nil
}

func ChessNotationToBoardIndex(notation string) (uint8, error) {

	if (len(notation) > 2) {
		return 0, errors.New("value error: chess notation must be 2 characters long")
	}

	fileChar := notation[0]
	file := uint8(fileChar) - uint8('A')
	if file > 8 {
		file = file + uint8('A') - uint8('a')
	}

	if file >= 8 {
		errorString := fmt.Sprintf("value error: file value %d does not fit in chess board", file)
		return 0, errors.New(errorString)
	}

	var rank uint8
	fmt.Sscanf(string(notation[1]), "%d", &rank)
	rank = 8 - rank
	// fmt.Printf("str: %s\n", string(notation[1]))
	// fmt.Printf("rank: %d\n", rank)

	index := uint8((rank<<3) + file)

	return index, nil
}