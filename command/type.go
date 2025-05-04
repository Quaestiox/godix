package command

import (
	"github.com/Quaestiox/godix/resp"
	"sync"
	"time"
)

type Args []resp.Val

type sv struct {
	value    string
	create_t time.Time
	expire   time.Time
}

func NewSV(v string) *sv {
	now := time.Now()
	return &sv{
		value:    v,
		create_t: now,
	}
}

func (sv *sv) Value() string {
	return sv.value
}

func (sv *sv) CreateTime() time.Time {
	return sv.create_t
}

func (sv *sv) Expire() time.Time {
	return sv.expire
}

func (sv *sv) setValue(s string) {
	sv.value = s
}

func (sv *sv) setExpire(duration time.Duration) {
	sv.expire = time.Now().Add(duration)
}

var WRCommand = []string{"SET", "DEL", "EXPIRE", "RENAME", "INCRBY", "INCR", "DECRBY", "DECR",
	"HSET", "HINCRBY", "HDEL",
	"LPUSH", "RPUSH", "LPOP", "RPOP", "LREM",
	"SADD", "SREM",
}

var Map = map[string]*sv{}
var MapLock = sync.RWMutex{}

var HMap = map[string]map[string]string{}
var HMapLock = sync.RWMutex{}

var LMap = map[string][]string{}
var LMapLock = sync.RWMutex{}

var SMap = map[string]map[string]bool{}
var SMapLock = sync.RWMutex{}
