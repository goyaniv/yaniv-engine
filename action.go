package main

// Asaf change state of player if he is allowed to asaf
func Asaf(g *Game, p *Player) bool {
	// check if anybody yanived yet
	if !g.HasPlayerYaniv() {
		return false
	}
	// check if player has the right deck weight for asaf
	if p.Hand.Value() > g.Params.YanivAt {
		return false
	}

	p.State.Asaf = true
	p.State.AsafRank = g.LastAsafRank() + 1
	return true
}

// Yaniv change state of player if he is allowed to yaniv
func Yaniv(g *Game, p *Player) bool {
	// check if anybody yanived yet
	if g.HasPlayerYaniv() {
		return false
	}
	if p.Hand.Value() > g.Params.YanivAt {
		return false
	}

	p.State.Yaniv = true
	return true
}
