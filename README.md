# GODIX

Simple redis-like database implemented in Go.

## Supported Commands

view all commands: [command/README.md](https://github.com/Quaestiox/godix/blob/master/command/README.md)

## Config
AOF is enabled by default. You can disable AOF data persistence by argument `aof`:
```cmd
./godix -aof=false
```
Default port is 6379. You can change the port by argument `port`:
```cmd
./godix -port=6666
```

## The differences from redis

Keys of different types are stored separately. This means that you can have both a list called 'foo' and a hash at the same time.

Have some commands unique to godix.