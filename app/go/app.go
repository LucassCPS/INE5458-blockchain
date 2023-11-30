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
	Manufacturer    string `json:"manufacturer"`
	Year      		string `json:"year"`
	Id				string `json:id`
}

func (m *AppChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (m *AppChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	switch fn {

	case "AddProduct":
		_, _ = strconv.Atoi(args[2])
		err := AddProduct(stub, args[0], args[1], args[2])
		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success([]byte(nil))

	case "GetAllProducts":
		result, err := GetAllProducts(stub)
		if err != nil {
			return shim.Error(err.Error())
		}

		resultBytes, err := json.Marshal(result)
		if err != nil {
			return shim.Error("Can't marshal result | "+err.Error())
		}

		return shim.Success(resultBytes)

	case "GetProduct":
		result, err := GetProduct(stub, args[0], args[1])
		if err != nil {
			return shim.Error(err.Error())
		}

		resultBytes, err := json.Marshal(result)
		if err != nil {
			return shim.Error("Can't marshal result | "+err.Error())
		}

		return shim.Success(resultBytes)

	default:
		return shim.Error("Function doesn't exist")
	}
}

func main() {
	if err := shim.Start(new(AppChaincode)); err != nil {
		log.Panicf("Erro ao tentar iniciar a chaincode da aplicação: %v", err)
	}
}

// Adiciona um produto
func AddProduct(stub shim.ChaincodeStubInterface, manufacturer string, year string, id string) error {
	product := &product{
		Manufacturer: manufacturer,
		Year: year,
		Id: id,
	}

	productBytes, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("Can't create product bytes | "+err.Error())
	}

	productID := manufacturer+"-"+id

	err = stub.PutState(productID, []byte(productBytes))
	if err != nil {
		return fmt.Errorf("Can't update world state | "+err.Error())
	}
	return nil
}

// Retorna todos os produtos
func GetAllProducts(stub shim.ChaincodeStubInterface) ([]product, error) {
	productIterator, err := stub.GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("Can't get world state | "+err.Error())
	}

	var products []product

	defer productIterator.Close()
	for productIterator.HasNext() {
		queryResponse, err := productIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("Can't get next product | "+err.Error())
		}
		var product product
		err = json.Unmarshal(queryResponse.Value, &product)
		if err != nil {
			return nil, fmt.Errorf("Can't get product info | "+err.Error())
		}
		products = append(products, product)
	}

	return products, nil
}

// Retorna um produto especifico
func GetProduct(stub shim.ChaincodeStubInterface, manufacturer string, id string) (*product, error) {
	productID := manufacturer+"-"+id
	productBytes, err := stub.GetState(productID)
	if err != nil {
		return nil, fmt.Errorf("Can't get world state | "+err.Error())
	}

	var product product
	err = json.Unmarshal(productBytes, &product)
	if err != nil {
		return nil, fmt.Errorf("Can't get product info | "+err.Error())
	}

	return &product, nil
}