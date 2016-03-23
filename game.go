package main

import (
	"encoding/json"
	"errors"
)

// Game struct
type Game struct {
	Name        string      `json:"name"`
	Round       int         `json:"round"`
	State       *GameState  `json:"state"`
	Params      *GameParams `json:"params"`
	Stack       *Stack      `json:"stack"`
	Players     []*Player   `json:"players"`
	stacktrash  *Stack
	hiddenstack *Stack
}

// GameState struct defines the state of the game
type GameState struct {
	Started bool `json:"started"`
	Ended   bool `json:"ended"`
	Flushed bool `json:"flushed"`
}

// GameParams defines the parameters at the game creation
type GameParams struct {
	YanivAt  int `json:"yaniv_at"`
	MaxScore int `json:"max_score"`
}

// GameStateNew Initialize a GameState object
func GameStateNew() *GameState {
	return &GameState{}
}

// GameParamsNew Initialize a GameParams object
func GameParamsNew(yanivat int, maxscore int) *GameParams {
	return &GameParams{
		YanivAt:  yanivat,
		MaxScore: maxscore,
	}
}

// GameNew Initialize a Game object
func GameNew(name string) *Game {
	hiddenstack := StackNew()
	hiddenstack.InitReference()
	hiddenstack.Shuffle()

	stack := StackNew()
	stack.Add(hiddenstack.Remove(0))
	return &Game{
		Name:        name,
		State:       GameStateNew(),
		Params:      GameParamsNew(5, 100),
		Players:     make([]*Player, 0),
		Stack:       stack,
		hiddenstack: hiddenstack,
		stacktrash:  StackNew(),
	}
}

// MarshalJSON render full JSON representation of a Game object
func (g *Game) MarshalJSON() ([]byte, error) {
	return json.Marshal(*g)
}

// MarshalJSONShort render the Short JSON representation of the Game object
func (g *Game) MarshalJSONShort() ([]byte, error) {
	return json.Marshal(struct {
		Name         string      `json:"name"`
		Round        int         `json:"round"`
		PlayersNames []string    `json:"players_names"`
		GameState    *GameState  `json:"state"`
		Params       *GameParams `json:"params"`
	}{
		Name:         g.Name,
		Round:        g.Round,
		PlayersNames: g.PlayersNames(),
		GameState:    g.State,
		Params:       g.Params,
	})
}

// PlayersNames return a string array with all the player name in the game
func (g *Game) PlayersNames() []string {
	var playersnames []string
	for _, player := range g.Players {
		playersnames = append(playersnames, player.Name)
	}
	return playersnames
}

// AddPlayer appends Player object to the Game Object
func (g *Game) AddPlayer(p *Player) {
	(*g).Players = append((*g).Players, p)
}

// HasPlayerYaniv checks if a player has yanived
func (g *Game) HasPlayerYaniv() bool {
	for _, p := range g.Players {
		if p.State.Yaniv == true {
			return true
		}
	}
	return false
}

// LastAsafRank return the ranking of the last asafer
func (g *Game) LastAsafRank() int {
	higherRank := 0
	for _, p := range g.Players {
		if p.State.asafRank > higherRank {
			higherRank = p.State.asafRank
		}
	}
	return higherRank
}

//PlayerPlaying return the player playing
func (g *Game) PlayerPlaying() *Player {
	for _, p := range g.Players {
		if p.State.Playing {
			return p
		}
	}
	return nil
}

//PlayerPreviousPlaying return the player who
//previously played
func (g *Game) PlayerPreviousPlaying() *Player {
	for _, p := range g.Players {
		if p.State.previousPlaying {
			return p
		}
	}
	return nil
}

// FindPlayer returns the player if exists
func (g *Game) FindPlayer(name string) *Player {
	for _, player := range g.Players {
		if player.Name == name {
			return player
		}
	}
	return nil
}

//FlushStack moves the previous visible cards in trash
func (g *Game) FlushStack() {
	if g.hiddenstack.Len() == 0 {
		//g.hiddenstack = g.stacktrash
		g.hiddenstack.AddStack(g.stacktrash)

		g.hiddenstack.Shuffle()
		//g.stacktrash = StackNew()
	}
	g.stacktrash.AddStack(g.Stack)
}

// RemovePlayer removes player if exists
func (g *Game) RemovePlayer(name string) error {
	for i, player := range g.Players {
		if player.Name == name {
			(*g).Players = append((*g).Players[:i], (*g).Players[i+1:]...)
			return nil
		}
	}
	return errors.New("Player does not exists in this game")
}

// NextPlayer change the player turn
func (g *Game) NextPlayer() {
	ppp := g.PlayerPreviousPlaying()
	pp := g.PlayerPlaying()
	pa, _ := g.PlayerAfter(g.PlayerPlaying())
	if ppp != nil {
		ppp.State.previousPlaying = false
	} else {
		pp.State.previousPlaying = true
	}
	pp.State.Playing = false
	pa.State.Playing = true
}

// PlayerAfter return playing after a player
func (g *Game) PlayerAfter(p *Player) (*Player, error) {
	for i, player := range g.Players {
		if player == p {
			// last player in slice
			if i == len(g.Players)-1 {
				return g.Players[0], nil
			}
			return g.Players[i+1], nil
		}
	}
	return nil, errors.New("Next player error")
}

//IsAllPlayersReady return true if all players in game are ready
func (g *Game) IsAllPlayersReady() bool {
	for _, p := range g.Players {
		if !p.State.Ready {
			return false
		}
	}
	return true
}

// Start the game and give cards to all players
func (g *Game) Start() error {
	if !g.IsAllPlayersReady() {
		return errors.New("All the player are not ready")
	}
	if len(g.Players) < 2 {
		return errors.New("There is not enought players")
	}
	g.State.Started = true
	g.Players[0].State.Playing = true
	for _, player := range g.Players {
		for i := 0; i < 5; i++ {
			player.Hand.Add(g.hiddenstack.Remove(0))
		}
	}
	return nil
}
