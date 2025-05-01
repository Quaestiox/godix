package main

import (
	"fmt"
	"github.com/Quaestiox/godix/cfg"

	"github.com/Quaestiox/godix/persistence"
	"github.com/Quaestiox/godix/resp"
	"net"
)

var config cfg.Config

func main() {
	Banner()

	config.Init()

	// godix server
	fmt.Println("Listening on port :" + config.Port)
	l, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		fmt.Println(err)
		return
	}

	// aof
	if config.AofOn {
		config.Aof, err = persistence.NewAOF(config.AofPath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer config.Aof.Close()
	}

	err = HandleAOF(config.Aof)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	go HandleExpire()

	for {
		reader := resp.NewReader(conn)
		value, err := reader.Read()
		if err != nil {
			fmt.Println(err)
			return
		}
		if value.Type() != "array" {
			fmt.Println("Invalid request, expected array")
			continue
		}
		res, err := HandleRequest(value, config.Aof)
		if err != nil && err.Error() != "Invalid command.\n" {
			fmt.Println(err)
			continue
		}
		writer := resp.NewWriter(conn)
		writer.Write(res)
	}
}
