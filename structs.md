# Structs
## Game
Name String `json:"name"`
YanivAt int `json:"yaniv_at"`
MaxScore int `json:"max_score"`
State GameState `json:"state"`
Round int `json:"round"`
Players []Player `json:"players"`

## GameState
Started bool `json:"started"`
Ended bool `json:"ended"`

## Player
Name string `json:"name"`
Hand Deck `json:"hand"`
State PlayerState `json:"state"`
Score Int `json:"score"`

## PlayerState
Yaniv bool `json:"yaniv"`
Asaf bool `json:"asaf"` 
Playing bool `json:"playing"`
Ready bool `json:"Ready"`
Loser bool `json:"Loser"`

## Deck
Cards []Card `json:"cards"`
Value int `json:"value"`
Size int `json:"size"`

## Action
Name string `json:"name"`
Options ActionOptions `json:"options"`

## ActionOptions
Discarded []int `json:"discarded"`
Taken int `json:"taken"`

