package persistence

import "testing"

func TestAOF_Path(t *testing.T) {
	p := "./test.aof"
	aof, err := NewAOF(p)
	if err != nil {
		t.Fatal("wrong")
	}
	path := aof.Path()
	if path != p {
		t.Fatal("wrong")
	}

}
