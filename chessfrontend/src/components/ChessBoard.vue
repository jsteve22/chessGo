<template>
  <div>
    <div>
      drag started in {{ startDrag }}
    </div>
    <div class="grid grid-cols-8" id="ChessBoardGrid">
      <div v-for="color in squares" v-bind:key="color.id">
        <chessboardSquare :color=color.color :id="(startDrag == -1 ) ? color.id : startDrag" :piece=color.piece :placable="placable[color.id]"
        @dragUpdate="(startSquare:number) => startDrag = startSquare" @pieceType="(piece:number) => pieces[color.id] = piece"/>
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
      elements.push( {color: color, id: i, piece: 0} );
    }
    this.squares = elements
    this.squares[0].piece = 1 + 8;
  },
  data() {
    interface squareValues {
      color: number,
      id: number,
      piece: number,
    }
    const empty:(squareValues)[] = [];
    const placable = Array(64).fill(false);
    const pieces = Array(64).fill(0);
    return {
      squares: empty,
      startDrag: -1,
      placable: placable,
      pieces: pieces,
    }
  }, 
  updated() {
    for (let i = 0; i < 64; i++) {
      const isPlacable = (i !== this.startDrag) && (this.pieces[i] === 0) && (this.startDrag !== -1);
      this.placable[i] = (isPlacable) ? true : false;
    }
    if (this.startDrag !== -1)
      this.pieces[this.startDrag] = 0;
  },
});
</script>