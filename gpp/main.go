package main

import (
	"blockchain/block"
	"blockchain/blockchain"
	"fmt"
	"github.com/gin-gonic/gin"
	"gpp/chain"
	"gpp/services"
	"log"
)

func Init() {
	sd := chain.LoadStateData()
	stts := chain.LoadStates()
	stts.Init(&sd)
	bc := blockchain.New(block.Data{})
	if err := chain.SaveStateData(&sd); err != nil {
		fmt.Println(err)
	}
	if err := chain.SaveBlockchain(&bc); err != nil {
		fmt.Println(err)
	}
}

func main() {
	if false {
		Init()
	} else {
		chainName := "baby_chain"
		server := gin.Default()
		basePath := server.Group("/" + chainName)
		services.RegisterClientRoutes(basePath)
		services.RegisterMessagingRoutes(basePath)
		log.Fatalln(server.Run(":9090"))
	}
}
