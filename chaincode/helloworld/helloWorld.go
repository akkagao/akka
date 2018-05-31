package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type HelloWorld struct {
}

func (h *HelloWorld) Init(stub shim.ChaincodeStubInterface) peer.Response {

	return shim.Success(nil)
}

func (h *HelloWorld) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	if fn == "set" {
		return h.set(stub, args)
	} else if fn == "get" {
		return h.get(stub, args)
	}

	return shim.Success(nil)
}

func (h *HelloWorld) set(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (h *HelloWorld) get(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	result, err := stub.GetState(args[0])
	if err != nil {
		shim.Error(err.Error())
	}
	return shim.Success(result)
}

func main() {
	err := shim.Start(new(HelloWorld))
	if err != nil {
		fmt.Println(err)
	}
}
