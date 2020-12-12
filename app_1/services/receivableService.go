/*
 * Copyright (c) 2020. DLT Experts.
 *  Authors: Mikael Lazarev, Ivan Fedorov
 */

package services

import (
	"encoding/hex"
	"github.com/MikaelLazarev/hlf-blockchain/app_1/core"
	"github.com/MikaelLazarev/hlf-blockchain/app_1/payloads"
	"github.com/ethereum/go-ethereum/crypto"
	uuid "github.com/satori/go.uuid"
	"strconv"
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
	hashStr := receivable.ID + receivable.Issuer + receivable.Payer + strconv.Itoa(receivable.Amount)
	receivable.Hash = hex.EncodeToString(crypto.Keccak256([]byte(hashStr)))

	return r.repo.Insert(&receivable)
}
