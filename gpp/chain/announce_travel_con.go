package chain

import (
	"blockchain/block"
	"blockchain/consensus"
)
import "blockchain/blockchain"

func validateAT(_ *blockchain.Blockchain, b block.Block) bool {
	header := b.Header
	if header["head"] != "TravelAnnouncement" {
		return false
	}
	if _, ok := header["signature"].(string); !ok {
		return false
	}
	data := b.Data()
	if _, ok := data["public_key"].(string); !ok {
		return false
	}
	if _, ok := data["from_lat"].(string); !ok {
		return false
	}
	if _, ok := data["from_lon"].(string); !ok {
		return false
	}
	if _, ok := data["to_lat"].(string); !ok {
		return false
	}
	if _, ok := data["to_lon"].(string); !ok {
		return false
	}
	if _, ok := data["time"].(string); !ok {
		return false
	}
	return true

}

func runAT(bc *blockchain.Blockchain, b block.Block) error {
	return nil
}

var CAnnounceTravel = consensus.Consensus{Validate: validateAT, Run: runAT}
