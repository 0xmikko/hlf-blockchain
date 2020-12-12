import React from "react";
import { Provider } from "react-redux";
import { BrowserRouter } from "react-router-dom";

import { AuthSwitcher } from "./screens/AuthSwitcher";
import configureStore from "./store";


const store = configureStore();

const App = () => {
  return (
      <Provider store={store}>
        <BrowserRouter>
          <AuthSwitcher />
        </BrowserRouter>
      </Provider>
  );
};

export default App;
