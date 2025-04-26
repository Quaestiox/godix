# GODIX

Simple redis-like database implemented in Go.

## Support Command
- ping
- set
- get
- hset
- hget
- aof 
  - clean
- about

## Config
AOF is enabled by default. You can disable AOF data persistence by flag:
```cmd
./godix -aof=false
```