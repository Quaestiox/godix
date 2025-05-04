package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func SIsMemeber(args Args, config cfg.Config) resp.Val {
	if len(args) != 2 {
		return resp.NewError("ERR", "wrong number of arguments for 'sismemeber' command.")
	}

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	s := args[0].Value().(string)
	member := args[1].Value().(string)
	res := 0
	SMapLock.RLock()
	set, ok := SMap[s]
	if !ok {
		SMapLock.RUnlock()
		return resp.NewInteger(0)
	}
	if set[member] {
		res = 1
	}

	SMapLock.RUnlock()

	return resp.NewInteger(res)
}
