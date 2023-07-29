<template>
  <div class="bg-slate-300 h-screen w-full flex justify-center items-center">
    <ChessBoard/>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import ChessBoard from './components/ChessBoard.vue';

export default defineComponent({
  name: 'App',
  components: {
    ChessBoard
  }
});
</script>
