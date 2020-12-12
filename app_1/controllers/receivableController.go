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
	r.GET("/:id/", withId(controller.Retrieve))
	r.POST("/sync/", controller.Sync)
}

// GET: /api/receivables/:id/
func (bc *receivableController) Retrieve(c *gin.Context, id string) {
	result, err := bc.service.Retrieve(id)
	if err != nil {
		errorhandler.ResponseWithAPIError(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

// GET: /api/receivables/
// Returns array of receivables
func (bc *receivableController) List(c *gin.Context) {

	result, err := bc.service.List()
	if err != nil {
		errorhandler.ResponseWithAPIError(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

// POST: /api/receivables/
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

// POST: /api/receivables/sync/
func (bc *receivableController) Sync(c *gin.Context) {
	var req payloads.SyncReq
	if err := c.BindJSON(&req); err != nil {
		errorhandler.ResponseWithAPIError(c, errorhandler.HttpBadRequestError(errors.New("Incorrect request")))
		return
	}

	if err := bc.service.Sync(req.ID); err != nil {
		errorhandler.ResponseWithAPIError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
