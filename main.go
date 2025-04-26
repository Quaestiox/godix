package main

import (
	"fmt"
	"github.com/Quaestiox/godix/persistence"
	"github.com/Quaestiox/godix/resp"
	"net"
)

var aofOn = true

func main() {
	var aof *persistence.AOF
	// godix server
	fmt.Println("Listening on port :6379")
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	// aof
	if aofOn {
		aof, err = persistence.NewAOF("godix.aof")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer aof.Close()
	}

	err = HandleAOF(aof)
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

	for {
		reader := resp.NewReader(conn)
		value, err := reader.Read()
		fmt.Println(value)
		if err != nil {
			fmt.Println(err)
			return
		}
		if value.Type() != "array" {
			fmt.Println("Invalid request, expected array")
			continue
		}
		res, err := HandleRequest(value, aof)
		if err != nil && err.Error() != "Invalid command.\n" {
			fmt.Println(err)
			continue
		}
		fmt.Println(res)
		writer := resp.NewWriter(conn)
		writer.Write(res)
	}
}
