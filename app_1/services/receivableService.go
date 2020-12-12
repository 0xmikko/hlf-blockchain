/*
 * Copyright (c) 2020. DLT Experts.
 *  Authors: Mikael Lazarev, Ivan Fedorov
 */

package services

import (
	"github.com/MikaelLazarev/hlf-blockchain/app_1/core"
	"github.com/MikaelLazarev/hlf-blockchain/app_1/payloads"
	uuid "github.com/satori/go.uuid"
)

type receivableService struct {
	repo core.ReceivableRepositoryI
}

func NewReceivableService(
	repo core.ReceivableRepositoryI,
) core.ReceivableServiceI {
	return &receivableService{
		repo: repo,
	}
}

func (r *receivableService) List() ([]core.Receivable, error) {
	var receivables []core.Receivable
	if err := r.repo.List(&receivables); err != nil {
		return nil, err
	}
	return receivables, nil
}

func (r *receivableService) Create(req *payloads.CreateReceivableReq) error {

	receivable := core.ReceivableFromReq(req)
	receivable.ID = uuid.NewV4().String()

	return r.repo.Insert(&receivable)
}
