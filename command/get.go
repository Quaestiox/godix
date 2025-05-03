package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func Get(args Args, config cfg.Config) resp.Val {
	if len(args) != 1 {
		return resp.NewError("ERR", "wrong number of arguments for 'get' command.")
	}
	key := args[0]
	MapLock.RLock()
	sv, ok := Map[key.Value().(string)]
	MapLock.RUnlock()
	if !ok {
		return resp.NewNullBulk()
	}
	return resp.NewBulk(sv.Value())
}
