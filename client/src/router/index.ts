import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import ChatHall from "../views/ChatHall.vue";
import ChatRoom from "../views/ChatRoom.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "chat-hall",
    component: ChatHall,
  },
  {
    path: "/chat-room",
    name: "chat-room",
    component: ChatRoom,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
