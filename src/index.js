import React from "react";
import ReactDOM from "react-dom";
import CRouter from "./routes";
import { Provider } from "react-redux";

import store from "./store";

ReactDOM.render(
  <Provider store={store}>
    <CRouter store={store} />
  </Provider>,
  document.getElementById("root")
);
