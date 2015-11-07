package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCardWeight(t *testing.T) {
	c := CardNew(1, 2, "spade")
	assert.Equal(t, c.Weight(), 2, "Size must be 2")
	c = CardNew(1, 12, "spade")
	assert.Equal(t, c.Weight(), 10, "Size must be 10")
}
