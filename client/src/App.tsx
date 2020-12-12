import React from "react";
import { Provider } from "react-redux";

import { TransferScreen } from "./screens/TransferScreen";
import configureStore from "./store";

const store = configureStore();

const App = () => {
  return (
    <Provider store={store}>
      <TransferScreen />
    </Provider>
  );
};

export default App;
