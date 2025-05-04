package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
	"slices"
	"strconv"
)

func LREM(args Args, config cfg.Config) resp.Val {
	if len(args) < 3 {
		return resp.NewError("ERR", "wrong number of arguments for 'lrem' command.")
	}

	if res := expectBulks(args); res != nil {
		return resp.NewError("ERR", res.Error())
	}

	list := args[0].Value().(string)
	value := args[1].Value().(string)
	count, err := strconv.Atoi(value)
	if err != nil {
		return resp.NewError("ERR", "argument '"+value+"' is not an integer")
	}
	target := args[2].Value().(string)

	LMapLock.Lock()
	res := 0
	_, ok := LMap[list]
	if !ok {
		LMapLock.Unlock()
		return resp.NewInteger(0)
	}

	if count == 0 {
		for idx := 0; idx < len(LMap[list]); {
			el := LMap[list][idx]
			if el == target {
				LMap[list] = slices.Delete(LMap[list], idx, idx+1)
				res++
			} else {
				idx++
			}
		}
	} else if count > 0 {
		for idx := 0; idx < len(LMap[list]); {
			el := LMap[list][idx]
			if el == target {
				LMap[list] = slices.Delete(LMap[list], idx, idx+1)
				count--
				res++
			} else {
				idx++
			}
		}
	} else {
		for idx := len(LMap[list]) - 1; idx >= 0; {
			el := LMap[list][idx]
			if el == target && count < 0 {
				LMap[list] = slices.Delete(LMap[list], idx, idx+1)
				count++
				res++
			} else {
				idx--
			}

		}
	}

	LMapLock.Unlock()

	return resp.NewInteger(res)
}
