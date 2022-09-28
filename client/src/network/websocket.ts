import { MessageCode } from "@/common/define";
import { Room } from "@/common/models";
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

// const getAllRooms = () => {
//   if (!ws) return;
//   ws.send(JSON.stringify({
//     code: MessageCode.AllRooms
//   }));
// };

const onOpen = (event: Event) => {
  console.log("Connected to server");
  const ws = store.state.ws;
  if (!ws) return;
  // get all rooms
  ws.send(JSON.stringify({
    code: MessageCode.AllRooms
  }));
};

const onClose = (event: CloseEvent) => {
  console.log("connection closed");
};

const onMessage = (event: MessageEvent) => {
  console.log(event.data);
  try {
    const message = JSON.parse(event.data);
    switch (message.code) {
      case MessageCode.HallChat:
        store.dispatch(ActionCommands.NEW_MESSAGE, message.data);
        break;
      case MessageCode.Success:
        store.dispatch(ActionCommands.NEW_MESSAGE, message.data);
        break;
      case MessageCode.AllRooms:
        handleAllRoomsMessage(message.data);
        break;
      default:
        console.log("Unknown message code");
    }
  } catch (e) {
    console.log("Invalid message");
  }
};

const onError = (event: Event) => {
  console.log(event);
};


function handleAllRoomsMessage(data: any) {
  console.log(data);
  const rooms: Room[] = data;
  console.log(rooms);
  store.dispatch(ActionCommands.ALL_ROOMS, rooms);
}
