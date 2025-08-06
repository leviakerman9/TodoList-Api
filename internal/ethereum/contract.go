package ethereum

import (

	"log"
	"math/big"

	"github.com/leviakerman9/TODOLIST-API/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/leviakerman9/TODOLIST-API/contracts"
)

var Todo *contracts.Todolistv2
var Auth *bind.TransactOpts

func InitContract(cfg *config.Config) {
	if EthClient == nil {
	log.Fatal("Ethereum client is nil. Did you call InitClient() first?")
}
	log.Println("Contract address:", cfg.ContractAddress)

	address := common.HexToAddress(cfg.ContractAddress)
	instance, err := contracts.NewTodolistv2(address, EthClient)
	if err != nil {
		log.Fatalf("Failed to load contract: %v", err)
	}
	Todo = instance

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	Auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111)) // Sepolia chainId
	if err != nil {
		log.Fatalf("Failed to create auth: %v", err)
	}
}
