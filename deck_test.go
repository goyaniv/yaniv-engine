package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeckValue(t *testing.T) {
	d := DeckNew()
	d.Add(CardNew(1, 1, "spade"))
	d.Add(CardNew(2, 2, "heart"))
	d.Add(CardNew(3, 3, "diam"))
	assert.Equal(t, d.Value(), 6, "Size must be 6")
}
func TestDeckSize(t *testing.T) {
	d := DeckNew()
	d.Add(CardNew(1, 1, "spade"))
	d.Add(CardNew(2, 2, "heart"))
	d.Add(CardNew(3, 3, "diam"))
	assert.Equal(t, d.Size(), 3, "Size must be 3")
}
func TestDeckMultiple(t *testing.T) {
	d := DeckNew()
	d.Add(CardNew(1, 1, "spade"))
	d.Add(CardNew(2, 2, "heart"))
	d.Add(CardNew(3, 3, "diam"))
	assert.False(t, d.IsMultiple(), "Deck must not be a Multiple")
	d.Remove(1)
	assert.Equal(t, d.Size(), 2, "Size must be 2")
	d.Remove(2)
	d.Add(CardNew(4, 3, "diam"))
	assert.True(t, d.IsMultiple(), "Deck must be a Multiple")
}
func TestDeckSequence(t *testing.T) {
	d := DeckNew()
	d.Add(CardNew(1, 1, "spade"))
	d.Add(CardNew(2, 2, "heart"))
	d.Add(CardNew(3, 3, "diam"))
	assert.False(t, d.IsSequence(), "Deck must not be a sequence")
	d.Remove(1)
	d.Remove(2)
	d.Add(CardNew(4, 4, "diam"))
	d.Add(CardNew(5, 5, "diam"))
	assert.True(t, d.IsSequence(), "Deck must be a sequence")
}
