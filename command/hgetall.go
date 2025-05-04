package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func HGetAll(args Args, config cfg.Config) resp.Val {
	if len(args) != 1 {
		return resp.NewError("ERR", "wrong number of arguments for 'hgetall' command.")
	}

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	hash := args[0].Value().(string)

	res := []resp.Val{}

	HMapLock.RLock()
	h, ok := HMap[hash]
	if !ok {
		HMapLock.RUnlock()
		return resp.NewArray(res...)
	}
	for k, v := range h {
		res = append(res, resp.NewBulk(k))
		res = append(res, resp.NewBulk(v))
	}
	HMapLock.RUnlock()
	return resp.NewArray(res...)

}
