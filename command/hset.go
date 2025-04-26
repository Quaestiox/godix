package command

import (
	"github.com/Quaestiox/godix/resp"
	"sync"
)

var HMap = map[string]map[string]string{}
var HMapLock = sync.RWMutex{}

func HSet(args Args) resp.Val {
	if len(args) != 3 {
		return resp.NewError("ERR", "wrong number of arguments for 'hset' command.")
	}

	if args[0].Type() != "bulk" || args[1].Type() != "bulk" || args[2].Type() != "bulk" {
		return resp.NewError("ERR", "Invalid arguments' format, expected bulk.")
	}

	hash := args[0].Value().(string)
	key := args[1].Value().(string)
	value := args[2].Value().(string)

	HMapLock.Lock()
	if _, ok := HMap[hash]; !ok {
		HMap[hash] = map[string]string{}
	}
	HMap[hash][key] = value
	HMapLock.Unlock()

	return resp.NewString("OK")
}
