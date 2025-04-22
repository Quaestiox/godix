package main

import (
	"fmt"
	"github.com/Quaestiox/godix/resp"
	"net"
)

func main() {
	fmt.Println("Listening on port :6379")
	l, err := net.Listen("tcp", ":6379")
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
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(value)
		writer := resp.NewWriter(conn)
		writer.Write(resp.NewString("OK"))
	}
}
