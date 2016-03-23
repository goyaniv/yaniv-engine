package main

import (
	"encoding/json"
	"errors"
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
	Yaniv           bool `json:"yaniv"`
	Asaf            bool `json:"asaf"`
	Playing         bool `json:"playing"`
	previousPlaying bool
	Ready           bool `json:"ready"`
	Loser           bool `json:"loser"`
	asafRank        int
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

// Discard cards
func (p *Player) Discard(discard []int) (*Stack, error) {
	// check if all the cards are present in player deck
	deckdiscard := StackNew()

	for _, cardid := range discard {
		if !p.HasCard(cardid) {
			return nil, errors.New("You can't discard these cards")
		}
		card := p.Hand.Remove(cardid)
		if card == nil {
			return nil, errors.New("You can't discard these cards")
		}
		deckdiscard.Add(card)
	}

	if !deckdiscard.IsValid() {
		p.Hand.AddStack(deckdiscard)
		return nil, errors.New("Invalid discarded deck")
	}
	return deckdiscard, nil
}

//HasCard check if the player own the card id
func (p *Player) HasCard(card int) bool {
	for _, c := range p.Hand.Cards {
		if card == c.ID {
			return true
		}
	}
	return false
}
