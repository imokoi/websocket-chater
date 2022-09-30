<template>
  <div class="home">
    <el-container class="container">
      <el-header><h2>Chatting Hall</h2></el-header>
      <el-container>
        <div style="margin-top: 20px">
          <el-aside class="side">
            <ListView
                list-type="player"
                :player-list=store.state.hallPlayers
            />
          </el-aside>
          <el-aside class="side" style="margin-top: 20px">
            <ListView
                list-type="room"
                :room-list=store.state.allRooms
                @join-room=joinRoom
            />
          </el-aside>
        </div>
        <el-container class="main-container">
          <el-main class="message-box">
            <MessageView
                :messages=store.state.messages
            />
          </el-main>
          <el-footer class="input-box">
            <InputBox
                @send-message=sendMessage
                @new-room=newRoom
            />
          </el-footer>
        </el-container>
      </el-container>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import {MessageCode} from "@/common/define";
import InputBox from "@/components/InputBox.vue";
import ListView from "@/components/ListView.vue";
import MessageView from "@/components/MessageView.vue";
import router from "@/router";
import store, {ActionCommands} from "@/store";

const ws = store.state.ws;

const sendMessage = (msg: string) => {
  if (!ws) return;
  ws.send(JSON.stringify({
    code: MessageCode.HallChatRequest,
    data: msg
  }));
};

const newRoom = () => {
  if (!ws) return;
  if (ws.OPEN === 1) {
    ws.send(JSON.stringify({
      code: MessageCode.NewRoomRequest,
      data: ""
    }));
  }
};

const joinRoom = (id: string) => {
  if (!ws) return;
  if (ws.OPEN === 1) {
    ws.send(JSON.stringify({
      code: MessageCode.JoinRoomRequest,
      data: id
    }));
  }
  store.dispatch(ActionCommands.SET_CURRENT_ROOM, id);
  router.push({name: "chat-room"});
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
  height: 400px;

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
