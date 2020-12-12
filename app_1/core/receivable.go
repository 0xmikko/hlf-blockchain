/*
 * Copyright (c) 2020. DLT Experts.
 *  Authors: Mikael Lazarev, Ivan Fedorov
 */

package core

import (
	"github.com/MikaelLazarev/hlf-blockchain/app_1/payloads"
	"time"
)

type (
	Receivable struct {
		ID     string    `json:"id"`
		Date   time.Time `json:"date"`
		Issuer int       `json:"issuer"`
		Payer  int       `json:"payer"`
		Amount int       `json:"amount"`
	}

	ReceivableRepositoryI interface {
		Insert(item *Receivable) error
		List(result *[]Receivable) error
	}

	ReceivableServiceI interface {
		List() ([]Receivable, error)
		Create(req *payloads.CreateReceivableReq) error
	}
)

func ReceivableFromReq(req *payloads.CreateReceivableReq) Receivable {
	var r Receivable
	r.Date = req.Date
	r.Issuer = req.Issuer
	r.Payer = req.Payer
	r.Amount = req.Amount
	return r
}
