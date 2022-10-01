<template>
  <div class="home">
    <h2>Chatting Hall</h2>
    <el-container class="container">
        <div class="side-box">
          <div class="side">
            <ListView
                list-type="player"
                :player-list=store.state.hallPlayers
            />
          </div>
          <div class="side" style="margin-top: 20px">
            <ListView
                list-type="room"
                :room-list=store.state.allRooms
                @join-room=joinRoom
            />
          </div>
        </div>
        <div class="main-container">
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
        </div>
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
  display: flex;
  flex-direction: column;
  background: antiquewhite;

  .container {
    height: 80%;
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: center;

    .side-box {
      display: flex;
      flex-direction: column;
      justify-content: flex-start;
      .side {
        height: 40%;
        border: #42b983 2px solid;
        overflow: auto;
      }
    }

    .main-container {
      height: 83%;
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      margin-left: 10px;
      width: 60%;

      .message-box {
        width: 100%;
        border: #42b983 2px solid;
        overflow: auto;
      }

      .input-box {
        border: #42b983 2px solid;
        width: 100%;
        margin-top: 30px;
        align-self: center;
        display: flex;
      }
    }
  }
}
</style>
