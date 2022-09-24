import {
  ActionTree,
  Commit,
  createStore,
  GetterTree,
  MutationTree
} from "vuex";

interface State {
  ws: WebSocket | null;
}

const state: State = {
  ws: null,
};

const MutationCommands = {
  SETUP_WS: "setupWebSocket",
};

const mutations: MutationTree<State> = {
  setupWebSocket(state: State, ws: WebSocket) {
    state.ws = ws;
  },
};

export const ActionCommands = {
  SETUP_WS: "setupWebSocket",
};

const actions: ActionTree<State, any> = {
  setupWebSocket(context: { commit: Commit }, ws: WebSocket) {
    context.commit(MutationCommands.SETUP_WS, ws);
  },
};

export const GetterCommands = {
  GET_WS: "getWebSocket",
};

const getters: GetterTree<State, any> = {
  getWebSocket(state: State): WebSocket | null {
    return state.ws;
  },
};

export default createStore({
  state,
  getters,
  mutations,
  actions,
});
