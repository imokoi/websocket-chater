<template>
  <div>
    <el-scrollbar height="800px">
      <p v-for="(item, index) in messages" :key=index class="message-scrollbar-item">
        {{ item }}
      </p>
    </el-scrollbar>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue';

defineProps({
  messages: {
    type: Array,
    default: [] as string[],
  },
});
</script>

<style lang="scss">
.message-scrollbar-item {
  height: auto;
  margin: 10px;
  text-align: left;
  border-radius: 4px;
  background-color: aliceblue;
}
</style>
