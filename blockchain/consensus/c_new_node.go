package consensus

import (
	"blockchain/block"
	"blockchain/blockchain"
	"blockchain/wallet"
	"errors"
	"net"
)

func validateNN(_ *blockchain.Blockchain, b block.Block) bool {
	header := b.Header
	if header["head"] != "NewNode" {
		return false
	}
	if _, ok := header["signature"].(string); !ok {
		return false
	}
	data := b.Data()
	if _, ok := data["public_key"].(string); !ok {
		return false
	}
	if _, ok := data["ip_address"].(string); !ok {
		return false
	}
	return true
}

func runNN(bc *blockchain.Blockchain, b block.Block) error {
	data := b.Data()
	header := b.Header
	publicKey, _ := data["public_key"].(string)
	ipAddress, _ := data["ip_address"].(string)
	signature, _ := header["signature"].(string)
	if len(publicKey) != 128 {
		return errors.New("invalid public key")
	}
	if net.ParseIP(ipAddress) == nil {
		return errors.New("invalid ip address")
	}
	hash := b.Hash()
	if !wallet.VerifySignature(publicKey, hash[:], []byte(signature)) {
		return errors.New("signature mismatch")
	}
	if err := bc.AddBlock(b); err != nil {
		return err
	}
	return nil
}

var CNewNode = Consensus{validateNN, runNN}
