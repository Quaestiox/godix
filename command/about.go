package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func About(args Args, config cfg.Config) resp.Val {
	return resp.NewBulk("Godix(github.com/Quaestiox/godix)")
}
