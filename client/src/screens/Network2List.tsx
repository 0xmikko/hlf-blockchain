import React, { useState } from "react";
import { Button } from "react-bootstrap";
import { useDispatch, useSelector } from "react-redux";
import actions from "../store/actions";
import { receivables2ListSelector } from "../store/network2";
import { DataListView } from "rn-web-components";
import {
  ReceivableHeader,
  ReceivableListItem2,
} from "../containers/ReceivableListItem";
import { BACKEND_ADDR_2 } from "../config";
import { Receivable } from "../core/receivable";
import { ReceivablesModal } from "../containers/ReceivablesModal";

export function Network2List(): React.ReactElement {
  const dispatch = useDispatch();
  const [receivable, setReceivable] = useState<Receivable | null>(null);
  const [showModal, setShowModal] = useState(false);

  const getList = (opHash: string) =>
    dispatch(actions.network2.getListReceivables(opHash));
  const data = useSelector(receivables2ListSelector);

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
          backgroundColor: "#CCCCCC",
          borderRadius: "20px",
          padding: "20px",
          textAlign: "center",
          marginTop: "30px",
        }}
      >
        <h1>Network #2</h1>
        <h3 style={{ marginBottom: 0 }}>Factoring network</h3>
        <span style={{ fontSize: "14px", fontWeight: "bold" }}>
          {BACKEND_ADDR_2}
        </span>
        <div style={{ flex: 1, height: "calc(100vh - 360px)" }}>
          <DataListView
            getList={getList}
            data={data}
            renderHeader={ReceivableHeader}
            renderItem={ReceivableListItem2}
            onSelect={onSelect}
          />
        </div>
        <Button onClick={() => getList("Update")}>Refresh List</Button>
      </div>
    </>
  );
}
