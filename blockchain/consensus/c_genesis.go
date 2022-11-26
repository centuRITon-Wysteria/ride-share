package consensus

import "blockchain/block"
import "blockchain/blockchain"

func validateG(_ *blockchain.Blockchain, b block.Block) bool {
	header := b.Header
	if header["head"] != "Genesis" {
		return false
	}
	return true
}

func runG(bc *blockchain.Blockchain, b block.Block) error {
	if err := bc.AddBlock(b); err != nil {
		return err
	}
	return nil
}

var CGenesis = Consensus{validateG, runG}
