package services

import (
	"blockchain/block"
	"blockchain/wallet"
	"fmt"
	"github.com/gin-gonic/gin"
	"gpp/chain"
	"net/http"
)

func SendMessage(ctx *gin.Context) {
	responseObj := new(Message)
	if err := ctx.BindJSON(&responseObj); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if responseObj.SendMessage == true {
		publicAddress := responseObj.PublicAddress
		toPublicAddress := responseObj.ToPublicAddress
		nodes, _ := sd["full_nodes"].(block.Data)
		ipAddress, _ := nodes[toPublicAddress].(string)
		message := responseObj.Message
		b := bc.MineBlock(
			block.Data{"head": "Message"},
			block.Data{"public_key": publicAddress, "message": message},
		)
		hash := b.Hash()
		signature, err := wallet.SignMessage(responseObj.PrivateKey, responseObj.PublicAddress, hash[:])
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		b.Header["signature"] = string(signature)
		//if err := cons.Exec(&bc, b); err != nil {
		//	ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		//	return
		//}
		if err := stts.Exec(&sd, b); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if err := chain.SaveStateData(&sd); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		dat, err := b.Save()
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if err := chain.SendBC(ipAddress, dat); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "message sent"})
		return
	}
	{
		fmt.Println(responseObj.Message)
	}
}

func RegisterMessagingRoutes(rg *gin.RouterGroup) {
	clientRoute := rg.Group("/msgservice")
	clientRoute.POST("/sendmsg", SendMessage)
}
