package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a Letter of Credit
type SmartContract struct {
	contractapi.Contract
}

// LetterOfCredit describes basic details of a Letter of Credit
type LetterOfCredit struct {
	ID         string `json:"ID"`
	Importer   string `json:"importer"`
	Exporter   string `json:"exporter"`
	Amount     int    `json:"amount"`
	ExpiryDate string `json:"expiryDate"`
	Status     string `json:"status"`
}

// InitLedger adds a base set of lcs to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	lcs := []LetterOfCredit{
		{ID: "lc1", Importer: "Importer1", Exporter: "Exporter1", Amount: 100000, ExpiryDate: "2023-12-31", Status: "issued"},
		{ID: "lc2", Importer: "Importer2", Exporter: "Exporter2", Amount: 200000, ExpiryDate: "2024-01-31", Status: "issued"},
	}

	for _, lc := range lcs {
		lcJSON, err := json.Marshal(lc)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(lc.ID, lcJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// IssueLC issues a new letter of credit to the world state with given details.
func (s *SmartContract) IssueLC(ctx contractapi.TransactionContextInterface, id string, importer string, exporter string, amount int, expiryDate string) error {
	exists, err := s.LCExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the lc %s already exists", id)
	}

	lc := LetterOfCredit{
		ID:         id,
		Importer:   importer,
		Exporter:   exporter,
		Amount:     amount,
		ExpiryDate: expiryDate,
		Status:     "issued",
	}
	lcJSON, err := json.Marshal(lc)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, lcJSON)
}

// ReadLC returns the lc stored in the world state with given id.
func (s *SmartContract) ReadLC(ctx contractapi.TransactionContextInterface, id string) (*LetterOfCredit, error) {
	lcJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if lcJSON == nil {
		return nil, fmt.Errorf("the lc %s does not exist", id)
	}

	var lc LetterOfCredit
	err = json.Unmarshal(lcJSON, &lc)
	if err != nil {
		return nil, err
	}

	return &lc, nil
}

// VerifyDocuments updates the status of the lc with given id in world state.
func (s *SmartContract) VerifyDocuments(ctx contractapi.TransactionContextInterface, id string) error {
	lc, err := s.ReadLC(ctx, id)
	if err != nil {
		return err
	}

	lc.Status = "documents_verified"
	lcJSON, err := json.Marshal(lc)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, lcJSON)
}

// SettlePayment updates the status of the lc with given id in world state.
func (s *SmartContract) SettlePayment(ctx contractapi.TransactionContextInterface, id string) error {
	lc, err := s.ReadLC(ctx, id)
	if err != nil {
		return err
	}

	if lc.Status != "documents_verified" {
		return fmt.Errorf("documents not verified for lc %s", id)
	}

	lc.Status = "payment_settled"
	lcJSON, err := json.Marshal(lc)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, lcJSON)
}

// LCExists returns true when lc with given ID exists in world state
func (s *SmartContract) LCExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	lcJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return lcJSON != nil, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		fmt.Printf("Error creating trade finance chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting trade finance chaincode: %s", err.Error())
	}
}
