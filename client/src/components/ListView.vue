<template>
  <div>
    <div v-if="listType==='player'">Player List:</div>
    <div v-else>Room List: </div>
    <div v-if="listType==='room'">
      <p v-for="item in roomList" :key="item.id" class="scrollbar-item">
        <el-button type="text" @click="joinRoom(item.id)">{{ item.id }}</el-button>
      </p>
    </div>
    <div v-else>
      <p v-for="item in playerList" :key="item.id" class="scrollbar-item">
        {{ item.name }}
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Player, Room } from "@/common/models";
import { defineEmits, defineProps } from "vue";

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

const emits = defineEmits(["join-room"]);
const joinRoom = (id: string) => {
  emits("join-room", id);
};
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
