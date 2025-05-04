package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func RPush(args Args, config cfg.Config) resp.Val {
	if len(args) < 2 {
		return resp.NewError("ERR", "wrong number of arguments for 'rpush' command.")
	}

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	list := args[0].Value().(string)
	el := []string{}
	for i := 1; i < len(args); i++ {
		el = append(el, args[i].Value().(string))
	}

	LMapLock.Lock()
	_, ok := LMap[list]
	if !ok {
		LMap[list] = []string{}
	}
	LMap[list] = append(LMap[list], el...)

	LMapLock.Unlock()

	return resp.NewInteger(len(LMap[list]))
}
