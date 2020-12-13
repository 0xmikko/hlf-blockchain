import {journaledOperation, LIST_FAILURE, LIST_REQUEST, LIST_SUCCESS} from "redux-data-connect";
import {createAction} from "redux-api-middleware";
import {endpoint, NETWORK1_PREFIX} from "./index";

export const createReceivable = (id: string, opHash: string) =>
    journaledOperation(
        createAction({
            endpoint: endpoint,
            method: "POST",
            headers: ({ "Content-Type": "application/json" }),
            body: "",
            types: [
                NETWORK1_PREFIX + "CREATE_REQUEST",
                NETWORK1_PREFIX + "CREATE_SUCCESS",
                NETWORK1_PREFIX + "CREATE_FAILURE",
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
                NETWORK1_PREFIX + LIST_REQUEST,
                NETWORK1_PREFIX + LIST_SUCCESS,
                NETWORK1_PREFIX + LIST_FAILURE,
            ],
        }),
        opHash
    );
