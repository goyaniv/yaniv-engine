package model

import (
	"encoding/json"
	"math/rand"
	"sort"
	"time"
)

type Stack struct {
	Cards []*Card `json:"cards"`
}

func (s *Stack) MarshalJSON() ([]byte, error) {
	// Return the JSON representation of the deck
	return json.Marshal(struct {
		Cards []int `json:"cards"`
	}{
		Cards: s.ArrayID(),
	})
}

func StackNew() *Stack {
	// Return the pointer of a new allocated Deck
	return &Stack{
		Cards: make([]*Card, 0),
	}
}

func (s *Stack) ArrayID() []int {
	// Return array of cards id
	arrayid := make([]int, 0)
	for _, card := range s.Cards {
		arrayid = append(arrayid, card.Id)
	}
	return arrayid
}

func (s *Stack) TopCardId() int {
	// Return the id of the card on top of the deck
	return s.Cards[0].Id
}

func (s *Stack) Add(c *Card) {
	// Add card to deck
	(*s).Cards = append((*s).Cards, c)
}

func (s *Stack) Remove(id int) *Card {
	// Remove card from deck and return the deleted card
	for i, card := range s.Cards {
		if card.Id == id {
			(*s).Cards = append((*s).Cards[:i], (*s).Cards[i+1:]...)
			return card
		}
	}
	return nil
}

func (s *Stack) AddStack(stack *Stack) {
	// Merge two deck together
	for i, _ := range stack.Cards {
		s.Add(stack.Remove(i))
	}
}

func (s *Stack) Shuffle() {
	// Shuffle the complete deck
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range s.Cards {
		j := rand.Intn(i + 1)
		(*s).Cards[i], (*s).Cards[j] = (*s).Cards[j], (*s).Cards[i]
	}
}

func (s *Stack) IsValid() bool {
	// Return true if the deck can be played in game
	return s.IsSequence() || s.IsMultiple() || len(s.Cards) == 1
}

func (s *Stack) IsMultiple() bool {
	// Return true if all cards have the same value and if the deck size
	// > 1
	if len(s.Cards) < 2 {
		return false
	}
	value := s.Cards[0].Value
	for _, card := range s.Cards {
		if card.Value != value {
			return false
		}
	}
	return true
}

func (s *Stack) IsSequence() bool {
	// Return true if the deck representation is a sequence
	// If less than 3 cards, not a seq.
	if s.Len() < 3 {
		return false
	}
	// check if all cards have same Symbol
	for _, card := range s.Cards {
		if card.Colour != s.Cards[0].Colour {
			return false
		}
	}
	// Sort the deck
	sort.Sort(s)
	for i, _ := range s.Cards {
		if s.Cards[0].Value != s.Cards[i].Value-i {
			return false
		}
	}
	return true
}

func (s Stack) Swap(i, j int) {
	s.Cards[i], s.Cards[j] = s.Cards[j], s.Cards[i]
}

func (s Stack) Less(i, j int) bool {
	return s.Cards[i].Value < s.Cards[j].Value
}

func (s Stack) Len() int {
	return len(s.Cards)
}
