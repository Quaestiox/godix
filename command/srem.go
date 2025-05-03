package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func SRem(args Args, config cfg.Config) resp.Val {
	if len(args) < 2 {
		return resp.NewError("ERR", "wrong number of arguments for 'srem' command.")
	}

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	value := args[0].Value().(string)
	els := []string{}
	for i := 1; i < len(args); i++ {
		els = append(els, args[i].Value().(string))
	}

	count := 0
	SMapLock.Lock()
	_, ok := SMap[value]
	if !ok {
		SMap[value] = map[string]bool{}
	}

	for _, el := range els {
		if SMap[value][el] {
			delete(SMap[value], el)
			count++
		}
	}

	SMapLock.Unlock()

	return resp.NewInteger(count)
}
