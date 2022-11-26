package chain

import (
	"blockchain/block"
	"blockchain/consensus"
)
import "blockchain/blockchain"

func validatePI(_ *blockchain.Blockchain, b block.Block) bool {
	header := b.Header
	if header["head"] != "PublicInfo" {
		return false
	}
	if _, ok := header["signature"].(string); !ok {
		return false
	}
	data := b.Data()
	if _, ok := data["public_key"].(string); !ok {
		return false
	}
	if _, ok := data["name"].(string); !ok {
		return false
	}
	if _, ok := data["driver"].(bool); !ok {
		return false
	}
	if _, ok := data["license"].(string); !ok {
		return false
	}
	if _, ok := data["email"].(string); !ok {
		return false
	}
	if _, ok := data["occupation"].(string); !ok {
		return false
	}
	if _, ok := data["hobbies"].(string); !ok {
		return false
	}
	if _, ok := data["skills"].(string); !ok {
		return false
	}
	if _, ok := data["interests"].(string); !ok {
		return false
	}
	if _, ok := data["others"].(string); !ok {
		return false
	}
	return true

}

func runPI(bc *blockchain.Blockchain, b block.Block) error {
	if err := bc.AddBlock(b); err != nil {
		return err
	}
	return nil
}

var CPublicInfo = consensus.Consensus{Validate: validatePI, Run: runPI}
