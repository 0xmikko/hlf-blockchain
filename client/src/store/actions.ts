/*
 * Copyright (c) 2020. Mikael Lazarev
 */

import * as network1 from './network1/actions';
import * as network2 from './network2/actions';
import * as operations from 'redux-data-connect/lib/operations/actions';

const actions = {
  network1,
  network2,
  operations,
};

export default actions;
