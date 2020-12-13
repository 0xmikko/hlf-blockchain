import React from "react";
import { Receivable } from "../core/receivable";
import { Button, Table } from "react-bootstrap";
import { useDispatch } from "react-redux";
import actions from "../store/actions";

export interface ReceivableListItemProps {
  data: Receivable;
  onSelect: (id: string) => void;
}

export function ReceivableListItem1({
  data,
  onSelect,
}: ReceivableListItemProps): React.ReactElement {
  const dispatch = useDispatch();
  const onSync = () =>
    dispatch(actions.network2.importReceivable(data.id, "Sync"));

  return (
      <tr onClick={() => onSelect(data.id)}>
        <td className="text-left tx-normal" style={{height: "60px"}}>{data.issuer}</td>
        <td className="text-left tx-normal">{data.payer}</td>
        <td className="text-center tx-normal">{data.amount}</td>
      <td>
        <Button size={"sm"} onClick={onSync}>
          Transfer
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
      <td className="text-left tx-normal" style={{height: "60px"}}>{data.issuer}</td>
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
      <thead style={{ borderColor: "black!important"}}>
        <th className="text-left tx-normal">Issuer</th>
        <th className="text-left tx-normal">Payer</th>
        <th>Amount</th>
        <th>Action</th>
      </thead>
      <tbody>{children}</tbody>
    </Table>
  );
}
