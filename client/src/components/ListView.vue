<template>
  <div>
    <el-scrollbar class="scrollbar">
      <div v-if="listType==='room'">
        <p v-for="item in roomList" :key="item.id" class="scrollbar-item">
          {{ item.id }}
        </p>
      </div>
      <div v-else>
        <p v-for="item in playerList" :key="item.id" class="scrollbar-item">
          {{ item.name }}
        </p>
      </div>
    </el-scrollbar>
  </div>
</template>

<script setup lang="ts">
import { Player, Room } from "@/common/models";
import { defineProps } from "vue";

defineProps({
  listType: {
    type: String,
    default: "room"
  },
  roomList: {
    type: Array as () => Room[],
    default: [] as Room[],
    required: false
  },
  playerList: {
    type: Array as () => Player[],
    default: [] as Player[],
    required: false
  }
});
</script>

<style lang="scss">
.scrollbar {
  height: 100%;
  width: 80%;
  background-color: antiquewhite;
}

.scrollbar-item {
  font-size: 10px;
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
