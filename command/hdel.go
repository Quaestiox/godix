package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func HDel(args Args, config cfg.Config) resp.Val {
	count := 0
	n := len(args)
	if n < 2 {
		return resp.NewError("ERR", "wrong number of arguments for 'hdel' command.")
	}

	hash := args[0].Value().(string)

	HMapLock.Lock()
	for i := 1; i < n; i++ {
		v := args[i].Value().(string)
		if _, ok := HMap[hash][v]; ok {
			delete(HMap[hash], v)
			count++
		}

	}
	HMapLock.Unlock()

	return resp.NewInteger(count)

}
