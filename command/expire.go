package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
	"strconv"
	"time"
)

func Expire(args Args, config cfg.Config) resp.Val {
	if len(args) != 2 {
		return resp.NewError("ERR", "wrong number of arguments for 'expire' command.")
	}

	key := args[0].Value().(string)
	duration, err := strconv.Atoi(args[1].Value().(string))

	if err != nil {
		return resp.NewError("ERR", "wrong duration")
	}

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	MapLock.Lock()
	sv, ok := Map[key]
	if !ok {
		MapLock.Unlock()
		return resp.NewInteger(0)
	}
	d := time.Duration(duration) * time.Second
	sv.setExpire(d)
	MapLock.Unlock()

	ExpireRecordLock.Lock()
	ExpireRecord[key] = time.Now().Add(d)
	ExpireRecordLock.Unlock()

	return resp.NewInteger(1)

}
