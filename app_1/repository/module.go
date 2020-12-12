/*
 * Copyright (c) 2020. DLT Experts.
 *  Authors: Mikael Lazarev, Ivan Fedorov
 */

package repository

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewHLFClient,
	NewReceivablesRepository,
)
