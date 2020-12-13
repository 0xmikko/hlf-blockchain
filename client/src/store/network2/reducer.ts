import { createDataLoaderReducer } from "redux-data-connect";
import { Receivable } from "../../core/receivable";
import { NETWORK2_PREFIX } from "./";

export default createDataLoaderReducer<Receivable>(NETWORK2_PREFIX);
