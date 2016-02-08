package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// LaunchHTTP initialize http routes
func LaunchHTTP() {
	r := gin.Default()

	// Games
	r.GET("/game/:game", handlergetgame)
	r.GET("/games", handlerlistgames)
	r.POST("/games", handlercreategame)
	//r.DELETE("/game/:game", handlerdeletegame)

	// Players
	r.POST("/game/:game/players", handleraddplayer)
	r.PUT("/game/:game/player/:player", handlerupdateplayer)
	r.DELETE("/game/:game/player/:player", handlerremoveplayer)

	// Actions
	r.POST("/game/:game/player/:player/action/asaf", handleractionasaf)
	r.POST("/game/:game/player/:player/action/yaniv", handleractionyaniv)
	r.POST("/game/:game/player/:player/action/takecard", handleractiontakecard)

	r.Run(":3001")
}

func handlerlistgames(c *gin.Context) {
	fmt.Println(s)
	c.JSON(200, s)
}

func handlergetgame(c *gin.Context) {
	game := s.FindGame(c.Param("game"))
	if game == nil {
		c.JSON(404, gin.H{"message": "game not found"})
	} else {
		c.JSON(200, game)
	}
}

func handlercreategame(c *gin.Context) {
	game := GameNew("")

	if c.BindJSON(game) != nil {
		c.JSON(400, gin.H{"message": "cannot decode the json sen t"})
	}
	if s.FindGame(game.Name) != nil {
		c.JSON(409, gin.H{"message": "game already exists"})
	} else {
		s.AddGame(game)
		c.JSON(200, game)
	}
}

func handleraddplayer(c *gin.Context) {
	game := s.FindGame(c.Param("game"))
	if game == nil {
		c.JSON(404, gin.H{"message": "game not found"})
	} else {
		player := PlayerNew("")
		if c.BindJSON(player) != nil {
			c.JSON(400, gin.H{"message": "cannot decode the json sent"})
		} else {
			if game.FindPlayer(player.Name) != nil {
				c.JSON(409, gin.H{"message": "player already exists in this game"})
			} else {
				game.AddPlayer(player)
				c.JSON(200, game)
			}
		}
	}
}

func handlerremoveplayer(c *gin.Context) {
	game := s.FindGame(c.Param("game"))
	if game == nil {
		c.JSON(404, gin.H{"message": "game not found"})
	} else {
		if game.FindPlayer(c.Param("player")) == nil {
			c.JSON(404, gin.H{"message": "player does not exists in this game"})
		} else {
			game.RemovePlayer(c.Param("player"))
			c.JSON(200, game)
		}
	}
}

func handlerupdateplayer(c *gin.Context) {
	game := s.FindGame(c.Param("game"))
	if game == nil {
		c.JSON(404, gin.H{"message": "game not found"})
	} else {
		player := game.FindPlayer(c.Param("player"))
		if player == nil {
			c.JSON(404, gin.H{"message": "player does not exists in this game"})
		} else {
			if c.BindJSON(player) == nil {
				c.JSON(200, game)
			} else {
				c.JSON(400, gin.H{"message": "cannot decode json sent"})
			}
		}
	}
}

func handleractionasaf(c *gin.Context) {
	game := s.FindGame(c.Param("game"))
	if game == nil {
		c.JSON(404, gin.H{"message": "game not found"})
	} else {
		player := game.FindPlayer(c.Param("player"))
		if player == nil {
			c.JSON(404, gin.H{"message": "player does not exists in this game"})
		} else {
			if game.State.Started {
				if Asaf(game, player) {
					c.JSON(200, game)
				} else {
					c.JSON(403, gin.H{"message": "player cannot asaf"})
				}
			} else {
				c.JSON(403, gin.H{"message": "game not started"})
			}
		}
	}
}

func handleractionyaniv(c *gin.Context) {
	game := s.FindGame(c.Param("game"))
	if game == nil {
		c.JSON(404, gin.H{"message": "game not found"})
	} else {
		player := game.FindPlayer(c.Param("player"))
		if player == nil {
			c.JSON(404, gin.H{"message": "player does not exists in this game"})
		} else {
			if game.State.Started {
				if Yaniv(game, player) {
					c.JSON(200, game)
				} else {
					c.JSON(403, gin.H{"message": "player cannot yaniv"})
				}
			} else {
				c.JSON(403, gin.H{"message": "game not started"})
			}
		}
	}
}

type incomingjson struct {
	Discard []int `json:"discard"`
	Take    int   `json:"take"`
}

func handleractiontakecard(c *gin.Context) {
	game := s.FindGame(c.Param("game"))
	if game == nil {
		c.JSON(404, gin.H{"message": "game not found"})
	} else {
		player := game.FindPlayer(c.Param("player"))
		if player == nil {
			c.JSON(404, gin.H{"message": "player does not exists in this game"})
		} else {
			var json incomingjson
			if c.BindJSON(&json) != nil {
				c.JSON(400, gin.H{"message": "cannot decode json sent"})
			} else {
				if game.State.Started {
					if err := Play(game, player, json.Discard, json.Take); err == nil {
						c.JSON(200, game)
					} else {
						c.JSON(403, gin.H{"message": err.Error()})
					}
				} else {
					c.JSON(403, gin.H{"message": "game not started"})
				}
			}
		}
	}
}
