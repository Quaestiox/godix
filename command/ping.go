package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func Ping(args Args, config cfg.Config) resp.Val {
	return resp.NewString("PONG")
}
