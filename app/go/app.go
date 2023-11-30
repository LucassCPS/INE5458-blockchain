package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	peer "github.com/hyperledger/fabric-protos-go/peer"
)

// Estrutura da chaincode
type AppChaincode struct {
}

// Estrutura de um produto
type product struct {
	Manufacturer string `json:"manufacturer"`
	ModelName    string `json:"modelName"`
	ModelId      int    `json:"modelId"`
	ExtraInfo    string `json:"extraInfo"`
}

func (m *AppChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func main() {
	if err := shim.Start(new(AppChaincode)); err != nil {
		log.Panicf("Erro ao tentar iniciar a chaincode da aplicação: %v", err)
	}
}

func (m *AppChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	switch fn {

	case "AddProduct":
		if len(args) != 4 {
			return shim.Error("Incorrect number of arguments. Expecting 4")
		}

		modelId, err := strconv.Atoi(args[2])
		if err != nil {
			return shim.Error(err.Error())
		}

		err = AddProduct(stub, args[0], args[1], modelId, args[3])
		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success([]byte(nil))

	case "GetProduct":
		if len(args) != 2 {
			return shim.Error("Incorrect number of arguments. Expecting 2")
		}

		modelId, err := strconv.Atoi(args[1])
		if err != nil {
			return shim.Error(err.Error())
		}

		result, err := GetProduct(stub, args[0], modelId)
		if err != nil {
			return shim.Error(err.Error())
		}

		resultBytes, err := json.Marshal(result)
		if err != nil {
			return shim.Error("Can't marshal result | " + err.Error())
		}

		return shim.Success(resultBytes)

	case "GetAllProducts":
		result, err := GetAllProducts(stub)
		if err != nil {
			return shim.Error(err.Error())
		}

		resultBytes, err := json.Marshal(result)
		if err != nil {
			return shim.Error("Can't marshal result | " + err.Error())
		}

		return shim.Success(resultBytes)

	case "GetProductsByManufacturer":
		if len(args) != 1 {
			return shim.Error("Incorrect number of arguments. Expecting 1")
		}

		result, err := GetProductsByManufacturer(stub, args[0])
		if err != nil {
			return shim.Error(err.Error())
		}

		resultBytes, err := json.Marshal(result)
		if err != nil {
			return shim.Error("Can't marshal result | " + err.Error())
		}

		return shim.Success(resultBytes)

	default:
		return shim.Error("Function doesn't exist")
	}
}

// Adiciona um produto
func AddProduct(stub shim.ChaincodeStubInterface, manufacturer string, modelName string, modelId int, extraInfo string) error {

	product := &product{
		Manufacturer: manufacturer,
		ModelName:    modelName,
		ModelId:      modelId,
		ExtraInfo:    extraInfo,
	}

	productBytes, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("Can't create product bytes | " + err.Error())
	}

	productId := modelName + "-" + strconv.Itoa(modelId)

	err = stub.PutState(productId, []byte(productBytes))
	if err != nil {
		return fmt.Errorf("Can't update world state | " + err.Error())
	}
	return nil
}

// Retorna um produto especifico
func GetProduct(stub shim.ChaincodeStubInterface, modelName string, modelId int) (*product, error) {
	productId := modelName + "-" + strconv.Itoa(modelId)

	productBytes, err := stub.GetState(productId)
	if err != nil {
		return nil, fmt.Errorf("Can't get world state | " + err.Error())
	}

	var product product
	err = json.Unmarshal(productBytes, &product)
	if err != nil {
		return nil, fmt.Errorf("Can't get product info | " + err.Error())
	}

	return &product, nil
}

// Retorna todos os produtos de um fabricante específico
func GetProductsByManufacturer(stub shim.ChaincodeStubInterface, manufacturer string) ([]product, error) {
	productIterator, err := stub.GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("Can't get world state | " + err.Error())
	}

	var products []product

	defer productIterator.Close()
	for productIterator.HasNext() {
		queryResponse, err := productIterator.Next()

		if err != nil {
			return nil, fmt.Errorf("Can't get next product | " + err.Error())
		}

		var currentProduct product

		err = json.Unmarshal(queryResponse.Value, &currentProduct)
		if err != nil {
			return nil, fmt.Errorf("Can't get product info | " + err.Error())
		}

		// Verifica se o fabricante do produto é o desejado
		if currentProduct.Manufacturer == manufacturer {
			products = append(products, currentProduct)
		}
	}

	return products, nil
}

// Retorna todos os produtos
func GetAllProducts(stub shim.ChaincodeStubInterface) ([]product, error) {
	productIterator, err := stub.GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("Can't get world state | " + err.Error())
	}

	var products []product

	defer productIterator.Close()
	for productIterator.HasNext() {
		queryResponse, err := productIterator.Next()

		if err != nil {
			return nil, fmt.Errorf("Can't get next product | " + err.Error())
		}

		var product product

		err = json.Unmarshal(queryResponse.Value, &product)
		if err != nil {
			return nil, fmt.Errorf("Can't get product info | " + err.Error())
		}
		products = append(products, product)
	}

	return products, nil
}
