import store, { ActionCommands } from "@/store";

const socketUrl = "ws://localhost:8888/ws";

export function initWebsocket() {
  const ws = new WebSocket(socketUrl);
  store.dispatch(ActionCommands.SETUP_WS, ws);
  ws.onopen = onOpen;
  ws.onclose = onClose;
  ws.onmessage = onMessage;
  ws.onerror = onError;
}

const onOpen = (event: Event) => {
  console.log("Connected to server");
  const ws = store.state.ws;
  if (ws) {
    ws.send("Hello from client");
  }
};

const onClose = (event: CloseEvent) => {
  console.log("connection closed");
};

const onMessage = (event: MessageEvent) => {
  const data = JSON.parse(event.data);
  console.log(data);
};

const onError = (event: Event) => {
  console.log(event);
};
