<template>
  <div>
    <!-- <div>
      drag started in {{ startDrag }}
      drag ended in {{ endDrag }}
      <div v-if="inPromotion">
        <button type="button" @click="board[endDrag].piece = 2" class="w-16 h-4 bg-emerald-600 hover:border-2 hover:border-orange-500">knight</button>
        <button type="button" @click="board[endDrag].piece = 3" class="w-16 h-4 bg-emerald-600 hover:border-2 hover:border-orange-500">bishop</button>
        <button type="button" @click="board[endDrag].piece = 4" class="w-16 h-4 bg-emerald-600 hover:border-2 hover:border-orange-500">rook</button>
        <button type="button" @click="board[endDrag].piece = 5" class="w-16 h-4 bg-emerald-600 hover:border-2 hover:border-orange-500">queen</button>
      </div>
      <button @click="fetchMoves" class="w-auto h-auto bg-emerald-200 rounded-lg p-1">
        fetch Moves
      </button>
      <button type="button" @click="() => {const p:number[] = []; board.forEach((square) => p.push(square.piece)); console.log(p);}" class="w-auto h-auto bg-blue-50 rounded-lg p-1">
        print board
      </button>
      <button type="button" @click="() => clearBoard()" class="w-auto h-auto bg-blue-50 rounded-lg p-1">
        clear board
      </button>
      <span class="text-xs">
        FEN = {{ currentFEN }}
      </span>
      <form>
      <input v-model="FENinput" class="text-xs w-64">
      <button @click="(event) => {event.preventDefault(); loadFEN(FENinput); generateFEN(); fetchMoves();}" class="w-auto h-auto bg-emerald-200 rounded-lg p-1">
        load FEN
      </button>
      </form>
    </div> -->
    <button @click="(event) => {event.preventDefault(); loadFEN(FENinput); generateFEN(); fetchMoves();}" class="w-auto h-auto bg-emerald-200 rounded-lg p-1">
      start new board
    </button>
    <div>
      half move = {{ halfClock }} | full move = {{ fullClock }} | game state = {{ (playing) ? "active" : "not active" }} 
      <span v-if="playing===false">| result = {{ (winner === -1) ? "DRAW" : (winner === 0) ? "WHITE WINS" : "BLACK WINS" }}</span>
    </div>
    <div class="grid grid-cols-8" id="ChessBoardGrid">
      <div v-for="square in board" v-bind:key="square.id">
        <chessboardSquare :color=square.color :id="(startDrag == -1 ) ? square.id : startDrag" :piece="board[square.id].piece" 
        :placable="board[square.id].placable" @dragStart="(startSquare:number) => startDrag = startSquare" 
        @pieceUpdate="(piece:number) => board[square.id].piece = piece" @dragEnd="(endSquare:number) => endDrag = endSquare"/>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import chessboardSquare from './ChessBoardSquare.vue';
import { FENByteToNumber, NumberToFENByte, IndexToChessNotation, ChessNotationToIndex, LoadCastlingFEN, GenerateCastlingFEN } from './ChessFunctions';

export default defineComponent({
  name: 'ChessBoard',
  components: {
    chessboardSquare
  },
  mounted() {
    interface Squares {
      color: number,
      id: number,
      piece: number,
      placable: boolean
    }
    const size = 8*8;
    const elements:(Squares)[] = [];
    const light = 0;
    const dark = 1;
    for (let i = 0; i < size; i++) {
      const rank = Math.floor(i/8);
      const file = i % 8;
      const color = ((rank%2 ^ file%2) === 0) ? light : dark; 
      elements.push( {color: color, id: i, piece: 0, placable: false} );
    }
    this.board = elements

    const chessboardElement = document.getElementById('ChessBoardGrid');
    chessboardElement?.addEventListener('drop', this.dropDraggable);
  },
  data() {
    interface Squares {
      color: number,
      id: number,
      piece: number,
      placable: boolean
    }
    const board:(Squares)[] = [];
    return {
      board: board,
      startDrag: -1,
      endDrag: -1,
      FENinput: 'rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1',
      currentFEN: '',
      moves: [],
      nextToPlay: 0,
      castling: [false, false, false, false],
      enPassant: -1,
      halfClock: 0,
      fullClock: 0,
      inPromotion: true,
      playing: false,
      winner: -1,
    }
  }, 
  updated() {
    // console.log(this.pieces);

    interface JsonMove {
      Start: number,
      End: number,
      Promotion: number,
      Castle: boolean,
    }

    let moves:Array<JsonMove> = this.moves;
    if (moves == null) {
      moves = [];
    }

    const pieceColor = (this.startDrag !== -1) ? this.board[this.startDrag].piece >> 3 : 0;
    for (let i = 0; i < 64; i++) {
      const differentColor = (this.board[i].piece >> 3) !== pieceColor;

      // go through the moves and see which moves are available
      let inMoves = false;
      for (let j = 0; j < moves.length; j++) {
        if (moves[j].Start === this.startDrag && moves[j].End === i) {
          inMoves = true;
        }
      }
      // const inMoves = true;

      const isPlacable = (i !== this.startDrag) && ((this.board[i].piece === 0) || differentColor) && (this.startDrag !== -1) && (inMoves);
      this.board[i].placable = (isPlacable) ? true : false;
    }
  },
  methods: {
    dropDraggable() { // this will check if a move is made when the player drops the piece
      console.log(`startDrag = ${this.startDrag}`);
      console.log(`endDrag = ${this.endDrag}`);
      if (this.endDrag === this.startDrag || this.board[this.endDrag].placable === false) {
        this.startDrag = -1;
        return;
      }

      if (this.startDrag !== -1) {  // a move is made
        const PIECE_COLOR = this.board[this.startDrag].piece >> 3;
        const PIECE = this.board[this.startDrag].piece & 7;
        const ROOK = 4;
        const KING = 6;
        const PAWN = 1;
        
        // update castling if the rooks move
        if ((PIECE === ROOK) && ((this.startDrag & 7) === 7)) {
          this.castling[2*PIECE_COLOR + 0] = false;
        }
        if ((PIECE === ROOK) && ((this.startDrag & 7) === 0)) {
          this.castling[2*PIECE_COLOR + 1] = false;
        }

        // update castling if king moves
        if (PIECE === KING) {
          this.castling[2*PIECE_COLOR + 0] = false;
          this.castling[2*PIECE_COLOR + 1] = false;
        }

        // check if a piece is captured
        if (this.board[this.endDrag].piece !== 0) {
          this.halfClock = -1;
        }

        // check a pawn is moved
        if (PIECE === PAWN) {
          this.halfClock = -1;
        }

        // a move is made
        this.board[this.endDrag].piece = this.board[this.startDrag].piece;
        this.board[this.startDrag].piece = 0;
        // console.log(this.board);

        this.nextToPlay = this.nextToPlay ^ 1;

        // delete piece of pawn takes en passant
        if ((PIECE === PAWN) && (this.endDrag === this.enPassant)) {
          const deletePawn = this.endDrag + 8 + (PIECE_COLOR*-16);
          this.board[deletePawn].piece = 0;
        }

        if ((PIECE === KING) && ((this.startDrag & 7) === 4) && ((this.endDrag & 7) === 6)) {
          this.board[this.startDrag+1].piece = this.board[this.startDrag+3].piece;
          this.board[this.startDrag+3].piece = 0;
        }

        if ((PIECE === KING) && ((this.startDrag & 7) === 4) && ((this.endDrag & 7) === 2)) {
          this.board[this.startDrag-1].piece = this.board[this.startDrag-4].piece;
          this.board[this.startDrag-4].piece = 0;
        }

        this.enPassant = -1;
        // update en passant if a pawn is moved twice
        if ((PIECE === PAWN) && (Math.abs((this.startDrag>>3)-(this.endDrag>>3)) === 2)) {
          this.enPassant = this.endDrag + 8 + (PIECE_COLOR*-16);
        }

        // if it is a PAWN promotion, then wait for the promotion to be made
        if ((PIECE === PAWN) && ((this.endDrag>>3)===0 || (this.endDrag>>3)===7)){
          this.inPromotion = true;
          this.board[this.endDrag].piece = 5 + (8*PIECE_COLOR);
        }

        // increment full move clock
        this.fullClock += 1;

        // increment the half move clock
        this.halfClock += 1;

        this.generateFEN();
        this.fetchMoves();
      }
      this.startDrag = -1;
    },

    clearBoard() {
      for (let i = 0; i < 64; i ++) {
        this.board[i].piece = 0;
      }
    },

    loadFEN(FEN:string) {
      // clear the board
      this.clearBoard()

      const split_string = FEN.split(' ');

      const FENBoard = split_string[0];
      let index = 0;
      FENBoard.split('').forEach((token) => {
        if (token === '/' || token === ' ')
          return;
        const emptySquares = Number(token);
        if (Number.isNaN(emptySquares)) {
          this.board[index].piece = FENByteToNumber(token);
          index++;
          return;
        }
        index += emptySquares;
      });

      if (split_string.length < 2)
        return;
      if (split_string[1] === 'w') {
        this.nextToPlay = 0;
      } else {
        this.nextToPlay = 1;
      }

      if (split_string.length < 3)
        return;
      this.castling = LoadCastlingFEN(split_string[2]);

      if (split_string.length < 4)
        return;
      this.enPassant = ChessNotationToIndex(split_string[3]);

      if (split_string.length < 5)
        return;
      this.halfClock = Number(split_string[4]);

      if (split_string.length < 6)
        return;
      this.fullClock = Number(split_string[5]);
    }, 

    generateFEN() {
      let fen = '';
      let empty = 0;
      this.board.forEach((square) => {
        if (square.piece !==0) {
          if (empty !== 0) {
            fen += `${empty}`;
          }
          fen += NumberToFENByte(square.piece);
          empty = 0;
        } else {
          empty++;
        }
        if (square.id % 8 === 7 && square.id !== 63) {
          if (empty !== 0) {
            fen += `${empty}`;
          }
          fen += '/';
          empty = 0;
        }
      })

      if (this.nextToPlay === 0) {
        fen += ' w';
      } else {
        fen += ' b';
      }

      // add castling to FEN 
      fen += ` ${GenerateCastlingFEN(this.castling)}`;

      // add en passant to FEN
      fen += ` ${IndexToChessNotation(this.enPassant)}`;

      // add half clock
      fen += ` ${this.halfClock}`;

      // add full clock
      fen += ` ${this.fullClock}`;

      this.currentFEN = fen;
    },

    async fetchMoves() {
      this.moves = []
      try {
        const res = await fetch( 
          "http://localhost:12345/FEN", {
            method: "POST", // *GET, POST, PUT, DELETE, etc.
            mode: "cors", // no-cors, *cors, same-origin
            body: this.currentFEN, // body data type must match "Content-Type" header
          }
        );
        const json = await res.json();
        console.log(json);
        this.moves = JSON.parse(JSON.stringify(json.Data));
        this.playing = json.Playing;
        this.winner = json.Winner;
        const fen = json.FEN;
        if (fen !== "") {
          this.loadFEN(fen)
        }
      } catch (error) {
        console.log(error);
      }
    }
  },
});
</script>