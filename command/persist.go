package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func Persist(args Args, config cfg.Config) resp.Val {
	if len(args) != 1 {
		return resp.NewError("ERR", "wrong number of arguments for 'persist' command.")
	}

	key := args[0].Value().(string)

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	MapLock.Lock()
	v, ok := Map[key]
	if !ok {
		MapLock.Unlock()
		return resp.NewInteger(0)
	}
	v.persist()
	MapLock.Unlock()

	ExpireRecordLock.Lock()
	_, ok = ExpireRecord[key]
	if !ok {
		ExpireRecordLock.Unlock()
		return resp.NewInteger(0)
	}
	delete(ExpireRecord, key)
	ExpireRecordLock.Unlock()

	return resp.NewInteger(1)

}
