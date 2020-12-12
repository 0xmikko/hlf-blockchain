/*
 * Copyright (c) 2020. DLT Experts.
 *  Authors: Mikael Lazarev, Ivan Fedorov
 */

package controllers

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Invoke(
		NewReceivablesController,
	),
	fx.Provide(NewServer))
