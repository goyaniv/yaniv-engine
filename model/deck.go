package model

import (
	"encoding/json"
)

type Deck struct {
	Stack
}

// Return the pointer of a new allocated Deck
func DeckNew() *Deck {
	return &Deck{*StackNew()}
}

// Return the JSON representation of the deck
func (d *Deck) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Cards []int `json:"cards"`
		Value int   `json:"value"`
		Size  int   `json:"size"`
	}{
		Cards: d.ArrayID(),
		Value: d.Value(),
		Size:  d.Size(),
	})

}

// Return the total value of deck
func (d *Deck) Value() int {
	j := 0
	for _, card := range d.Cards {
		j = j + card.Value
	}
	return j
}

// Return the number of cards in the deck
func (d *Deck) Size() int {
	return len(d.Cards)
}

// Initialize the complete card game
func (d *Deck) InitReferenceDeck() {
	j := 1
	colours := [...]string{"spade", "heart", "diam", "club"}
	for _, colour := range colours {
		for i := 1; i < 14; i++ {
			(*d).Cards = append((*d).Cards, CardNew(j, i, colour))
			j++
		}
	}
}
