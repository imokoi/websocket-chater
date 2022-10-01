<template>
  <div class="message-view-container">
    <el-scrollbar height="400px">
      <p v-for="(item, index) in messages" :key=index class="message-scrollbar-item">
        {{ item }}
      </p>
    </el-scrollbar>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from "vue";

defineProps({
  messages: {
    type: Array,
    default: [] as string[]
  }
});
</script>

<style lang="scss">
.message-view-container {
  height: 100%;
  width: 100%;
}

.message-scrollbar-item {
  height: auto;
  margin: 10px;
  text-align: left;
  border-radius: 4px;
}
</style>
