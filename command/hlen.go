package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func HLen(args Args, config cfg.Config) resp.Val {
	if len(args) != 1 {
		return resp.NewError("ERR", "wrong number of arguments for 'hget' command.")
	}
	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	hash := args[0].Value().(string)

	HMapLock.RLock()
	h, ok := HMap[hash]
	if !ok {
		HMapLock.RUnlock()
		return resp.NewInteger(0)
	}
	length := len(h)
	HMapLock.RUnlock()

	return resp.NewInteger(length)

}
