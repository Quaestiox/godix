package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
	"time"
)

func PTTL(args Args, config cfg.Config) resp.Val {
	if len(args) != 1 {
		return resp.NewError("ERR", "wrong number of arguments for 'pttl' command.")
	}

	key := args[0].Value().(string)

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	ExpireRecordLock.RLock()
	exp, ok := ExpireRecord[key]
	if !ok {
		ExpireRecordLock.RUnlock()
		return resp.NewInteger(-1)
	}
	res := int(exp.Sub(time.Now()).Milliseconds())
	ExpireRecordLock.RUnlock()

	return resp.NewInteger(res)

}
