<template>
  <div class="home">
    <h2>Chatting Room</h2>
    <el-container class="container">
      <div class="side-box">
        <div class="side">
          <ListView
              list-type="player"
              :player-list=store.state.roomPlayers
          />
        </div>
      </div>
      <div class="main-container">
        <el-main class="message-box">
          <MessageView
              :messages=store.state.roomMessages
          />
        </el-main>
        <el-footer class="input-box">
          <InputBox
              @send-message=sendMessage
          />
        </el-footer>
      </div>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { MessageCode } from "@/common/define";
import InputBox from "@/components/InputBox.vue";
import store from "@/store";
import ListView from "../components/ListView.vue";
import MessageView from "../components/MessageView.vue";

const ws = store.state.ws;
const sendMessage = (msg: string) => {
  if (!ws) return;
  ws.send(JSON.stringify({
    code: MessageCode.RoomChatRequest,
    data: { id: store.state.currentRoomId, message: msg }
  }));
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
    height: 400px;
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
  }
}
</style>
