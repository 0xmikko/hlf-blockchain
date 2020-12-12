/*
 * Copyright (c) 2020. DLT Experts.
 *  Authors: Mikael Lazarev, Ivan Fedorov
 */

package core

import (
	"github.com/MikaelLazarev/hlf-blockchain/app_1/payloads"
)

type (
	Receivable struct {
		ID     string `json:"id"`
		Issuer string `json:"issuer"`
		Payer  string `json:"payer"`
		Amount int    `json:"amount"`
		Hash   string `json:"hash"`
	}

	ReceivableRepositoryI interface {
		Insert(item *Receivable) error
		FindOneByID(result *Receivable, id string) error
		List(result *[]Receivable) error
		GetFromAnotherChain(id string) error
	}

	ReceivableServiceI interface {
		Retrieve(id string) (*Receivable, error)
		List() ([]Receivable, error)
		Create(req *payloads.CreateReceivableReq) error
		Sync(id string) error
	}
)

func ReceivableFromReq(req *payloads.CreateReceivableReq) Receivable {
	var r Receivable
	r.Issuer = req.Issuer
	r.Payer = req.Payer
	r.Amount = req.Amount
	return r
}
