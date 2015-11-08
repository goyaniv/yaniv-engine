package main

import "errors"

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

// Play allows the player to discard cards and take one
func Play(g *Game, p *Player, discard []int, take int) error {

	if p != g.PlayerPlaying() {
		return errors.New("Not the player turn")
	}
	//if !discard.IsValid() {
	//	return errors.New("Invalid discarded cards")
	//}
	if !g.Stack.Contains(take) && take != 0 {
		return errors.New("Invalid taken card")
	}
	cardtaken := g.Stack.Remove(take)
	g.FlushStack()
	//g.Stack.AddStack(discard)
	p.Hand.Add(cardtaken)

	return nil
}
