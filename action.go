package main

import (
	"errors"
	"fmt"
	"sort"
)

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
	p.State.asafRank = g.LastAsafRank() + 1
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
	if g.HasPlayerYaniv() {
		return errors.New("The round has stopped")
	}
	// If player is not playing, he just have the square
	// or the flash options
	if p != g.PlayerPlaying() {
		// if it seems to be a square or a flash
		if (g.Stack.Len()+len(discard) == 4) || (p == g.PlayerPreviousPlaying() && len(discard) == 1) {
			fmt.Println("on rentre dans le fast")
			var idcardstack []int
			for _, card := range g.Stack.Cards {
				idcardstack = append(idcardstack, card.ID)
			}
			for _, idcard := range discard {
				idcardstack = append(idcardstack, idcard)
			}
			sort.Ints(idcardstack)
			for i, idcard := range idcardstack {
				// if not first loop
				if i != 0 {
					// +13 because there is 13 cards and we are
					// searching for card with exact same values
					if idcardstack[i-1] != idcard+13 {
						return errors.New("Not the player turn")
					}
				}
			}
			// ok, all checks passed, he can fast his square or flash
			discardedcards, err := p.Discard(discard)
			if err != nil {
				return err
			}
			g.Stack.AddStack(discardedcards)
		}
	} else {
		fmt.Println("on rentre dans le non fast")

		if !g.Stack.Contains(take) && take != 0 {
			return errors.New("Invalid taken card")
		}
		discardedcards, err := p.Discard(discard)
		if err != nil {
			return err
		}
		if take == 0 {
			p.Hand.Add(g.hiddenstack.Remove(0))
		} else {
			p.Hand.Add(g.Stack.Remove(take))
		}
		g.FlushStack()
		g.Stack.AddStack(discardedcards)
	}
	return nil
}
