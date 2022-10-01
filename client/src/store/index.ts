import {DialogMessage, Player, Room} from "@/common/models";
import {
  ActionTree,
  Commit,
  createStore, MutationTree
} from "vuex";

interface State {
  ws: WebSocket | null;
  messages: Array<DialogMessage>;
  allRooms: Array<Room>;
  hallPlayers: Array<Player>;
  currentRoomId: string;
  roomPlayers: Array<Player>;
  roomMessages: Array<DialogMessage>;
}

const state: State = {
  ws: null,
  messages: [] as DialogMessage[],
  allRooms: [] as Room[],
  hallPlayers: [] as Player[],
  currentRoomId: "",
  roomPlayers: [] as Player[],
  roomMessages: [] as DialogMessage[],
};

const MutationCommands = {
  SETUP_WS: "setupWebSocket",
  NEW_MESSAGE: "newMessage",
  ALL_ROOMS: "allRoomsRefresh",
  HALL_PLAYERS: "hallPlayersRefresh",
  SET_CURRENT_ROOM: "setCurrentRoom",
  ROOM_PLAYERS: "roomPlayersRefresh",
  ROOM_MESSAGES: "roomMessagesRefresh",
};

const mutations: MutationTree<State> = {
  setupWebSocket(state: State, ws: WebSocket) {
    state.ws = ws;
  },

  newMessage(state: State, msg: DialogMessage) {
    state.messages.push(msg);
  },

  allRoomsRefresh(state: State, rooms: Room[]) {
    state.allRooms = rooms;
  },

  setCurrentRoom(state: State, roomId: string) {
    state.currentRoomId = roomId;
  },

  hallPlayersRefresh(state: State, players: Player[]) {
    state.hallPlayers = players;
  },

  roomPlayersRefresh(state: State, players: Player[]) {
    state.roomPlayers = players;
  },

  roomMessagesRefresh(state: State, message: DialogMessage) {
    state.roomMessages.push(message);
  }
};

export const ActionCommands = {
  SETUP_WS: "setupWebSocket",
  NEW_MESSAGE: "newMessage",
  ALL_ROOMS: "allRoomsRefresh",
  HALL_PLAYERS: "hallPlayersRefresh",
  SET_CURRENT_ROOM: "setCurrentRoom",
  ROOM_PLAYERS: "roomPlayersRefresh",
  ROOM_MESSAGES: "roomMessagesRefresh",
};

const actions: ActionTree<State, any> = {
  setupWebSocket(context: { commit: Commit }, ws: WebSocket) {
    context.commit(MutationCommands.SETUP_WS, ws);
  },

  newMessage(context: { commit: Commit }, msg: DialogMessage) {
    context.commit(MutationCommands.NEW_MESSAGE, msg);
  },

  allRoomsRefresh(context: { commit: Commit }, rooms: Room[]) {
    context.commit(MutationCommands.ALL_ROOMS, rooms);
  },

  setCurrentRoom(context: { commit: Commit }, roomId: string) {
    context.commit(MutationCommands.SET_CURRENT_ROOM, roomId);
  },

  hallPlayersRefresh(context: { commit: Commit }, players: Player[]) {
    context.commit(MutationCommands.HALL_PLAYERS, players);
  },

  roomPlayersRefresh(context: { commit: Commit }, players: Player[]) {
    context.commit(MutationCommands.ROOM_PLAYERS, players);
  },

  roomMessagesRefresh(context: { commit: Commit }, messages: DialogMessage) {
    context.commit(MutationCommands.ROOM_MESSAGES, messages);
  }
};


export default createStore({
  state,
  mutations,
  actions,
});
