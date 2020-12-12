/*
 * Copyright (c) 2020. DLT Experts.
 *  Authors: Mikael Lazarev, Ivan Fedorov
 */

package services

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewReceivableService,
)
