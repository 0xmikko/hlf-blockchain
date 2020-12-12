/*
 * Copyright (c) 2020. DLT Experts.
 *  Authors: Mikael Lazarev, Ivan Fedorov
 */

package config

import (
	"go.uber.org/fx"
)

var Module = fx.Option(
	fx.Provide(NewConfig))
