package main

import (
	"fmt"
	"github.com/goyaniv/yaniv-engine/model"
)

func main() {

	s := model.ServerNew()
	p := model.PlayerNew("gégé")
	p.Hand.Add(model.CardNew(1, 1, "spade"))
	g := model.GameNew("pouette")
	g.AddPlayer(p)
	json, _ := s.MarshalJSON()
	fmt.Println(string(json))
}
