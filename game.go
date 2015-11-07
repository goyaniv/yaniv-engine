package main

import (
	"encoding/json"
)

// Game struct
type Game struct {
	Name    string      `json:"name"`
	Round   int         `json:"round"`
	State   *GameState  `json:"state"`
	Params  *GameParams `json:"params"`
	Stack   *Stack      `json:"stack"`
	Players []*Player   `json:"players"`
}

// GameState struct defines the state of the game
type GameState struct {
	Started bool `json:"started"`
	Ended   bool `json:"ended"`
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
func GameParamsNew() *GameParams {
	return &GameParams{}
}

// GameNew Initialize a Game object
func GameNew(name string) *Game {
	return &Game{
		Name:    name,
		State:   GameStateNew(),
		Params:  GameParamsNew(),
		Players: make([]*Player, 0),
		Stack:   StackNew(),
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