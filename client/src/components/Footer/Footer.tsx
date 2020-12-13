import React from "react";
import {Container} from "react-bootstrap";

export function Footer() {
  return (
    <footer className={"footer-container"}>

      <div className={"footer-content"} >
        <Container style={{paddingTop: '10px', paddingBottom: "10px"}}>
                &copy; Copyright 2020, DLT Experts. All rights reserved
        </Container>
      </div>
    </footer>
  );
}
