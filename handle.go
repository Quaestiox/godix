package main

import (
	"fmt"
	"github.com/Quaestiox/godix/resp"
	"strings"
)

type Args []resp.Val

var Handlers = map[string]func(Args) resp.Val{
	"PING": ping,
}

func ping(args Args) resp.Val {
	return resp.NewString("PONG")
}

func HandleRequest(arr []resp.Val) (resp.Val, error) {
	if len(arr) == 0 {
		return nil, fmt.Errorf("Invalid request, expected array length > 0.\n")
	}

	cmd := arr[0]
	if cmd.Type() != "bulk" {
		return nil, fmt.Errorf("Invalid format, expected bulk.\n")
	}
	command := cmd.Value().(string)
	handler, ok := Handlers[strings.ToUpper(command)]
	if !ok {
		return resp.NewError("ERR", "unknown command:"+command), fmt.Errorf("Invalid command.\n")
	}
	args := arr[1:]
	res := handler(args)
	return res, nil
}
