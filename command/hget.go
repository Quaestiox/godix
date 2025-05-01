package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func HGet(args Args, config cfg.Config) resp.Val {
	if len(args) != 2 {
		return resp.NewError("ERR", "wrong number of arguments for 'hget' command.")
	}

	if args[0].Type() != "bulk" || args[1].Type() != "bulk" {
		return resp.NewError("ERR", "Invalid arguments' format, expected bulk.")
	}

	hash := args[0].Value().(string)
	key := args[1].Value().(string)

	HMapLock.RLock()
	value, ok := HMap[hash][key]
	HMapLock.RUnlock()
	if !ok {
		return resp.NewNullBulk()
	}
	return resp.NewBulk(value)

}
