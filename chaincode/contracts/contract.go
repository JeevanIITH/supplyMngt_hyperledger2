package contracts

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically
type Asset struct {
	// AppraisedValue int    `json:"AppraisedValue"`
	// Color          string `json:"Color"`
	ProductID             string `json:"ProductID"`
	Name          string `json:"Name"`
	
	// Size           int    `json:"Size"`
	Description	    string `json:"Description"`
	ManufacturingDate string `json:"ManufacturingDate"`
	BatchNumber     string `json:"BatchNumber"`

	SupplyDate 		string `json:"SupplyDate"`
	WarehouseLocation  	string `json:"WarehouseLocation"`

	WholesaleDate 	string `json:"WholesaleDate"`
	WholesaleLocation 	string `json:"WholesaleLocation"`
	WholesaleQuantity 	string `json:"WholesaleQuantity"`

	

	Status          string  `json:"Status"`  // sold , NA
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Asset{

		{ProductID: "asset1", Name :"iPhone 16", Description :"iPhone 16", ManufacturingDate : "2021-01-01", BatchNumber : "B1", SupplyDate : "2021-01-02", WarehouseLocation : "W1", WholesaleDate : "2021-01-03", WholesaleLocation : "WL1", WholesaleQuantity : "100", Status : "NA" },
		{ProductID: "asset2", Name :"iPhone 15 Pro", Description: "iPhone 15 Pro", ManufacturingDate : "2021-01-01", BatchNumber : "B1", SupplyDate : "2021-01-02", WarehouseLocation : "W1", WholesaleDate : "2021-01-03", WholesaleLocation : "WL1", WholesaleQuantity : "100", Status : "sold"},


		// {ProductID: "asset1", Description :"iPhone 16", Status : "Ordered" },
		// {ProductID: "asset2", Description: "iPhone 15 Pro",   Status : "Shipped"},
		// {ProductID: "asset3", Description: "iPhone 15", Status : "Delivered"},

	}

	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ProductID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateProduct(ctx contractapi.TransactionContextInterface, productid string , name string , description string , manufacturingDate string, batchNumber string) error {
	exists, err := s.AssetExists(ctx, productid)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", productid)
	}

	asset := Asset{
		ProductID:             productid,
		Name:			name,
		Description:		description,
		ManufacturingDate:	manufacturingDate,
		BatchNumber:		batchNumber,
		
		SupplyDate:         "NA",
		WarehouseLocation:   "NA",
		WholesaleDate:  "NA",
		WholesaleLocation: "NA",
		WholesaleQuantity:  "NA",

		Status:  "created",

	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(productid, assetJSON)
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id string, status string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	// overwriting original asset with new asset
	asset := Asset{
		ProductID:             id,
		Status:			status,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

func (s *SmartContract)  SupplyProduct(ctx contractapi.TransactionContextInterface, id string, supplyDate string, warehouseLocation string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}
	asset.SupplyDate = supplyDate
	asset.WarehouseLocation = warehouseLocation
	asset.Status = "Supplied"
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, assetJSON)
	
}

func (s *SmartContract)  WholesaleProduct(ctx contractapi.TransactionContextInterface, id string, wholesaleDate string , wholeSaleLocation string , wholesaleQuantity string) error {  
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}
	asset.WholesaleDate = wholesaleDate
	asset.WholesaleLocation = wholeSaleLocation
	asset.WholesaleQuantity = wholesaleQuantity
	asset.Status = "Wholesaled"
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, assetJSON)
}

func (s *SmartContract) QueryProduct(ctx contractapi.TransactionContextInterface, id string ) (*Asset, error) {
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return nil,err
	}

	return asset , nil
	
} 

func (s *SmartContract) UpdateProductStatus(ctx contractapi.TransactionContextInterface, id string , status string ) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	asset.Status = status
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, assetJSON)
}

// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// AssetExists returns true when asset with given ProductID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// // TransferAsset updates the owner field of asset with given id in world state, and returns the old owner.
// func (s *SmartContract) TransferAsset(ctx contractapi.TransactionContextInterface, id string, newOwner string) (string, error) {
// 	asset, err := s.ReadAsset(ctx, id)
// 	if err != nil {
// 		return "", err
// 	}

// 	oldOwner := asset.Owner
// 	asset.Owner = newOwner

// 	assetJSON, err := json.Marshal(asset)
// 	if err != nil {
// 		return "", err
// 	}

// 	err = ctx.GetStub().PutState(id, assetJSON)
// 	if err != nil {
// 		return "", err
// 	}

// 	return oldOwner, nil
// }

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}
