/*
 * Copyright (c) 2020. DLT Experts.
 *  Authors: Mikael Lazarev, Ivan Fedorov
 */

package main

import (
	"context"
	"log"
	"time"

	"github.com/MikaelLazarev/hlf-blockchain/app_1/config"
	"github.com/MikaelLazarev/hlf-blockchain/app_1/controllers"
	"github.com/MikaelLazarev/hlf-blockchain/app_1/repository"
	"github.com/MikaelLazarev/hlf-blockchain/app_1/services"

	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		// Provide all the constructors we need, which teaches Fx how we'd like to
		// construct the *log.Logger, http.Handler, and *http.ServeMux types.
		// Remember that constructors are called lazily, so this block doesn't do
		// much on its own.
		config.Module,
		controllers.Module,
		repository.Module,
		services.Module,

		// Since constructors are called lazily, we need some invocations to
		// kick-start our application. In this case, we'll use Register. Since it
		// depends on an http.Handler and *http.ServeMux, calling it requires Fx
		// to build those types using the constructors above. Since we call
		// NewMux, we also register Lifecycle hooks to start and stop an HTTP
		// server.
		fx.Invoke(controllers.StartServer),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	<-app.Done()
}
