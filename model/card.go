package model

type Card struct {
	Id     int    `json:"id"`
	Value  int    `json:"_"`
	Colour string `json:"_"`
}

func CardNew(id int, value int, colour string) *Card {
	return &Card{id, value, colour}
}

func (c *Card) Weight() int {
	if c.Value > 10 {
		return 10
	}
	return c.Value
}
