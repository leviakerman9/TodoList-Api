package ethereum

import (
	// "context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct{
  EthClient *ethclient.Client

}

func InitiateClient(rpcURL string) (*Client,error){
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum node: %v", err)
	}
	log.Println("Ethereum client connected successfully")

	return &Client{
		EthClient : client,
	},nil
}
