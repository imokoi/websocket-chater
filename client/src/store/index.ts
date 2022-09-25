import {
  ActionTree,
  Commit,
  createStore, MutationTree
} from "vuex";

interface State {
  ws: WebSocket | null;
  messages: Array<string>;
}

const state: State = {
  ws: null,
  messages: [] as string[],
};

const MutationCommands = {
  SETUP_WS: "setupWebSocket",
  NEW_MESSAGE: "newMessage",
};

const mutations: MutationTree<State> = {
  setupWebSocket(state: State, ws: WebSocket) {
    state.ws = ws;
  },

  newMessage(state: State, msg: string) {
    state.messages.push(msg);
  }
};

export const ActionCommands = {
  SETUP_WS: "setupWebSocket",
  NEW_MESSAGE: "newMessage",
};

const actions: ActionTree<State, any> = {
  setupWebSocket(context: { commit: Commit }, ws: WebSocket) {
    context.commit(MutationCommands.SETUP_WS, ws);
  },

  newMessage(context: { commit: Commit }, msg: string) {
    context.commit(MutationCommands.NEW_MESSAGE, msg);
  }
};


export default createStore({
  state,
  mutations,
  actions,
});
