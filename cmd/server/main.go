package main

import (
	"fmt"

	// "github.com/ethereum/go-ethereum"
	"github.com/leviakerman9/TODOLIST-API/config"
	"github.com/leviakerman9/TODOLIST-API/internal/api"
	"github.com/leviakerman9/TODOLIST-API/internal/ethereum"
)

func main() {

	fmt.Println("this is the entry point")
	cfg := config.LoadConfig()

	ethereum.InitClient(cfg.RPCURL)
	ethereum.InitContract(cfg)

	router := api.SetupRouter()
	router.Run(":8080")
}