# Redis commands
## Cards management
### Add cards in deck
```
rpush player:bob:cards 8
rpush player:bob:cards 16
rpush player:bob:cards 6
rpush player:bob:cards 1
```
### Get All cards in deck
`lrange player:bob:cards 0 -1`
### Remove card in deck
`lrem player:bob:cards 1 8`
