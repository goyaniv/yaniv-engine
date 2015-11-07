package main

import (
	"encoding/json"
)

// Player struct
type Player struct {
	Name  string       `json:"name"`
	Score int          `json:"score"`
	Hand  *Deck        `json:"hand"`
	State *PlayerState `json:"state"`
}

// PlayerState struct
type PlayerState struct {
	Yaniv    bool `json:"yaniv"`
	Asaf     bool `json:"asaf"`
	Playing  bool `json:"playing"`
	Ready    bool `json:"ready"`
	Loser    bool `json:"loser"`
	AsafRank int  `json:"-"`
}

// PlayerNew initialize a new player with name
func PlayerNew(name string) *Player {
	return &Player{Name: name, Hand: DeckNew(), State: _PlayerStateNew()}
}

func _PlayerStateNew() *PlayerState {
	return &PlayerState{}
}

// MarshalJSON returns the JSON representation of a player
func (p *Player) MarshalJSON() ([]byte, error) {
	return json.Marshal(*p)
}
