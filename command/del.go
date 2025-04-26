package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func Del(args Args, config cfg.Config) resp.Val {
	if len(args) != 1 {
		return resp.NewError("ERR", "wrong number of arguments for 'del' command.")
	}

	key := args[0].Value().(string)

	MapLock.Lock()
	delete(Map, key)
	MapLock.Unlock()

	return resp.NewString("OK")
}
