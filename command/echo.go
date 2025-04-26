package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
)

func Echo(args Args, config cfg.Config) resp.Val {
	s := ""
	for _, v := range args {
		v1 := v.Value().(string)
		s += v1
	}
	return resp.NewString(s)
}
