package ethereum

import (
	// "context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var EthClient *ethclient.Client

func InitClient(rpcURL string) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum node: %v", err)
	}
	EthClient = client
	log.Println("Ethereum client connected")
}
