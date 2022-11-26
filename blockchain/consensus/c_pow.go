package consensus

import (
	"blockchain/block"
	"blockchain/blockchain"
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/big"
)

var difficulty = 2

type PoW struct {
	block  *block.Block
	target *big.Int
}

func RunPow(_ *blockchain.Blockchain, blk block.Block) error {
	powObj := newPow(&blk)
	powObj.run()
	return nil
}

func ValidatePow(_ *blockchain.Blockchain, blk block.Block) bool {
	powObj := newPow(&blk)
	return powObj.validate()
}
func newPow(b *block.Block) *PoW {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))
	pow := &PoW{b, target}
	return pow
}

func (p *PoW) initializeNonce(nonce int64) []byte {
	hash := p.block.Hash()
	h := hash[:]
	dat := bytes.Join(
		[][]byte{
			h,
			toBytes(nonce),
			toBytes(int64(difficulty)),
		},
		[]byte{},
	)
	return dat
}

func (p *PoW) run() {
	var intHash big.Int
	var nonce int64 = 0
	for true {
		dat := p.initializeNonce(nonce)
		Hash := sha256.Sum256(dat)
		intHash.SetBytes(Hash[:])
		if intHash.Cmp(p.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	p.block.Header["Nonce"] = nonce
}

func (p *PoW) validate() bool {
	var intHash big.Int
	nonce, _ := p.block.Header["Nonce"].(int64)
	dat := p.initializeNonce(nonce)
	hash := sha256.Sum256(dat)
	intHash.SetBytes(hash[:])
	return intHash.Cmp(p.target) == -1
}

func toBytes(number int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, number)
	if err != nil {
		fmt.Println(err)
	}
	return buff.Bytes()
}

var CPow = Consensus{ValidatePow, RunPow}
