import React, { useEffect, useState } from "react";
import { Receivable } from "../core/receivable";
import { Button, Table } from "react-bootstrap";
import { useDispatch, useSelector } from "react-redux";
import actions from "../store/actions";
import { operationSelector } from "redux-data-connect";

export interface ReceivableListItemProps {
  data: Receivable;
  onSelect: (id: string) => void;
}

export function ReceivableListItem1({
  data,
  onSelect,
}: ReceivableListItemProps): React.ReactElement {
  const dispatch = useDispatch();
  const [hash, setHash] = useState("0");
  const [btnName, setBtnName] = useState("Transfer");

  const onSync = () => {
    const newHash = Date.now().toString();
    dispatch(actions.network2.importReceivable(data.id, newHash));
    setHash(newHash);
  };

  const operation = useSelector(operationSelector(hash));

  useEffect(() => {
    if (hash !== "0") {
      switch (operation?.status) {
        default:
        case "STATUS.LOADING":
          setBtnName("Transferring...");
          break;

        case "STATUS.FAILURE":
          alert(
            "Receivable transfer failed" + (operation.error || "Unknown error")
          );
          setBtnName("Failed");
          setHash("0");
          break;

        case "STATUS.UPDATING":
        case "STATUS.SUCCESS":
          alert("Receivable transferred");
          setBtnName("Transferred");
          setHash("0");
          break;
      }
    }
  });

  return (
    <tr>
      <td
        className="text-left tx-normal"
        style={{ height: "60px" }}
        onClick={() => onSelect(data.id)}
      >
        {data.issuer}
      </td>
      <td className="text-left tx-normal" onClick={() => onSelect(data.id)}>
        {data.payer}
      </td>
      <td className="text-center tx-normal" onClick={() => onSelect(data.id)}>
        {data.amount}
      </td>
      <td>
        <Button size={"sm"} onClick={onSync}>
          {btnName}
        </Button>
      </td>
    </tr>
  );
}

export function ReceivableListItem2({
  data,
  onSelect,
}: ReceivableListItemProps): React.ReactElement {
  return (
    <tr onClick={() => onSelect(data.id)}>
      <td className="text-left tx-normal" style={{ height: "60px" }}>
        {data.issuer}
      </td>
      <td className="text-left tx-normal">{data.payer}</td>
      <td className="text-center tx-normal">{data.amount}</td>
      <td></td>
    </tr>
  );
}

export function ReceivableHeader({
  children,
}: React.PropsWithChildren<any>): React.ReactElement {
  return (
    <Table
      className="table-dashboard"
      hover={true}
      style={{
        color: "black",

        marginTop: "20px",
      }}
    >
      <thead style={{ borderColor: "black!important" }}>
        <th className="text-left tx-normal">Issuer</th>
        <th className="text-left tx-normal">Payer</th>
        <th>Amount</th>
        <th>Action</th>
      </thead>
      <tbody>{children}</tbody>
    </Table>
  );
}
