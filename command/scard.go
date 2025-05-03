package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func SCard(args Args, config cfg.Config) resp.Val {
	if len(args) != 1 {
		return resp.NewError("ERR", "wrong number of arguments for 'scard' command.")
	}

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	value := args[0].Value().(string)

	SMapLock.RLock()

	count := 0
	set, ok := SMap[value]
	if !ok {
		SMapLock.RUnlock()
		return resp.NewInteger(0)
	}
	for range set {
		count++
	}

	SMapLock.RUnlock()

	return resp.NewInteger(count)
}
