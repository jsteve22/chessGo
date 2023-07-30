<template>
  <div :id="squareId" :class="squareStyling" >
    <img v-if="pieceType!==0" :id="pieceId" :src="getImagePath(pieceType)" class="active:cursor-grabbing" draggable="true" />
    <div v-if="placable && pieceType===0" class="w-1 h-1 bg-white">
    </div>
    <!--
    <div v-if="pieceType!==0" :id="pieceId" class="w-full h-full flex justify-center items-center" draggable="true" >
      <img :src="piecePath" draggable="true"/>
    </div>
    -->
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'chessboardSquare',
  props: {
    color: Number,
    id: Number,
    piece: Number,
    placable: Boolean,
  },
  data() {
    const id = this.id || 0;
    return {
      squareStyling: 'square h-24 w-24 select-none active:cursor-grabbing flex justify-center items-center bg-green-200',
      piecePath: '',
      pieceType: 0,
      squareId: `square-${id}`,
      pieceId: `piece-${id}`,
      sourceId: -1
    }
  }, 
  mounted() {
    // const light = 1;
    const dark = 1;
    if (this.color == dark) {
      this.squareStyling += ' bg-red-200';
    }

    this.pieceType = this.piece || 0

    const element = document.getElementById(this.squareId);
    element?.addEventListener('dragstart', this.startDraggable);
    element?.addEventListener('dragover', (event:DragEvent) => {event?.preventDefault();}, false);
    element?.addEventListener('dragenter', this.enterDraggable);
    element?.addEventListener('dragleave', this.leaveDraggable);
    element?.addEventListener('drop', this.dropDraggable);

  }, 
  updated() {
    const id = this.id || 0;
    this.sourceId = id;
    this.pieceType = this.piece || 0;
  },
  methods: {
    startDraggable() {
      this.$emit('dragStart', this.id);
      this.$emit('dragEnd', this.id);
    },
    enterDraggable(event:DragEvent) {
      const target = event.target;
      if (target instanceof Element) {
        if (target.classList.contains('square')) {
          target.classList.add('border-8');
          target.classList.add('border-amber-500');
        }
      }
      return;
    },
    leaveDraggable(event:DragEvent) {
      const target = event.target;
      if (target instanceof Element) {
        if (target.classList.contains('square')) {
          target.classList.remove('border-8');
          target.classList.remove('border-amber-500');
        }
      }
      return;
    },
    dropDraggable(event:DragEvent) {
      const source = document.getElementById(`square-${this.sourceId}`);
      const element = source?.children[0];
      event?.preventDefault();
      const target = event.target;

      if (target instanceof Element) {
        if (target.classList.contains('square')) {
          const targetIdNumber = Number(target.id.split('-')[1]);
          this.$emit('dragEnd', targetIdNumber);
          target.classList.remove('border-8');
          target.classList.remove('border-amber-500');
          // console.log('made it here in dropped');
          console.log(event);
          if (element && this.placable) {
            console.log(`dropped ${element.id} in ${target.id}`);
            // element.classList.add('border-4');
            // element.classList.add('border-blue-200');
            // element.remove();
            // target.appendChild(element);
            // element.id = 'piece-' + target.id.split('-')[1];
            // console.log(`new element id ${element.id}`);
          }
        }
        if (target.parentElement?.classList.contains('square')) {
          const targetIdNumber = Number(target.parentElement?.id.split('-')[1]);
          this.$emit('dragEnd', targetIdNumber);
          target.parentElement?.classList.remove('border-8');
          target.parentElement?.classList.remove('border-amber-500');
        }
      }
    },
    getImagePath(piece:number) {
      const white = 0; 
      const PAWN = 1; const KNIGHT = 2; const BISHOP = 3; 
      const ROOK = 4; const QUEEN = 5; const KING = 6;
      const pieceColor = piece >> 3;
      const pieceType = piece & 7;
      let image_name = (pieceColor == white) ? 'white_' : 'black_';
      switch (pieceType) {
        case PAWN:
          image_name += 'pawn.png';
          break
        case KNIGHT:
          image_name += 'knight.png';
          break
        case BISHOP:
          image_name += 'bishop.png';
          break
        case ROOK:
          image_name += 'rook.png';
          break
        case QUEEN:
          image_name += 'queen.png';
          break
        case KING:
          image_name += 'king.png';
          break
        }
        return require('@/assets/piece_pics/' + image_name);
        // console.log(this.piecePath);
    },
    testPieceChange() {
      this.pieceType++;
      if (this.pieceType === 7)
        this.pieceType = 9;
      if (this.pieceType === 7 + 8)
        this.pieceType = 0;
      this.$emit('pieceUpdate', this.pieceType);
    }
  },
});
</script>