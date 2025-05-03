package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
	"strconv"
)

func Decr(args Args, config cfg.Config) resp.Val {
	if len(args) != 1 {
		return resp.NewError("ERR", "wrong number of arguments for 'decr' command.")
	}
	key := args[0].Value().(string)
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
	n := i - 1
	sv.setValue(strconv.Itoa(n))
	MapLock.RUnlock()

	return resp.NewInteger(n)
}
