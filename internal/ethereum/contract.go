package ethereum

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/leviakerman9/TODOLIST-API/config"
	"github.com/leviakerman9/TODOLIST-API/contracts"
)

type Todolist struct{
   Todo *contracts.Todolistv2
   Auth *bind.TransactOpts
}


func InitiateContract(cfg *config.Config,cl *Client)(*Todolist,error) {
	if cl.EthClient == nil {
	log.Fatal("Ethereum client is nil.Ethereum node connection is not setup properly. First call InitClient()")
}
	log.Println("Contract address:", cfg.ContractAddress)

	address := common.HexToAddress(cfg.ContractAddress)
	instance, err := contracts.NewTodolistv2(address, cl.EthClient)
	if err != nil {
		log.Fatalf("Failed to load contract: %v", err)
	}
	// Todo = instance

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(cfg.ChainID)) // Sepolia chainId
	if err != nil {
		log.Fatalf("Failed to create auth: %v", err)
	}

	return &Todolist{
		Todo : instance,
		Auth: auth,
	},nil
}
