package command

import (
	"github.com/Quaestiox/godix/resp"
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

func (sv *sv) setExpire(duration time.Duration) {

	sv.expire = time.Now().Add(duration)
}

var WRCommand = []string{"SET", "HSET", "RENAME", "DEL", "HDEL", "EXPIRE"}
