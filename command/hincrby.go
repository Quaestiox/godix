package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
	"strconv"
)

func HIncrBy(args Args, config cfg.Config) resp.Val {
	if len(args) != 3 {
		return resp.NewError("ERR", "wrong number of arguments for 'hincrby' command.")
	}
	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	hash := args[0].Value().(string)
	field := args[1].Value().(string)
	value := args[2].Value().(string)
	val, err := strconv.Atoi(value)
	if err != nil {
		return resp.NewError("ERR", "argument '"+value+"' is not an integer")
	}

	HMapLock.Lock()
	_, ok := HMap[hash]
	if !ok {
		HMap[hash] = map[string]string{}
	}
	mp, ok := HMap[hash][field]
	if !ok {
		HMap[hash][field] = "0"
		mp = HMap[hash][field]
	}
	v, err := strconv.Atoi(mp)
	if err != nil {
		HMapLock.Unlock()
		return resp.NewError("ERR", "'"+mp+"' is not an integer")
	}
	newv := v + val
	HMap[hash][field] = strconv.Itoa(newv)
	HMapLock.Unlock()

	return resp.NewInteger(newv)

}
