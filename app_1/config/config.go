/*
 * Copyright (c) 2020. DLT Experts.
 *  Authors: Mikael Lazarev, Ivan Fedorov
 */

package config

type Config struct {
	DiscoveryAsLocalhost string `env:"DISCOVERY_AS_LOCALHOST" validate:"required"`
	Env                  string `env:"ENV" default:"development" validate:"required"`
	Port                 string `env:"PORT" default:"8080" validate:"required"`
}
