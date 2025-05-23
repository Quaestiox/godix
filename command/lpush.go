package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func LPush(args Args, config cfg.Config) resp.Val {
	if len(args) < 2 {
		return resp.NewError("ERR", "wrong number of arguments for 'lpush' command.")
	}

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	list := args[0].Value().(string)
	el := []string{}
	for i := len(args) - 1; i >= 1; i-- {
		el = append(el, args[i].Value().(string))
	}

	LMapLock.Lock()
	_, ok := LMap[list]
	if !ok {
		LMap[list] = []string{}
	}
	LMap[list] = append(el, LMap[list]...)

	LMapLock.Unlock()

	return resp.NewInteger(len(LMap[list]))
}
