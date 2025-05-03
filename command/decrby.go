package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
	"strconv"
)

func DecrBy(args Args, config cfg.Config) resp.Val {
	if len(args) != 2 {
		return resp.NewError("ERR", "wrong number of arguments for 'decrby' command.")
	}

	key := args[0].Value().(string)

	v := args[1].Value().(string)
	num, err := strconv.Atoi(v)
	if err != nil {
		return resp.NewError("ERR", "argument '"+v+"' is not an integer")
	}

	MapLock.RLock()
	sv, ok := Map[key]
	if !ok {
		Map[key] = NewSV("0")
		sv = Map[key]
	}
	i, err := strconv.Atoi(sv.Value())
	if err != nil {
		MapLock.RUnlock()
		return resp.NewError("ERR", "value is not an integer.")
	}
	n := i - num
	sv.setValue(strconv.Itoa(n))
	MapLock.RUnlock()

	return resp.NewInteger(n)
}
