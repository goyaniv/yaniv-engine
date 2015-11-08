package main

import "github.com/gin-gonic/gin"

// LaunchHTTP initialize http routes
func LaunchHTTP() {
	r := gin.Default()

	// Games
	r.GET("/game/:game", gingetgame)
	r.GET("/games", ginlistgames)
	r.POST("/games", gincreategame)
	//r.DELETE("/game/:game", gindeletegame)

	// Players
	r.POST("/game/:game/players", ginaddplayer)
	r.PUT("/game/:game/player/:player", ginupdateplayer)
	r.DELETE("/game/:game/player/:player", ginremoveplayer)

	// Actions
	r.POST("/game/:game/player/:player/action/asaf", ginactionasaf)
	r.POST("/game/:game/player/:player/action/yaniv", ginactionyaniv)
	r.POST("/game/:game/player/:player/action/takecard", ginactiontakecard)

	r.Run(":8000")
}

func ginlistgames(c *gin.Context) {
	c.JSON(200, s)
}

func gingetgame(c *gin.Context) {
	game := s.FindGame(c.Param("game"))
	if game == nil {
		c.JSON(404, gin.H{"message": "game not found"})
	} else {
		c.JSON(200, game)
	}
}

func gincreategame(c *gin.Context) {
	game := GameNew("")

	if c.BindJSON(game) != nil {
		c.JSON(400, gin.H{"message": "cannot decode the json sent"})
	}
	if s.FindGame(game.Name) != nil {
		c.JSON(409, gin.H{"message": "game already exists"})
	} else {
		s.AddGame(game)
		c.JSON(200, game)
	}
}

func ginaddplayer(c *gin.Context) {
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

func ginremoveplayer(c *gin.Context) {
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

func ginupdateplayer(c *gin.Context) {
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

func ginactionasaf(c *gin.Context) {
	game := s.FindGame(c.Param("game"))
	if game == nil {
		c.JSON(404, gin.H{"message": "game not found"})
	} else {
		player := game.FindPlayer(c.Param("player"))
		if player == nil {
			c.JSON(404, gin.H{"message": "player does not exists in this game"})
		} else {
			if Asaf(game, player) {
				c.JSON(200, game)
			} else {
				c.JSON(403, gin.H{"message": "player cannot asaf"})
			}
		}
	}
}

func ginactionyaniv(c *gin.Context) {
	game := s.FindGame(c.Param("game"))
	if game == nil {
		c.JSON(404, gin.H{"message": "game not found"})
	} else {
		player := game.FindPlayer(c.Param("player"))
		if player == nil {
			c.JSON(404, gin.H{"message": "player does not exists in this game"})
		} else {
			if Yaniv(game, player) {
				c.JSON(200, game)
			} else {
				c.JSON(403, gin.H{"message": "player cannot yaniv"})
			}
		}
	}
}

type incomingjson struct {
	Discard []int `json:"discard"`
	Take    int   `json:"take"`
}

func ginactiontakecard(c *gin.Context) {
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
				if err := Play(game, player, json.Discard, json.Take); err == nil {
					c.JSON(200, game)
				} else {
					c.JSON(403, gin.H{"message": err.Error()})
				}
			}
		}
	}
}
