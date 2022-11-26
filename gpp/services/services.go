package services

import (
	"gpp/chain"
)

import (
	"github.com/gin-gonic/gin"
)

var bc = chain.LoadBlockchain()
var sd = chain.LoadStateData()
var cons = chain.LoadConsensus()
var stts = chain.LoadStates()

func RegisterClientRoutes(rg *gin.RouterGroup) {
	clientRoute := rg.Group("/service")

	clientRoute.POST("/newnode", NewNode)
	clientRoute.POST("/sync", Sync)
	clientRoute.POST("/publicinfo", PublicInfo)
	clientRoute.POST("/announcetravel", AnnounceTravel)
	clientRoute.POST("/messaging", MessageService)
}
