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
		if value.Type() != "array" {
			fmt.Println("Invalid request, expected array")
			continue
		}
		arr := value.Value().([]resp.Val)
		res, err := HandleRequest(arr)
		if err != nil && err.Error() != "Invalid command.\n" {
			fmt.Println(err)
			continue
		}
		fmt.Println(res)
		writer := resp.NewWriter(conn)
		writer.Write(res)
	}
}
