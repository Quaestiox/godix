package command

import "github.com/Quaestiox/godix/resp"

func Get(args Args) resp.Val {
	if len(args) != 1 {
		return resp.NewError("ERR", "wrong number of arguments for 'set' command.")
	}
	key := args[0]
	MapLock.RLock()
	value, ok := Map[key.Value().(string)]
	MapLock.RUnlock()
	if !ok {
		return resp.NewNullBulk()
	}
	return resp.NewBulk(value)
}
