# GODIX

Simple redis-like database implemented in Go.

## Support Command
- ping
- echo
- set
- get
- del
- hset
- hget
- hdel
- aof 
  - clean

... *view all commands in command/README.md*


## Config
AOF is enabled by default. You can disable AOF data persistence by argument `aof`:
```cmd
./godix -aof=false
```
Default port is 6379. You can change the port by argument `port`:
```cmd
./godix -port=6666
```