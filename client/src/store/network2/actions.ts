import {journaledOperation, LIST_FAILURE, LIST_REQUEST, LIST_SUCCESS} from "redux-data-connect";
import {createAction} from "redux-api-middleware";
import {endpoint, NETWORK2_PREFIX} from "./index";

export const importReceivable = (id: string, opHash: string) =>
    journaledOperation(
        createAction({
            endpoint: `${endpoint}sync/`,
            method: "POST",
            headers: ({ "Content-Type": "application/json" }),
            body: JSON.stringify({id}),
            types: [
                NETWORK2_PREFIX + "SYNC_REQUEST",
                NETWORK2_PREFIX + "SYNC_SUCCESS",
                NETWORK2_PREFIX + "SYNC_FAILURE",
            ],
        }),
        opHash
    );

export const getListReceivables = (opHash: string) =>
    journaledOperation(
        createAction({
            endpoint: endpoint,
            method: "GET",
            headers: ({ "Content-Type": "application/json" }),
            types: [
                NETWORK2_PREFIX + LIST_REQUEST,
                NETWORK2_PREFIX + LIST_SUCCESS,
                NETWORK2_PREFIX + LIST_FAILURE,
            ],
        }),
        opHash
    );
