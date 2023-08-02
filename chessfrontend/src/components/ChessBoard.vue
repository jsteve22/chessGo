<template>
  <div>
    <div>
      drag started in {{ startDrag }}
      drag ended in {{ endDrag }}
      <button @click="fetchMoves" class="w-auto h-auto bg-emerald-200 rounded-lg p-1">
        fetch Moves
      </button>
      <br/>
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

export default defineComponent({
  name: 'ChessBoard',
  components: {
    chessboardSquare
  },
  mounted() {
    const size = 8*8;
    const elements = [];
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
    dropDraggable() {
      console.log(`startDrag = ${this.startDrag}`);
      console.log(`endDrag = ${this.endDrag}`);
      if (this.endDrag === this.startDrag || this.board[this.endDrag].placable === false) {
        this.startDrag = -1;
        return;
      }
      if (this.startDrag !== -1) {
        this.board[this.endDrag].piece = this.board[this.startDrag].piece;
        this.board[this.startDrag].piece = 0;
        // console.log(this.board);
        this.nextToPlay = this.nextToPlay ^ 1;
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
          this.board[index].piece = this.FENByteToNumber(token);
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

    }, 
    generateFEN() {
      let fen = '';
      let empty = 0;
      this.board.forEach((square) => {
        if (square.piece !==0) {
          if (empty !== 0) {
            fen += `${empty}`;
          }
          fen += this.NumberToFENByte(square.piece);
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

      this.currentFEN = fen;
    },
    FENByteToNumber(token:string) {
      switch (token) {
        case 'P':
          return 1;
        case 'N':
          return 2;
        case 'B':
          return 3;
        case 'R':
          return 4;
        case 'Q':
          return 5;
        case 'K':
          return 6;
        case 'p':
          return 1 + 8;
        case 'n':
          return 2 + 8;
        case 'b':
          return 3 + 8;
        case 'r':
          return 4 + 8;
        case 'q':
          return 5 + 8;
        case 'k':
          return 6 + 8;
      }
      return 0;
    },
    NumberToFENByte(piece:number) {
      switch (piece) {
        case 1:
          return 'P';
        case 2:
          return 'N';
        case 3:
          return 'B';
        case 4:
          return 'R';
        case 5:
          return 'Q';
        case 6:
          return 'K';
        case 1 + 8:
          return 'p';
        case 2 + 8:
          return 'n';
        case 3 + 8:
          return 'b';
        case 4 + 8:
          return 'r';
        case 5 + 8:
          return 'q';
        case 6 + 8:
          return 'k';
      }
      return '';
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
        this.moves = JSON.parse(JSON.stringify(json.Data));
      } catch (error) {
        console.log(error);
      }
    }
  },
});
</script>