import {getFullUrl} from "redux-data-connect";
import {BACKEND_ADDR_1} from "../../config";
import {RootState} from "../index";

export const NETWORK1_PREFIX = 'N1_RECEIVABLES@@';
export const endpoint = getFullUrl('/api/receivables/', {host: BACKEND_ADDR_1});

export const receivables1ListSelector = (state: RootState) => state.receivables_1.List.data
