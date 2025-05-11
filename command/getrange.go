package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
	"strconv"
)

func GetRange(args Args, config cfg.Config) resp.Val {
	if len(args) != 3 {
		return resp.NewError("ERR", "wrong number of arguments for 'getrange' command.")
	}

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	key := args[0].Value().(string)
	v1 := args[1].Value().(string)
	v2 := args[2].Value().(string)
	start, err1 := strconv.Atoi(v1)
	end, err2 := strconv.Atoi(v2)

	if err1 != nil || err2 != nil {
		return resp.NewError("ERR", "argument should be integers")
	}

	res := ""

	MapLock.RLock()

	l, ok := Map[key]
	v := l.value
	if !ok {
		MapLock.RUnlock()
		return resp.NewBulk(res)
	}
	length := len(v)
	if start < 0 {
		start = length + start
	}
	if end < 0 {
		end = length + end
	}
	if start > end {
		MapLock.RUnlock()
		return resp.NewBulk(res)
	}
	for i := start; i <= end; i++ {
		res += string(v[i])
	}
	MapLock.RUnlock()

	return resp.NewBulk(res)
}
