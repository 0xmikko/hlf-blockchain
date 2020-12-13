import React from "react";
import { Modal } from "react-bootstrap";
import { Receivable } from "../core/receivable";

export interface ReceivablesModalProps {
  data: Receivable | null;
  show: boolean;
  onHide: () => void;
}

export function ReceivablesModal({
  data,
  show,
  onHide,
}: ReceivablesModalProps): React.ReactElement {
  if (data === null) return <div />;
  return (
    <Modal
      show={show}
      onHide={onHide}
      size={"lg"}
      centered
    >
      <Modal.Header closeButton>
        <Modal.Title>Receivable: {data.id}</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        ID: {data.id} <br />
        Issuer: {data.issuer} <br/>
        Payer: {data.payer} <br/>
        Amount: {data.amount} <br/>
        Hash: {data.hash} <br/>
      </Modal.Body>
    </Modal>
  );
}
