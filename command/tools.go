package command

import (
	"errors"
)

func expectBulks(args Args) error {
	for _, v := range args {
		if v.Type() != "bulk" {
			return errors.New("Invalid arguments' format, expected bulk.")
		}
	}
	return nil
}
