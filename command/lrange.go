package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
	"strconv"
)

func LRange(args Args, config cfg.Config) resp.Val {
	if len(args) != 3 {
		return resp.NewError("ERR", "wrong number of arguments for 'smemebers' command.")
	}

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	list := args[0].Value().(string)
	v1 := args[1].Value().(string)
	v2 := args[2].Value().(string)
	start, err1 := strconv.Atoi(v1)
	end, err2 := strconv.Atoi(v2)

	if err1 != nil || err2 != nil {
		return resp.NewError("ERR", "argument should be integers")
	}

	res := []resp.Val{}

	LMapLock.RLock()

	l, ok := LMap[list]
	if !ok {
		LMapLock.RUnlock()
		return resp.NewArray(res...)
	}
	length := len(l)
	if start < 0 {
		start = length + start
	}
	if end < 0 {
		end = length + end
	}
	if start > end {
		LMapLock.RUnlock()
		return resp.NewArray(res...)
	}
	for i := start; i <= end; i++ {
		res = append(res, resp.NewBulk(l[i]))
	}
	LMapLock.RUnlock()

	return resp.NewArray(res...)
}
