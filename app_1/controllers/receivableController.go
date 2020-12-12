/*
 * Copyright (c) 2020. DLT Experts.
 *  Authors: Mikael Lazarev, Ivan Fedorov
 */

package controllers

import (
	"errors"
	"github.com/MikaelLazarev/hlf-blockchain/app_1/config"
	"github.com/MikaelLazarev/hlf-blockchain/app_1/core"
	"github.com/MikaelLazarev/hlf-blockchain/app_1/errorhandler"
	"github.com/MikaelLazarev/hlf-blockchain/app_1/payloads"
	"github.com/gin-gonic/gin"
	"net/http"
)

type receivableController struct {
	service core.ReceivableServiceI
}

// receivableController: manage receivables
func NewReceivablesController(config *config.Config, g *gin.Engine,
	rs core.ReceivableServiceI) {

	controller := receivableController{
		service: rs,
	}

	r := g.Group("/api/receivables/")

	r.GET("/", controller.List)
	r.POST("/", controller.Create)
}

// GET: /api/ERC20s/
// Returns array of active ERC20s
func (bc *receivableController) List(c *gin.Context) {

	result, err := bc.service.List()
	if err != nil {
		errorhandler.ResponseWithAPIError(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

//// POST: /api/ERC20s/
func (bc *receivableController) Create(c *gin.Context) {
	var req payloads.CreateReceivableReq
	if err := c.BindJSON(&req); err != nil {
		errorhandler.ResponseWithAPIError(c, errorhandler.HttpBadRequestError(errors.New("Incorrect request")))
		return
	}

	if err := bc.service.Create(&req); err != nil {
		errorhandler.ResponseWithAPIError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
