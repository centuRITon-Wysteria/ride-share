package states

import (
	"blockchain/block"
	"encoding/json"
)

type StateData block.Data

func (sd *StateData) Save() ([]byte, error) {
	return json.Marshal(sd)
}

func Load(save []byte) (StateData, error) {
	var sd StateData
	if err := json.Unmarshal(save, &sd); err != nil {
		return StateData{}, err
	}
	d := block.Data(sd)
	if err := d.Validate(); err != nil {
		return nil, err
	}
	return StateData(d), nil
}

type State struct {
	Initialize func(*StateData)
	Validate   func(*StateData, block.Block) bool
	Run        func(*StateData, block.Block) error
}

type States []State

func (st *States) Init(sd *StateData) {
	for _, state := range *st {
		state.Initialize(sd)
	}
}

func (st *States) Exec(sd *StateData, b block.Block) error {
	if err := b.Validate(); err != nil {
		return err
	}
	for _, state := range *st {
		if state.Validate(sd, b) {
			if err := state.Run(sd, b); err != nil {
				return err
			}
		}
	}
	return nil
}

func New(sts ...State) States {
	return append(States{SNode, SGenesis}, sts...)
}
