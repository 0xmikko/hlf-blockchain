import {getFullUrl} from "redux-data-connect";
import {BACKEND_ADDR_2} from "../../config";
import {RootState} from "../index";

export const NETWORK2_PREFIX = 'N2_RECEIVABLES@@';
export const endpoint = getFullUrl('/api/receivables/', {host: BACKEND_ADDR_2});

export const receivables2ListSelector = (state: RootState) => state.receivables_2.List.data
