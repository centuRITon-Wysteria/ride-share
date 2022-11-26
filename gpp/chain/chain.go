package chain

import (
	"blockchain/blockchain"
	"blockchain/consensus"
	"blockchain/db/jsoner"
	"blockchain/states"
)

import (
	"fmt"
	"os"
)

var bcFile = "../blockchain/db/blockchain.bin"
var sdFile = "../blockchain/db/statedata.bin"

func LoadBlockchain() blockchain.Blockchain {
	bchData, err := jsoner.ReadData(bcFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	bc, err := blockchain.Load(bchData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return bc
}

func SaveBlockchain(bc *blockchain.Blockchain) error {
	save, err := bc.Save()
	if err != nil {
		return err
	}
	if err := jsoner.WriteData(bcFile, save); err != nil {
		return err
	}
	return nil
}

func LoadStateData() states.StateData {
	sdData, err := jsoner.ReadData(sdFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sd, err := states.Load(sdData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return sd
}

func SaveStateData(sd *states.StateData) error {
	save, err := sd.Save()
	if err != nil {
		return err
	}
	if err := jsoner.WriteData(sdFile, save); err != nil {
		return err
	}
	return nil
}

func LoadConsensus() consensus.CAlgo {
	return consensus.New(CPublicInfo, CAnnounceTravel)
}

func LoadStates() states.States {
	return states.New(SPublicInfo, SAts)
}
