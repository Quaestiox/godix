package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func LPop(args Args, config cfg.Config) resp.Val {
	if len(args) != 1 {
		return resp.NewError("ERR", "wrong number of arguments for 'lpop' command.")
	}

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	list := args[0].Value().(string)

	LMapLock.Lock()
	_, ok := LMap[list]
	if !ok {
		LMapLock.Unlock()
		return resp.NewError("ERR", "no list called '"+list+"'")
	}
	if len(LMap[list]) == 0 {
		LMapLock.Unlock()
		return resp.NewNullBulk()
	}
	v := LMap[list][0]
	LMap[list] = LMap[list][1:]

	LMapLock.Unlock()

	return resp.NewBulk(v)
}
