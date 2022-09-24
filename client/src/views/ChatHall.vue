<template>
  <div class="home">
    <el-container class="container">
      <el-header><h2>Imokoi Chatting Hall</h2></el-header>
      <el-container>
        <el-aside class="side">
          <ListView />
        </el-aside>
        <el-container class="main-container">
          <el-main class="message-box">
            <MessageView :messages=messages />
          </el-main>
          <el-footer class="input-box">
            <InputBox @send-message=sendMessage />
          </el-footer>
        </el-container>
        <el-aside class="side">
          <ListView />
        </el-aside>
      </el-container>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import InputBox from "@/components/InputBox.vue";
import ListView from "@/components/ListView.vue";
import MessageView from "@/components/MessageView.vue";
import store, { GetterCommands } from "@/store";
import { ref } from "vue";

const messages = ref([] as string[])

const ws = store.getters[GetterCommands.GET_WS];

if (!ws) {
  console.log("ws is not initialized");
}

const sendMessage = (msg: string) => {
  console.log(msg);
  messages.value.push(msg);
  ws.send(msg);
}
</script>

<style lang="scss">
.home {
  height: 100%;
  width: 100%;
}
.container {
  height: 100%;
  width: 100%;

  .side {
    height: 600px;
    width: 200px;
  }
}
.main-container {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  height: 80%;

  .message-box {
    flex: 2;
    overflow-y: scroll;
  }

  .input-box {
    display: flex;
    height: 100px;
    background-color: aliceblue;
  }
}
</style>
