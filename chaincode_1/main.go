package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"
	"strconv"
)

const salt = "Joiejofe"

// SmartContract provides functions for managing an Receivable
type SmartContract struct {
	contractapi.Contract
}

// Receivable describes basic details of what makes up a simple asset
type Receivable struct {
	ID     string `json:"id"`
	Issuer string `json:"issuer"`
	Payer  string `json:"payer"`
	Amount int    `json:"amount"`
	Hash   string `json:"hash"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Receivable{
		{ID: "rec1", Issuer: "GazpromNeft", Payer: "MarineLand", Amount: 300},
		{ID: "rec2", Issuer: "GazpromNeft", Payer: "MarineLines", Amount: 200},
		{ID: "rec3", Issuer: "GazpromNeft", Payer: "SovKomFlot", Amount: 1000},
	}

	for _, asset := range assets {
		asset.Hash = s.getHash(asset)
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateReceivable issues a new asset to the world state with given details.
func (s *SmartContract) CreateReceivable(ctx contractapi.TransactionContextInterface, id, issuer, payer string, amount int) error {
	exists, err := s.ReceivableExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	asset := Receivable{
		ID:     id,
		Issuer: issuer,
		Payer:  payer,
		Amount: amount,
	}

	asset.Hash = s.getHash(asset)

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) GetReceivable(ctx contractapi.TransactionContextInterface, id string) (*Receivable, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset Receivable
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// ReceivableExists returns true when asset with given ID exists in world state
func (s *SmartContract) ReceivableExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllReceivables(ctx contractapi.TransactionContextInterface) ([]*Receivable, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Receivable
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Receivable
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

func (s *SmartContract) getHash(r Receivable) string {
	hashStr := r.ID + r.Issuer + r.Payer + strconv.Itoa(r.Amount) + salt
	return hex.EncodeToString(crypto.Keccak256([]byte(hashStr)))
}

func main() {
	assetChaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating asset-transfer-basic chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
	}
}
