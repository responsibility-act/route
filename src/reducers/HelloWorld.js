import * as type from "../actions/type";

const helloWorld = (state, action) => {
  if (!state) {
    state = {
      text: ""
    };
  }
  switch (action.type) {
    case type.HELLO_WORLD_CHANGE:
      return { ...state, text: action.data };
    default:
      return state;
  }
};

export default helloWorld;
