package main

import (
	"errors"
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
	nbflash, errflash := canFlash(g, discard)

	// If player is not playing, he just have the square
	// or the flash options
	if p != g.PlayerPlaying() {
		// he fails flash
		if errflash != nil {
			return errflash
		}
		// big flash or flash?
		if nbflash == 4 || (nbflash == 2 && g.PlayerPreviousPlaying() == p) {
			discardedcards, err := p.Discard(discard)
			if err != nil {
				return err
			}
			g.Stack.AddStack(discardedcards)
		}
	} else {
		if !g.Stack.Contains(take) && take != 0 {
			return errors.New("Invalid taken card")
		}
		if errflash == nil && nbflash == 4 {
			discardedcards, err := p.Discard(discard)
			if err != nil {
				return err
			}
			g.Stack.AddStack(discardedcards)
		} else {
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
	}
	return nil
}

func canFlash(g *Game, discard []int) (int, error) {
	// if it seems to be a square or a flash
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
			if idcardstack[i-1]%13 != idcard%13 {
				return 0, errors.New("Can't Flash theses cards")
			}
		}
	}
	return len(idcardstack), nil
}
