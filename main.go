package main

import (
	"fmt"
)

func main() {
	s := ServerNew()
	p := PlayerNew("gégé")
	p.Hand.Add(CardNew(1, 1, "spade"))
	p.Hand.Add(CardNew(2, 2, "spade"))
	g := GameNew("pouette")
	g.AddPlayer(p)
	s.AddGame(g)
	json, _ := s.MarshalJSON()
	fmt.Println(string(json))
	json, _ = g.MarshalJSON()
	fmt.Println(string(json))
}
