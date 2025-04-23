package command

import "github.com/Quaestiox/godix/resp"

func Ping(args Args) resp.Val {
	return resp.NewString("PONG")
}
