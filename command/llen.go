package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func LLen(args Args, config cfg.Config) resp.Val {
	if len(args) != 1 {
		return resp.NewError("ERR", "wrong number of arguments for 'llen' command.")
	}

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	list := args[0].Value().(string)

	LMapLock.RLock()
	l, ok := LMap[list]
	if !ok {
		LMapLock.RUnlock()
		return resp.NewInteger(0)
	}
	length := len(l)
	LMapLock.RUnlock()

	return resp.NewInteger(length)
}
