import React, { useState } from "react";
import { Button } from "react-bootstrap";
import { DataListView } from "rn-web-components";
import { useDispatch, useSelector } from "react-redux";
import actions from "../store/actions";
import { receivables1ListSelector } from "../store/network1";
import {
  ReceivableHeader,
  ReceivableListItem1,
} from "../containers/ReceivableListItem";
import { BACKEND_ADDR_1 } from "../config";
import { Receivable } from "../core/receivable";
import { ReceivablesModal } from "../containers/ReceivablesModal";

export function Network1Form(): React.ReactElement {
  const dispatch = useDispatch();
  const [receivable, setReceivable] = useState<Receivable | null>(null);
  const [showModal, setShowModal] = useState(false);

  const getList = (opHash: string) =>
    dispatch(actions.network1.getListReceivables(opHash));
  const data = useSelector(receivables1ListSelector);

  const onSelect = (id: string) => {
    const found = data?.filter((e) => e.id === id);
    if (found.length === 0) return;
    setReceivable(found[0]);
    setShowModal(true);
  };

  const onHide = () => setShowModal(false);

  return (
    <>
      <ReceivablesModal data={receivable} show={showModal} onHide={onHide} />
      <div
        style={{
          backgroundColor: "#bee8ff",
          borderRadius: "20px",
          padding: "20px",
          textAlign: "center",
          marginTop: "30px",
        }}
      >
        <h1>Network #1</h1>
        <h3 style={{ marginBottom: 0 }}>Oil supplier network</h3>
        <span style={{ fontSize: "14px", fontWeight: "bold" }}>
          {BACKEND_ADDR_1}
        </span>

        <div style={{ flex: 1, height: "calc(100vh - 360px)" }}>
          <DataListView
            getList={getList}
            data={data}
            renderHeader={ReceivableHeader}
            renderItem={ReceivableListItem1}
            onSelect={onSelect}
          />
        </div>
        <Button onClick={() => getList("Update")}>Refresh List</Button>
      </div>
    </>
  );
}
