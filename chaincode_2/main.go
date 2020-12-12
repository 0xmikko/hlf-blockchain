package main

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const salt = "Joiejofe"
var network1Nodes = []string{"http://130.193.59.251:8080", "http://130.193.59.251:5000"}
const minConsensusShare = 50

// SmartContract provides functions for managing an Receivable
type SmartContract struct {
	contractapi.Contract
}

// Receivable describes basic details of what makes up a simple asset
type Receivable struct {
	ID     string `json:"id" binding:"required"`
	Issuer string `json:"issuer" binding:"required"`
	Payer  string `json:"payer" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
	Hash   string `json:"hash" binding:"required"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Receivable{
		{ID: "rec1", Issuer: "GazpromNeft", Payer: "MarineLand", Amount: 300},
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
func (s *SmartContract) ImportReceivable(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.ReceivableExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	client := &http.Client{}

	var errorsQty int
	var correctQty int

	var validResponse Receivable

	for _, node := range network1Nodes {

		// Making request
		url := node + "/api/receivables/" + id + "/"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}
		resp, err := client.Do(req)
		if resp.StatusCode != http.StatusOK {
			errorsQty++
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			errorsQty++
			continue
		}

		// Parsing result
		var rec Receivable
		if err := json.Unmarshal(body, &rec); err != nil {
			errorsQty++
			continue
		}

		// Check that
		if s.checkValidity(rec) == false {
			errorsQty++
			continue
		}

		if validResponse.Hash != "" && validResponse.Hash != rec.Hash{
			return errors.New("Got two different hashes from nodes")
		}

		validResponse = rec

		correctQty++
	}

	if correctQty < len(network1Nodes) * minConsensusShare / 100 {
		return errors.New("Consensus wasn't achieved")
	}

	assetJSON, err := json.Marshal(validResponse)
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

// Check is Receivable valid
func (s *SmartContract) checkValidity(r Receivable) bool {
	hashAlt := s.getHash(r)
	return hashAlt == r.Hash
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
