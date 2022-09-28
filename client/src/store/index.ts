import { Room } from "@/common/models";
import {
  ActionTree,
  Commit,
  createStore, MutationTree
} from "vuex";

interface State {
  ws: WebSocket | null;
  messages: Array<string>;
  allRooms: Array<Room>;
}

const state: State = {
  ws: null,
  messages: [] as string[],
  allRooms: [] as Room[],
};

const MutationCommands = {
  SETUP_WS: "setupWebSocket",
  NEW_MESSAGE: "newMessage",
  ALL_ROOMS: "allRoomsRefresh",
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
  }
};

export const ActionCommands = {
  SETUP_WS: "setupWebSocket",
  NEW_MESSAGE: "newMessage",
  ALL_ROOMS: "allRoomsRefresh",
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
  }
};


export default createStore({
  state,
  mutations,
  actions,
});
