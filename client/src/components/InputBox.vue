<template>
  <div class="input-container">
    <el-button class="button" type="primary" @click="newRoom"
      >New Room</el-button
    >
    <el-input class="input" v-model="input" placeholder="Please input" />
    <el-button class="button" type="primary" @click="send">Send</el-button>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
const input = ref("");

const newRoom = () => {
  console.log("new room");
};

const send = () => {
  console.log(input.value);
};
</script>

<style lang="scss">
.input-container {
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: center;

  .input {
    height: 50px;
    width: 50%;
    align-self: center;
  }

  .button {
    height: 50px;
    width: 100px;
    margin-left: 20px;
    margin-right: 20px;
    align-self: center;
  }
}
</style>
