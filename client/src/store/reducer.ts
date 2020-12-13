/*
 * Copyright (c) 2020. Mikael Lazarev
 */

import { combineReducers } from "redux";
import receivables_1 from './network1/reducer'
import receivables_2 from './network1/reducer'
import { operationReducer } from "redux-data-connect";

export default combineReducers({
    receivables_1,
    receivables_2,
    operations: operationReducer,
});
