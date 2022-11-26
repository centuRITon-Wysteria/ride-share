package chain

import (
	"blockchain/block"
	"blockchain/states"
)

func initializeATS(sd *states.StateData) {
	(*sd)["open_announcements"] = block.Data{}
}

func validateATS(_ *states.StateData, b block.Block) bool {
	header := b.Header
	if header["head"] != "AnnounceTravel" {
		return false
	}
	return true
}

func runATS(sd *states.StateData, b block.Block) error {
	data := b.Data()
	oa, _ := (*sd)["open_announcements"].(block.Data)
	oa[block.HashB([]byte(data.String())).Hex()] = data
	return nil
}

var SAts = states.State{Initialize: initializeATS, Validate: validateATS, Run: runATS}
