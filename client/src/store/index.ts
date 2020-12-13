/*
 * Copyright (c) 2020. Mikael Lazarev
 */

import { applyMiddleware, compose, createStore } from "redux";
import reducer from "./reducer";
import thunk from "redux-thunk";
import { composeWithDevTools } from "redux-devtools-extension";
import {createMiddleware} from 'redux-api-middleware';

export type RootState = ReturnType<typeof reducer>;

let composeEnhancers: typeof compose;

if (!process.env.NODE_ENV || process.env.NODE_ENV === "development") {
  composeEnhancers = composeWithDevTools({});
} else {
  composeEnhancers = compose;
}

export default function configureStore() {
  return createStore(
    reducer,
    composeEnhancers(applyMiddleware(thunk, createMiddleware()))
  );
}
