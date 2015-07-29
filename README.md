# yaniv-engine
Fast, battle tested, Yaniv Engine with fully REST API

# REST API

## Game
### List All games
`GET` /game
```
{
  [
    {
      "name": "toto",
      "yaniv_at": 5,
      "max_score": 100,
      "started": true,
      "ended": false,
      "round": 1,
      "players": [
        {
          "name": "Yann"
        },
        {
          "name": "Etienne"
        }
      ]
    },
    {
      "name": "tata",
      "yaniv_at": 5,
      "max_score": 100,
      "started": true,
      "ended": false,
      "round": 17,
      "players": [
        {
          "name": "Brousse"
        },
        {
          "name": "raphZen"
        }
      ]
    }
  ]
}
```
### Get game info
`GET` /game/:name?viewer=Yann
```
{
  {
    "name": "toto",
    "yaniv_at": 5,
    "max_score": 100,
    "started": true,
    "ended": false,
    "round": 1,
    "stack": [{
        "id": 5,
    }],
    "user": {
        "name": "Yann",
        "hand": {
          "cards": [1,2,3],
          "value": 7,
          "size": 5
        },
        "score": 23,
        "state": {
          "yaniv": false,
          "asaf": false,
          "playing": false,
          "ready": true
        }
    },
    "opponents": [{
        "name": "Etienne",
        "hand": {
          "size": 5
        },
        "score": 23,
        "state": {
          "yaniv": false,
          "asaf": false,
          "playing": false,
          "ready": true
        }
    }],
    "action_last": {
      "name": "card_take",
      "options": {
        "discarded": [3, 4, 5],
        "taken": 6
      }
    } 
  }
}
```
### Create game
`POST` /game
```
{
  "name": "toto",
  "yaniv_at": 5,
  "max_score": 100
}
```
### Delete game
`DELETE` /game/:name

## Player
### Get player info
`GET` /game/:name/player/:name
### Add Player
`POST` /game/:name/player
```
{
  "name": "Yann",
  "visitor": true,
  "ready": false
}
```
### Delete Player
`DELETE` /game/:name/player/:name

### Update Player
`UPDATE` /game/:name/player/:name
{
  "name": "Etienne"
}

## Action
### Take card
`POST` /game/:name/player/:name/action/takecard
```
{
  "take": 3,
  "discard": [4,5,6]
}
```
### Yaniv
`POST` /game/:name/player/:name/action/yaniv

### Asaf
`POST` /game/:name/player/:name/action/asaf
```
{
  "try_asaf": true
}
```
```
