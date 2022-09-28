import { Player, Room } from "@/common/models";
import {
  ActionTree,
  Commit,
  createStore, MutationTree
} from "vuex";

interface State {
  ws: WebSocket | null;
  messages: Array<string>;
  allRooms: Array<Room>;
  hallPlayers: Array<Player>;
  roomPlayers: Array<Player>
}

const state: State = {
  ws: null,
  messages: [] as string[],
  allRooms: [] as Room[],
  hallPlayers: [] as Player[],
  roomPlayers: [] as Player[],
};

const MutationCommands = {
  SETUP_WS: "setupWebSocket",
  NEW_MESSAGE: "newMessage",
  ALL_ROOMS: "allRoomsRefresh",
  HALL_PLAYERS: "hallPlayersRefresh",
  ROOM_PLAYERS: "roomPlayersRefresh",
};

const mutations: MutationTree<State> = {
  setupWebSocket(state: State, ws: WebSocket) {
    state.ws = ws;
  },

  newMessage(state: State, msg: string) {
    state.messages.push(msg);
  },

  allRoomsRefresh(state: State, rooms: Room[]) {
    state.allRooms = rooms;
  },

  hallPlayersRefresh(state: State, players: Player[]) {
    state.hallPlayers = players;
  },

  roomPlayersRefresh(state: State, players: Player[]) {
    state.roomPlayers = players;
  }
};

export const ActionCommands = {
  SETUP_WS: "setupWebSocket",
  NEW_MESSAGE: "newMessage",
  ALL_ROOMS: "allRoomsRefresh",
  HALL_PLAYERS: "hallPlayersRefresh",
  ROOM_PLAYERS: "roomPlayersRefresh",
};

const actions: ActionTree<State, any> = {
  setupWebSocket(context: { commit: Commit }, ws: WebSocket) {
    context.commit(MutationCommands.SETUP_WS, ws);
  },

  newMessage(context: { commit: Commit }, msg: string) {
    context.commit(MutationCommands.NEW_MESSAGE, msg);
  },

  allRoomsRefresh(context: { commit: Commit }, rooms: Room[]) {
    context.commit(MutationCommands.ALL_ROOMS, rooms);
  },

  hallPlayersRefresh(context: { commit: Commit }, players: Player[]) {
    context.commit(MutationCommands.HALL_PLAYERS, players);
  },

  roomPlayersRefresh(context: { commit: Commit }, players: Player[]) {
    context.commit(MutationCommands.ROOM_PLAYERS, players);
  }
};


export default createStore({
  state,
  mutations,
  actions,
});
