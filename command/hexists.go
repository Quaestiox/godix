package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func HExists(args Args, config cfg.Config) resp.Val {

	if len(args) != 2 {
		return resp.NewError("ERR", "wrong number of arguments for 'hexists' command.")
	}

	hash := args[0].Value().(string)
	key := args[1].Value().(string)

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	HMapLock.RLock()
	_, ok := HMap[hash][key]
	HMapLock.RUnlock()
	if !ok {
		return resp.NewInteger(0)
	}

	return resp.NewInteger(1)

}
