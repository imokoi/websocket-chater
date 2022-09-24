<template>
  <div>
    <el-scrollbar class="scrollbar">
      <p v-for="item in 50" :key="item" class="scrollbar-item">
        {{ item }}
      </p>
    </el-scrollbar>
  </div>
</template>

<script setup lang="ts"></script>

<style lang="scss">
.scrollbar {
  height: 100%;
  width: 80%;
  background-color: antiquewhite;
}

.scrollbar-item {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 50px;
  margin: 10px;
  text-align: center;
  border-radius: 4px;
  background: var(--el-color-primary-light-9);
  color: var(--el-color-primary);
}
</style>
