<template>
  <nav>
    <router-link to="/">Chat Hall</router-link> |
    <router-link to="/chat-room">Chat room</router-link>
  </nav>
  <router-view />
</template>

<script setup lang="ts">
const ws = new WebSocket("ws://localhost:8888/ws");

const onMessage = function (event: MessageEvent) {
  console.log(event.data);
};

const onError = function (event: Event) {
  console.log(event);
};

ws.onopen = () => {
  console.log("Connected to server");
  ws.send("Hello from client");
};

ws.onclose = () => {
  console.log("connection closed");
};
ws.onmessage = onMessage;
ws.onerror = onError;
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  width: 100%;
  height: 100%;
}

html,
body {
  height: 100%;
  width: 100%;
}

nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
