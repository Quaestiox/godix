package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func SMemebers(args Args, config cfg.Config) resp.Val {
	if len(args) != 1 {
		return resp.NewError("ERR", "wrong number of arguments for 'smemebers' command.")
	}

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	value := args[0].Value().(string)
	res := []resp.Val{}

	SMapLock.RLock()
	set, ok := SMap[value]
	if !ok {
		SMapLock.RUnlock()
		return resp.NewArray()
	}
	for member := range set {
		res = append(res, resp.NewBulk(member))
	}

	SMapLock.RUnlock()

	return resp.NewArray(res...)
}
