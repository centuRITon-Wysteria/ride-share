package services

import (
	"blockchain/block"
	"blockchain/wallet"
	"fmt"
	"github.com/gin-gonic/gin"
	"gpp/chain"
	"net/http"
)

func AnnounceTravel(ctx *gin.Context) {
	responseObj := new(AnnounceTravelPost)
	if err := ctx.BindJSON(responseObj); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	b := bc.MineBlock(
		block.Data{"head": "AnnounceTravel"},
		block.Data{
			"public_key": responseObj.PublicKey,
			"from_lat":   responseObj.FromLat,
			"from_lon":   responseObj.FromLon,
			"to_lat":     responseObj.ToLat,
			"to_lon":     responseObj.ToLon,
			"time":       responseObj.Time,
		},
	)
	hash := b.Hash()
	signature, err := wallet.SignMessage(responseObj.PrivateKey, responseObj.PublicKey, hash[:])
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	b.Header["signature"] = string(signature)
	if err := cons.Exec(&bc, b); err != nil {
		fmt.Println("hi", err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := stts.Exec(&sd, b); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := chain.SaveBlockchain(&bc); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := chain.SaveStateData(&sd); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "ok"})

	go func() {
		if err := chain.Sync(&bc, &sd); err != nil {
			fmt.Println(err)
		}
	}()
}
