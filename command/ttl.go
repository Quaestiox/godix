package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
	"time"
)

func TTL(args Args, config cfg.Config) resp.Val {
	if len(args) != 1 {
		return resp.NewError("ERR", "wrong number of arguments for 'ttl' command.")
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
	res := int(exp.Sub(time.Now()).Seconds())
	ExpireRecordLock.RUnlock()

	return resp.NewInteger(res)

}
