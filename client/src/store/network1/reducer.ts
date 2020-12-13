import { createDataLoaderReducer } from "redux-data-connect";
import { Receivable } from "../../core/receivable";
import { NETWORK1_PREFIX } from "./";

export default createDataLoaderReducer<Receivable>(NETWORK1_PREFIX);
