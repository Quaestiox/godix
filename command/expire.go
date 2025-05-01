package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
	"strconv"
	"sync"
	"time"
)

var ExpireRecord = map[string]time.Time{}
var ExpireRecordLock = sync.RWMutex{}

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
	sv := Map[key]
	d := time.Duration(duration) * time.Second
	sv.setExpire(d)
	MapLock.Unlock()

	ExpireRecordLock.Lock()
	ExpireRecord[key] = time.Now().Add(d)
	ExpireRecordLock.Unlock()

	return resp.NewString("OK")

}
