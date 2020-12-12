/*
 * Copyright (c) 2020. DLT Experts.
 *  Authors: Mikael Lazarev, Ivan Fedorov
 */

package repository

import (
	"encoding/json"
	"github.com/MikaelLazarev/hlf-blockchain/app_1/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"log"
	"strconv"
)

type receivablesRepository struct {
	hlf      *gateway.Network
	contract *gateway.Contract
}

func NewReceivablesRepository(hlf *gateway.Network) core.ReceivableRepositoryI {

	contract := hlf.GetContract("basic")

	log.Println("--> Submit Transaction: InitLedger, function creates the initial set of assets on the ledger")
	result, err := contract.SubmitTransaction("InitLedger")
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	}
	log.Println(string(result))

	return &receivablesRepository{
		hlf: hlf,
	}
}

// Insert New Receivable to HLF
func (r *receivablesRepository) Insert(item *core.Receivable) error {
	_, err := r.contract.SubmitTransaction("CreateReceivable",
		item.ID,
		item.Issuer,
		item.Payer,
		strconv.Itoa(item.Amount),
		item.Hash,
	)
	return err
}

// List All Receivables from HLF
func (r *receivablesRepository) List(result *[]core.Receivable) error {

	recBytes, err := r.contract.EvaluateTransaction("GetAllReceivables")
	if err != nil {
		return err
	}
	if err := json.Unmarshal(recBytes, &result); err != nil {
		return err
	}
	return nil

}
