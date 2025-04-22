package resp

import "strconv"

type Bulk struct {
	typ  string
	bulk string
}

func NewBulk(bulk string) *Bulk {
	return &Bulk{
		typ:  "bulk",
		bulk: bulk,
	}
}

func (b *Bulk) Marshal() (bytes []byte) {
	bytes = append(bytes, BULK)
	bytes = append(bytes, strconv.Itoa(len(b.bulk))...)
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, b.bulk...)
	bytes = append(bytes, '\r', '\n')
	return
}

func (b *Bulk) Type() string {
	return b.typ
}

type Array struct {
	typ   string
	array []Val
}

func NewArray(values ...Val) *Array {

	return &Array{
		typ:   "array",
		array: values,
	}
}

func (arr *Array) Marshal() (bytes []byte) {
	length := len(arr.array)
	bytes = append(bytes, ARRAY)
	bytes = append(bytes, strconv.Itoa(length)...)
	bytes = append(bytes, '\r', '\n')

	for i := 0; i < length; i++ {
		bytes = append(bytes, arr.array[i].Marshal()...)
	}
	return
}

func (arr *Array) Type() string {
	return arr.typ
}

type Err struct {
	msg string
}

func (e *Err) Marshal() (bytes []byte) {
	bytes = append(bytes, ERROR)
	bytes = append(bytes, e.msg...)
	bytes = append(bytes, '\r', '\n')
	return
}

func (e *Err) Type() string {
	return ""
}

type Val interface {
	Marshal() []byte
	Type() string
}
