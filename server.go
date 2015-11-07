package main

import (
	"strings"
)

// Server manages games
type Server struct {
	Games []*Game `json:"games"`
}

// ServerNew initialize a new Server
func ServerNew() *Server {
	return &Server{make([]*Game, 0)}
}

// MarshalJSON return the JSON representation of server struct
func (s *Server) MarshalJSON() ([]byte, error) {
	var json string
	var gametab []string
	jsonerr := []byte("[]")
	for _, g := range s.Games {
		gamebyte, err := g.MarshalJSONShort()
		if err != nil {
			return jsonerr, err
		}
		gamestring := string(gamebyte)
		gametab = append(gametab, gamestring)
	}
	json = strings.Join(gametab, ",")
	json = strings.Join([]string{"[", json, "]"}, "")
	return []byte(json), nil
}

// AddGame add game to the server
func (s *Server) AddGame(g *Game) {
	(*s).Games = append((*s).Games, g)
}

// FindGame search for a game object in the server
func (s *Server) FindGame(name string) *Game {
	for _, game := range s.Games {
		if game.Name == name {
			return game
		}
	}
	return nil
}
