package main

var s *Server

func main() {
	s = ServerNew()
	p1 := PlayerNew("josé")
	p2 := PlayerNew("robert")

	g := GameNew("166")
	g.AddPlayer(p1)
	g.AddPlayer(p2)
	s.AddGame(g)
	g.Start()
	p1.State.Ready = false
	p2.State.Ready = false
	p1.State.Playing = false
	p1.State.previousPlaying = false
	p2.State.Playing = true
	p2.State.previousPlaying = false
	g.State.Started = true
	g.Stack = StackNew()
	g.Stack.Add(CardNew(1, 1, "spade"))
	p1.Hand = DeckNew()
	p1.Hand.Add(CardNew(14, 1, "heart"))
	LaunchHTTP()
}
