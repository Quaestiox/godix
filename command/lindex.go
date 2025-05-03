package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
	"strconv"
)

func LIndex(args Args, config cfg.Config) resp.Val {
	if len(args) != 2 {
		return resp.NewError("ERR", "wrong number of arguments for 'lindex' command.")
	}

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	list := args[0].Value().(string)
	value := args[1].Value().(string)
	index, err := strconv.Atoi(value)
	if err != nil {
		return resp.NewError("ERR", "argument '"+value+"' is not an integer")
	}

	LMapLock.RLock()
	l, ok := LMap[list]
	if !ok {
		LMapLock.RUnlock()
		return resp.NewError("ERR", "no list called '"+list+"'")
	}
	length := len(l)
	if index < 0 {
		index = length + index
	}
	if index < 0 || index >= length {
		LMapLock.RUnlock()
		return resp.NewNullBulk()
	}
	v := l[index]

	LMapLock.RUnlock()

	return resp.NewBulk(v)
}
