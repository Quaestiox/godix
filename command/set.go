package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func Set(args Args, config cfg.Config) resp.Val {
	if len(args) != 2 {
		return resp.NewError("ERR", "wrong number of arguments for 'set' command.")
	}

	key := args[0]
	value := args[1]

	if key.Type() != "bulk" || value.Type() != "bulk" {
		return resp.NewError("ERR", "Invalid arguments' format, expected bulk.")
	}

	MapLock.Lock()
	Map[key.Value().(string)] = NewSV(value.Value().(string))
	MapLock.Unlock()

	return resp.NewString("OK")

}
