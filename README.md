# yaniv-engine
Fast, battle tested, Yaniv Engine with fully REST API

# REST API

## Game
### Create
`POST /games`
#### Request data
```
{
	"name": String,
	"params": {
		"max_score": Number,
		"yaniv_at": Number
	}
}
```
#### Response data
=> `GET /game/:name`

### Delete
`DELETE /game/:name`

### Get
`GET /game/:name`
#### Response data
```
{
	"name": String,
	"params": {
		"max_score": Number,
		"yaniv_at": Number
	},
	"players": [{
		"hand": {
			"cards": [Number],
			"size": Number,
			"value": Number
		},
		"name": String,
		"score": Number,
		"state": {
			"asaf": Boolean,
			"loser": Boolean,
			"playing": Boolean,
			"ready": Boolean,
			"yaniv": Boolean
		}
	}],
	"round": Number,
	"stack": {
		"cards": [Number]
	},
	"state": {
		"ended": Boolean,
		"started": Boolean,
		"yaniv_battle": Boolean
	}
}
```

### List
`GET /games`
#### Response data
```
[{
	"name": String,
	"params": {
		"max_score": Number,
		"yaniv_at": Number
	},
	"players": [{
		"name": String
	}],
	"round": Number,
	"state": {
		"ended": Boolean,
		"started": Boolean
	}
}]
```

## Player
### Add
`POST /game/:name/players`
#### Request data
```
{
	"name": String
}
```
#### Response data
=> `GET /game/:name`

### Remove
`DELETE /game/:name/player/:name`
#### Response data
=> `GET /game/:name`

### Update
`PUT /game/:name/player/:name`
#### Request data
```
{
	"name": String#Optional,
	"state": {
		"ready" Boolean
	}#Optional
}
```
#### Response data
=> `GET /game/:name`

## Action
### Asaf
`POST /game/:name/player/:name/action/asaf`
#### Request data
```
Boolean
```
#### Response data
=> `GET /game/:name`

### Take card
`POST /game/:name/player/:name/action/takecard`
#### Request data
```
{
	"discard": [Number],
	"take": Number
}
```
#### Response data
=> `GET /game/:name`

### Yaniv
`POST /game/:name/player/:name/action/yaniv`
#### Response data
=> `GET /game/:name`
