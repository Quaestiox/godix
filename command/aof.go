package command

import (
	"github.com/Quaestiox/godix/cfg"
	"github.com/Quaestiox/godix/resp"
	"os"
)

func AOF(args Args, config cfg.Config) resp.Val {
	if len(args) < 1 {
		return resp.NewError("ERR", "wrong number of arguments for 'aof' command.")
	}

	cmd := args[0].Value().(string)
	switch cmd {
	case "clean":
		err := AOFClean(&config)
		if err != nil {
			return resp.NewError("ERR", err.Error())
		}
	default:
		return resp.NewError("ERR", "unknown command: "+cmd+".")
	}

	return resp.NewString("OK")
}

func AOFClean(c *cfg.Config) error {
	if !c.AofOn {
		os.Remove(c.AofPath)
		os.Create(c.AofPath)
		return nil
	}
	return c.Aof.Clean()

}
