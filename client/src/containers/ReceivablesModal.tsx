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

  if (data === null) return <div />
  return (
    <Modal show={show} onHide={onHide}>
      <Modal.Header closeButton>
        <Modal.Title>Modal heading</Modal.Title>
      </Modal.Header>
      <Modal.Body>Woohoo, you're reading this text in a modal!</Modal.Body>
    </Modal>
  );
}
