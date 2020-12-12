/*
 * Copyright (c) 2020. DLT Experts.
 *  Authors: Mikael Lazarev, Ivan Fedorov
 */

package controllers

import (
	"errors"
	"github.com/MikaelLazarev/hlf-blockchain/app_1/errorhandler"
	"github.com/gin-gonic/gin"
)

func withId(handler func(c *gin.Context, id string)) gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, ok := c.Params.Get("id")
		if !ok {
			errorhandler.ResponseWithAPIError(c, errorhandler.HttpBadRequestError(errors.New("Cant get ID")))
			return
		}

		handler(c, ID)
	}
}
