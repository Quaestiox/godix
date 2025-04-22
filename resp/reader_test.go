package resp

import (
	"fmt"
	"strings"
	"testing"
)

func Test_ReadLine(t *testing.T) {
	strR := strings.NewReader("hello\r\n")
	reader := NewReader(strR)
	res, _, _ := reader.reader.ReadLine()
	if string(res) != "hello" {
		t.Fatal("wrong")
	}
}

func Test_ReadNum(t *testing.T) {
	strR := strings.NewReader("123\r\n")
	reader := NewReader(strR)
	res, len, _ := reader.readNum()
	fmt.Println(res, " ", len)
	if res != 123 {
		t.Fatal("wrong")
	}
	if len != 3 {
		t.Fatal("wrong")
	}
}

func Test_ReadBulk(t *testing.T) {
	strR := strings.NewReader("5\r\nhello\r\n")
	reader := NewReader(strR)
	bulk, _ := reader.readBulk()
	if bulk.typ != "bulk" || bulk.bulk != "hello" {
		t.Fatal("wrong")
	}
}

func Test_ReadArray(t *testing.T) {
	strR := strings.NewReader("2\r\n$5\r\nhello\r\n$3\r\nyou\r\n")
	reader := NewReader(strR)
	array, _ := reader.readArray()
	if array.Type() != "array" || len(array.array) != 2 || array.array[0].Type() != "bulk" ||
		array.array[1].Type() != "bulk" {
		t.Fatal("wrong")
	}
}

func Test_Read(t *testing.T) {
	strR := strings.NewReader("*2\r\n$5\r\nhello\r\n$3\r\nyou\r\n")
	reader := NewReader(strR)
	res, _ := reader.Read()
	if res.Type() != "array" {
		t.Fatal("wrong")
	}
}
