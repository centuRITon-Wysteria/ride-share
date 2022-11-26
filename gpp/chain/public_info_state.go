package chain

import (
	"blockchain/block"
	"blockchain/states"
)

func initializePIS(sd *states.StateData) {
	(*sd)["public_info"] = block.Data{}
}

func validatePIS(_ *states.StateData, b block.Block) bool {
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

func runPIS(sd *states.StateData, b block.Block) error {
	data := b.Data()
	publicInfo, _ := (*sd)["public_info"].(block.Data)
	publicKey, _ := data["public_key"].(string)
	publicInfo[publicKey] = data
	return nil
}

var SPublicInfo = states.State{Initialize: initializePIS, Validate: validatePIS, Run: runPIS}
