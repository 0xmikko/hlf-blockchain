/*
 * Copyright (c) 2020. Mikael Lazarev
 */

import React from 'react';
import {Container, Nav, Navbar} from 'react-bootstrap';

export function AppBar() : React.ReactElement {

  return (
    <Navbar
        bg="dark"
        expand="md"
        style={{ backgroundColor: `EEE!important`, height: '65px' }}
      >
      <Container>
        <Navbar.Brand>
            <span
              style={{color: 'white', fontSize: '18px', fontWeight: 'bold'}}>
              HLF Blockchain Transfer
            </span>
        </Navbar.Brand>

        <Navbar
          style={{
            width: '100%',
            display: 'flex',
            flexDirection: 'row',
            justifyContent: 'space-between',
          }}>
          <Nav
            className="navbar-menu"
            style={{justifyContent: 'center', width: '100%'}}></Nav>
        </Navbar>
        {/*<Navbar.Toggle aria-controls="basic-navbar-nav" />*/}
      </Container>
    </Navbar>
  );
};
