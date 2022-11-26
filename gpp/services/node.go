package services

import (
	"blockchain/block"
	"blockchain/wallet"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"gpp/chain"
	"io"
	"net/http"
)

func NewNode(ctx *gin.Context) {
	responseObj := new(NewNodePost)
	if err := ctx.BindJSON(responseObj); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ipAddress := responseObj.IpAddress

	returnObj := new(NewNodeResponse)
	publicKey, privateKey, err := wallet.GeneratePublicAddressAndKey()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	returnObj.PublicKey = publicKey
	returnObj.PrivateKey = privateKey

	b := bc.MineBlock(
		block.Data{"head": "NewNode"},
		block.Data{"public_key": publicKey, "ip_address": ipAddress},
	)
	hash := b.Hash()
	signature, err := wallet.SignMessage(privateKey, publicKey, hash[:])
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	b.Header["signature"] = string(signature)
	if err := cons.Exec(&bc, b); err != nil {
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
	ctx.IndentedJSON(http.StatusOK, returnObj)

	go func() {
		if err := chain.Sync(&bc, &sd); err != nil {
			fmt.Println(err)
		}
	}()
}

func PublicInfo(ctx *gin.Context) {
	responseObj := new(PublicInfoPost)
	if err := ctx.BindJSON(responseObj); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	postBody, _ := json.Marshal(map[string]string{
		"ip_address": responseObj.IpAddress,
	})
	responseBody := bytes.NewBuffer(postBody)
	url := "http://localhost:9090" + "/baby_chain/service/newnode"
	resp, err := http.Post(url, "application/json", responseBody)
	fmt.Println("hi")
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if resp.StatusCode != http.StatusOK {
		ctx.IndentedJSON(resp.StatusCode, gin.H{"message": resp})
		return
	}
	var NNRes NewNodeResponse
	if err := json.Unmarshal(body, &NNRes); err != nil {
		ctx.IndentedJSON(resp.StatusCode, gin.H{"message": resp})
		return
	}
	b := bc.MineBlock(
		block.Data{"head": "PublicInfo"},
		block.Data{
			"public_key": NNRes.PublicKey,
			"name":       responseObj.Name,
			"driver":     responseObj.Driver,
			"license":    responseObj.License,
			"email":      responseObj.Email,
			"occupation": responseObj.Occupation,
			"hobbies":    responseObj.Hobbies,
			"skills":     responseObj.Skills,
			"interests":  responseObj.Interests,
			"others":     responseObj.Others,
		},
	)
	hash := b.Hash()
	signature, err := wallet.SignMessage(NNRes.PrivateKey, NNRes.PublicKey, hash[:])
	b.Header["signature"] = string(signature)
	if err := cons.Exec(&bc, b); err != nil {
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

	ctx.IndentedJSON(http.StatusOK, NNRes)

	go func() {
		if err := chain.Sync(&bc, &sd); err != nil {
			fmt.Println(err)
		}
	}()
}
