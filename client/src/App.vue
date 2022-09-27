<template>
  <nav>
    <router-link to="/">Chat Hall</router-link>
    |
    <router-link to="/chat-room">Chat room</router-link>
  </nav>
  <router-view/>
</template>

<script setup lang="ts">
import { initWebsocket } from './network/websocket';

initWebsocket();
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
