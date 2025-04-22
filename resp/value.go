package resp

import "strconv"

type Val interface {
	Marshal() []byte
	Type() string
	Value() any
}

type String struct {
	typ string
	str string
}

func NewString(str string) *String {
	return &String{
		typ: "string",
		str: str,
	}
}

func (s *String) Marshal() (bytes []byte) {
	bytes = append(bytes, STRING)
	bytes = append(bytes, s.str...)
	bytes = append(bytes, '\r', '\n')
	return
}

func (s *String) Type() string {
	return s.typ
}

func (s *String) Value() any {
	return s.str
}

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

func (b *Bulk) Value() any {
	return b.bulk
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

func (arr *Array) Value() any {
	return arr.array
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
	typ    string
	prefix string
	msg    string
}

func NewError(prefix, msg string) *Err {
	return &Err{
		typ:    "error",
		prefix: prefix,
		msg:    msg,
	}
}

func (e *Err) Marshal() (bytes []byte) {
	bytes = append(bytes, ERROR)
	bytes = append(bytes, e.prefix...)
	bytes = append(bytes, ' ')
	bytes = append(bytes, e.msg...)
	bytes = append(bytes, '\r', '\n')
	return
}

func (e *Err) Type() string {
	return e.typ
}

func (e *Err) Value() any {
	return e.msg
}

type Null struct {
	typ string
}

func NewNull() *Null {
	return &Null{
		typ: "null",
	}
}

func (n *Null) Type() string {
	return n.typ
}

func (n *Null) Marshal() (bytes []byte) {
	bytes = append(bytes, []byte("_/r/n")...)
	return
}

func (n *Null) Value() any {
	return nil
}
