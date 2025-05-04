package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func HSet(args Args, config cfg.Config) resp.Val {
	if len(args) < 3 || len(args)%2 != 1 {
		return resp.NewError("ERR", "wrong number of arguments for 'hset' command.")
	}

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	kv := []string{}
	count := 0
	hash := args[0].Value().(string)
	for idx, v := range args {
		if idx > 0 {
			kv = append(kv, v.Value().(string))
		}
	}

	HMapLock.Lock()
	if _, ok := HMap[hash]; !ok {
		HMap[hash] = map[string]string{}
	}
	for i := 0; i < len(args)-1; {
		if _, ok := HMap[hash][kv[i]]; !ok {
			count++
		}
		HMap[hash][kv[i]] = kv[i+1]
		i += 2
	}
	HMapLock.Unlock()

	return resp.NewInteger(count)
}
