package resp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Reader struct {
	reader *bufio.Reader
}

func NewReader(rd io.Reader) *Reader {
	return &Reader{reader: bufio.NewReader(rd)}
}

func (r *Reader) Read() (Val, error) {
	typ, err := r.reader.ReadByte()
	if err != nil {
		return &Err{msg: "Fail to read byte!"}, err
	}

	switch typ {
	case BULK:
		return r.readBulk()
	case ARRAY:
		return r.readArray()
	default:
		fmt.Printf("Unknown type: %v", string(typ))
		return &Err{msg: "Unknown type:" + string(typ)}, nil
	}

}

func (r *Reader) readArray() (*Array, error) {
	v := NewArray()
	length, _, err := r.readNum()
	if err != nil {
		return v, err
	}
	v.array = make([]Val, length)
	for i := 0; i < length; i++ {
		val, err := r.Read()
		if err != nil {
			return v, err
		}
		v.array[i] = val
	}
	return v, nil
}

func (r *Reader) readBulk() (*Bulk, error) {
	v := NewBulk("")
	num, _, err := r.readNum()
	if err != nil {
		return v, err
	}
	buf := make([]byte, num)
	r.reader.Read(buf)
	v.bulk = string(buf)
	r.reader.ReadLine()
	return v, nil
}

func (r *Reader) readNum() (res, n int, err error) {
	line, _, err := r.reader.ReadLine()
	n = len(line)
	if err != nil {
		return 0, 0, err
	}
	num, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, n, err
	}
	return int(num), n, nil
}
