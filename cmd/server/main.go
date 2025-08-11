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

	client ,_:=ethereum.InitiateClient(cfg.RPCURL) //this function call will initaite and setup client and with connect to ethereum node 
	ethInstance,_:=ethereum.InitiateContract(cfg,client) //this function call will allows us to load contact and signer


	handler := &api.APIHandler{Eth: ethInstance}

	router := api.SetupRouter(handler) 
	router.Run(":8080")
}