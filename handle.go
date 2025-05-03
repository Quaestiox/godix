package main

import (
	"fmt"
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/command"
	"github.com/Quaestiox/godix/persistence"
	"github.com/Quaestiox/godix/resp"
	"slices"
	"strings"
	"time"
)

var Handlers = map[string]func(command.Args, cfg.Config) resp.Val{
	"PING":     command.Ping,
	"SET":      command.Set,
	"GET":      command.Get,
	"DEL":      command.Del,
	"EXISTS":   command.Exists,
	"RENAME":   command.Rename,
	"HSET":     command.HSet,
	"HGET":     command.HGet,
	"HDEL":     command.HDel,
	"HEXISTS":  command.HExists,
	"AOF":      command.AOF,
	"ABOUT":    command.About,
	"ECHO":     command.Echo,
	"EXPIRE":   command.Expire,
	"TTL":      command.TTL,
	"PTTL":     command.PTTL,
	"INCR":     command.Incr,
	"INCRBY":   command.IncrBy,
	"DECR":     command.Decr,
	"DECRBY":   command.DecrBy,
	"LINDEX":   command.LIndex,
	"LPUSH":    command.LPush,
	"LPOP":     command.LPop,
	"LLEN":     command.LLen,
	"SADD":     command.SAdd,
	"SMEMBERS": command.SMemebers,
	"SREM":     command.SRem,
	"SCARD":    command.SCard,
}

func HandleRequest(value resp.Val, aof *persistence.AOF) (resp.Val, error) {
	arr := value.Value().([]resp.Val)
	if len(arr) == 0 {
		return nil, fmt.Errorf("Invalid request, expected array length > 0.\n")
	}

	cmd := arr[0]
	if cmd.Type() != "bulk" {
		return nil, fmt.Errorf("Invalid format, expected bulk.\n")
	}
	cmdS := cmd.Value().(string)
	handler, ok := Handlers[strings.ToUpper(cmdS)]
	if !ok {
		return resp.NewError("ERR", "unknown command: "+cmdS), fmt.Errorf("Invalid command.\n")
	}
	args := arr[1:]
	res := handler(args, config)
	if config.AofOn && slices.Contains(command.WRCommand, strings.ToUpper(cmdS)) {
		err := aof.Write(value)
		if err != nil {
			return resp.NewError("ERR", "internal error"), err
		}
	}
	return res, nil
}

func HandleAOF(aof *persistence.AOF) error {
	if aof == nil {
		return nil
	}
	return aof.Read(func(value resp.Val) {
		arr := value.Value().([]resp.Val)
		cmd := arr[0]
		cmdS := cmd.Value().(string)
		handler, _ := Handlers[strings.ToUpper(cmdS)]
		args := arr[1:]
		handler(args, config)
	})
}

func HandleExpire() {
	ticker := time.NewTicker(config.ExpireTick)
	defer ticker.Stop()
	for range ticker.C {
		command.MapLock.Lock()
		now := time.Now()
		for k, sv := range command.Map {
			if !sv.Expire().IsZero() && now.After(sv.Expire()) {
				delete(command.Map, k)
			}
		}
		command.MapLock.Unlock()
	}
}
