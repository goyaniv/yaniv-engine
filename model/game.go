package model

import (
	"encoding/json"
)

type Game struct {
	Name    string      `json:"name"`
	Round   int         `json:"round"`
	State   *GameState  `json:"state"`
	Params  *GameParams `json:"params"`
	Stack   *Stack      `json:"stack"`
	Players []*Player   `json:"players"`
}

func GameStateNew() *GameState {
	return &GameState{}
}

func GameParamsNew() *GameParams {
	return &GameParams{}
}

type GameState struct {
	Started bool `json:"started"`
	Ended   bool `json:"ended"`
}

type GameParams struct {
	YanivAt  int `json:"yaniv_at"`
	MaxScore int `json:"max_score"`
}

func GameNew(name string) *Game {
	return &Game{
		Name:    name,
		State:   GameStateNew(),
		Params:  GameParamsNew(),
		Players: make([]*Player, 0),
		Stack:   StackNew(),
	}
}

func (g *Game) MarshalJSON() ([]byte, error) {
	return json.Marshal(*g)
}

// Return the Short JSON representation of the game
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

func (g *Game) PlayersNames() []string {
	playersnames := make([]string, 0)
	for _, player := range g.Players {
		playersnames = append(playersnames, player.Name)
	}
	return playersnames
}

func (g *Game) AddPlayer(p *Player) {
	(*g).Players = append((*g).Players, p)
}
