package main

import (
	"encoding/json"
	"math/rand"
	"sort"
	"time"
)

// Stack struct is a stack of cards
type Stack struct {
	Cards []*Card `json:"cards"`
}

// MarshalJSON render the json representation of a stack
func (s *Stack) MarshalJSON() ([]byte, error) {
	// Return the JSON representation of the deck
	return json.Marshal(struct {
		Cards []int `json:"cards"`
	}{
		Cards: s.ArrayID(),
	})
}

// StackNew Initialize a Stack object
func StackNew() *Stack {
	// Return the pointer of a new allocated Deck
	return &Stack{
		Cards: make([]*Card, 0),
	}
}

// ArrayID return cards id in array
func (s *Stack) ArrayID() []int {
	// Return array of cards id
	var arrayid []int
	for _, card := range s.Cards {
		arrayid = append(arrayid, card.ID)
	}
	return arrayid
}

// TopCardID returns the id of the card on top of the stack
func (s *Stack) TopCardID() int {
	return s.Cards[0].ID
}

// Add card to stack
func (s *Stack) Add(c *Card) {
	(*s).Cards = append((*s).Cards, c)
}

// Remove card id in the stack
func (s *Stack) Remove(id int) *Card {
	// Remove 0 takes the first card on the stack
	if id == 0 {
		//shift first card
		first := (*s).Cards[0]
		(*s).Cards = (*s).Cards[1:]
		return first
	}
	// Remove card from deck and return the deleted card
	for i, card := range s.Cards {
		if card.ID == id {
			(*s).Cards = append((*s).Cards[:i], (*s).Cards[i+1:]...)
			return card
		}
	}
	return nil
}

// Contains return true if stack contains card id
func (s *Stack) Contains(id int) bool {
	// Remove card from deck and return the deleted card
	for _, card := range s.Cards {
		if card.ID == id {
			return true
		}
	}
	return false
}

// AddStack to another Stack
func (s *Stack) AddStack(stack *Stack) {
	for _ = range stack.Cards {
		s.Add(stack.Remove(0))
	}
}

// Shuffle stack
func (s *Stack) Shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range s.Cards {
		j := rand.Intn(i + 1)
		(*s).Cards[i], (*s).Cards[j] = (*s).Cards[j], (*s).Cards[i]
	}
}

// IsValid return true if the stack can be dropped in game
func (s *Stack) IsValid() bool {
	if s.IsConsistent() {
		return s.IsSequence() || s.IsMultiple() || len(s.Cards) == 1
	}
	return false
}

// IsConsistent checks if the stack seems legit
func (s *Stack) IsConsistent() bool {
	if len(s.Cards) > 5 {
		return false
	}
	for _, card := range s.Cards {
		if card == nil {
			return false
		}
	}
	for i := 0; i < len(s.Cards)-1; i++ {
		for j := i + 1; j < len(s.Cards); j++ {
			if s.Cards[i].ID == s.Cards[j].ID {
				return false
			}
		}
	}
	return true
}

// IsMultiple return true if all cards have the same value and size > 1
func (s *Stack) IsMultiple() bool {
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

// IsSequence return true if >= 3 cards, have same symbol and value increments
func (s *Stack) IsSequence() bool {
	// If less than 3 cards, not a seq.
	if s.Len() < 3 {
		return false
	}
	// check if all cards have same Symbol/Colour
	for _, card := range s.Cards {
		if card.Colour != s.Cards[0].Colour {
			return false
		}
	}
	// Sort the stack
	sort.Sort(s)
	for i := range s.Cards {
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

// InitReference initialize the complete yaniv card game
func (s *Stack) InitReference() {
	j := 1
	colours := [...]string{"spade", "heart", "diam", "club"}
	for _, colour := range colours {
		for i := 1; i < 14; i++ {
			(*s).Cards = append((*s).Cards, CardNew(j, i, colour))
			j++
		}
	}
}
