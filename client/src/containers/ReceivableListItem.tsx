import React from "react";
import { Receivable } from "../core/receivable";
import { Button } from "react-bootstrap";
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
    <tr>
      <td>{data.issuer}</td>
      <td>{data.payer}</td>
      <td>{data.amount}</td>
      <td>
        <Button size={"sm"} onClick={onSync}>
          Sync
        </Button>{" "}
      </td>
    </tr>
  );
}

export function ReceivableListItem2({
  data,
  onSelect,
}: ReceivableListItemProps): React.ReactElement {
  return (
    <tr>
      <td>{data.issuer}</td>
      <td>{data.payer}</td>
      <td>{data.amount}</td>
      <td></td>
    </tr>
  );
}

export function ReceivableHeader({
  children,
}: React.PropsWithChildren<any>): React.ReactElement {
  return (
    <table>
      <thead>
        <th>Issuer</th>
        <th>Payer</th>
        <th>Amount</th>
        <th>Action</th>
      </thead>
      <tbody>{children}</tbody>
    </table>
  );
}
