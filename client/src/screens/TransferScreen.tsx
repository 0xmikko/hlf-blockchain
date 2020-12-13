import React from "react";
import { Container, Row, Col } from "react-bootstrap";
import { Network1Form } from "./Network1Form";
import { Network2List } from "./Network2List";
import { AppBar } from "../components/AppBar/AppBar";
import { Footer } from "../components/Footer/Footer";
import { Helmet } from "react-helmet";

export function TransferScreen(): React.ReactElement {
  return (
    <>
      <Helmet title={"HLF Blockchain Transfer"} />
      <Container fluid style={{ padding: 0, margin: 0 }}>
        <AppBar />
        <Container style={{ width: "100%" }}>
          <Row>
            <Col xl={6} lg={6} sm={12} xs={12}>
              <Network1Form />
            </Col>
            <Col xl={6} lg={6} sm={12} xs={12}>
              <Network2List />
            </Col>
          </Row>
        </Container>
        <Footer />
      </Container>
    </>
  );
}
