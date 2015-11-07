package main

import (
	"fmt"
)

func main() {
	s := ServerNew()
	p := PlayerNew("gégé")
	p.Hand.Add(CardNew(1, 1, "spade"))
	g := GameNew("pouette")
	g.AddPlayer(p)
	s.AddGame(g)
	json, _ := s.MarshalJSON()
	fmt.Println(string(json))
}
