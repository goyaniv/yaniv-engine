package test

import (
	"github.com/goyaniv/yaniv-engine/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCardWeight(t *testing.T) {
  c := model.CardNew(1,2,"spade")
	assert.Equal(t, c.Weight(), 2, "Size must be 2")
  c = model.CardNew(1,12,"spade")
	assert.Equal(t, c.Weight(), 10, "Size must be 10")
}
