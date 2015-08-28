package model

import (
	"encoding/json"
)

type Player struct {
	Name  string       `json:"name"`
	Score int          `json:"score"`
	Hand  *Deck        `json:"hand"`
	State *PlayerState `json:"state"`
}

type PlayerState struct {
	Yaniv    bool `json:"yaniv"`
	Asaf     bool `json:"asaf"`
	Playing  bool `json:"playing"`
	Ready    bool `json:"ready"`
	Loser    bool `json:"loser"`
	AsafRank int  `json:"-"`
}

func PlayerNew(name string) *Player {
	return &Player{Name: name, Hand: DeckNew(), State: _PlayerStateNew()}
}

func _PlayerStateNew() *PlayerState {
	return &PlayerState{}
}

func (p *Player) MarshalJSON() ([]byte, error) {
	return json.Marshal(*p)
}
