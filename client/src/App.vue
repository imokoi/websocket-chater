<template>
  <nav>
    <router-link to="/">Home</router-link>
    |
    <router-link to="/about">About</router-link>
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
