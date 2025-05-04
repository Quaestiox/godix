package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func RPop(args Args, config cfg.Config) resp.Val {
	if len(args) != 1 {
		return resp.NewError("ERR", "wrong number of arguments for 'rpop' command.")
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
	length := len(LMap[list])
	if length == 0 {
		LMapLock.Unlock()
		return resp.NewNullBulk()
	}
	v := LMap[list][length-1]
	LMap[list] = LMap[list][:length-1]

	LMapLock.Unlock()

	return resp.NewBulk(v)
}
