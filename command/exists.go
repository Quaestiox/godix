package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func Exists(args Args, config cfg.Config) resp.Val {
	if len(args) != 1 {
		return resp.NewError("ERR", "wrong number of arguments for 'exists' command.")
	}

	key := args[0].Value().(string)

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	MapLock.RLock()
	_, ok := Map[key]
	MapLock.RUnlock()
	if !ok {
		return resp.NewInteger(0)
	}

	return resp.NewInteger(1)
}
