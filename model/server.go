package model

import (
	"strings"
)

type Server struct {
	Games []*Game `json:"games"`
}

func ServerNew() *Server {
	return &Server{make([]*Game, 0)}
}

func (s *Server) MarshalJSON() ([]byte, error) {
	// Return the JSON representation of the server
	var json string
	var gametab []string
	jsonerr := []byte("[]")
	for _, g := range s.Games {
		gamebyte, err := g.MarshalJSONShort()
		gamestring := string(gamebyte)
		if err != nil {
			return jsonerr, err
		}
		gametab = append(gametab, gamestring)
	}
	json = strings.Join(gametab, ",")
	json = strings.Join([]string{"[", json, "]"}, "")
	return []byte(json), nil
}

func (s *Server) AddGame(g *Game) {
	(*s).Games = append((*s).Games, g)
}
