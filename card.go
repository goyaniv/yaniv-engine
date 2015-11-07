package main

// Card struct
type Card struct {
	ID     int    `json:"id"`
	Value  int    `json:"_"`
	Colour string `json:"_"`
}

// CardNew returns a new card pointer
func CardNew(id int, value int, colour string) *Card {
	return &Card{id, value, colour}
}

// Weight returns the real value of card in a yaniv game
func (c *Card) Weight() int {
	if c.Value > 10 {
		return 10
	}
	return c.Value
}
