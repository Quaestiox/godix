package persistence

import (
	"bufio"
	"fmt"
	"github.com/Quaestiox/godix/resp"
	"io"
	"os"
	"sync"
	"time"
)

type AOF struct {
	file   *os.File
	reader *bufio.Reader
	mu     sync.Mutex
}

func NewAOF(path string) (*AOF, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	aof := &AOF{
		file:   f,
		reader: bufio.NewReader(f),
	}

	go func() {
		for {
			aof.mu.Lock()
			err := aof.file.Sync()
			if err != nil {
				fmt.Println(err)
			}
			aof.mu.Unlock()
			time.Sleep(time.Second)
		}
	}()
	return aof, nil
}

func (aof *AOF) Close() error {
	aof.mu.Lock()
	defer aof.mu.Unlock()
	return aof.file.Close()
}

func (aof *AOF) Write(v resp.Val) error {
	aof.mu.Lock()
	defer aof.mu.Unlock()
	bytes := v.Marshal()
	_, err := aof.file.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func (aof *AOF) Read(callback func(value resp.Val)) error {
	aof.mu.Lock()
	defer aof.mu.Unlock()

	reader := resp.NewReader(aof.file)
	for {
		value, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		callback(value)

	}
	return nil
}
