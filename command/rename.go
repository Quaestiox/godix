package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func Rename(args Args, config cfg.Config) resp.Val {

	if len(args) != 2 {
		return resp.NewError("ERR", "wrong number of arguments for 'rename' command.")
	}
	key := args[0].Value().(string)
	newkey := args[1].Value().(string)

	MapLock.Lock()
	sv, ok := Map[key]
	if !ok {
		MapLock.Unlock()
		return resp.NewNullBulk()
	}
	delete(Map, key)
	Map[newkey] = sv
	MapLock.Unlock()

	if !ok {
		return resp.NewNullBulk()
	}

	ExpireRecordLock.Lock()
	exp, ok := ExpireRecord[key]
	delete(ExpireRecord, key)
	if ok {
		ExpireRecord[newkey] = exp
	}
	ExpireRecordLock.Unlock()

	return resp.NewString("OK")
}
