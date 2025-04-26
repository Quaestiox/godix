package cfg

import (
	"flag"
	"github.com/Quaestiox/godix/persistence"
)

type Config struct {
	AofOn   bool
	Aof     *persistence.AOF
	AofPath string
	Port    string
}

func (c *Config) Init() {
	c.AofPath = "./godix.aof"
	flag.BoolVar(&c.AofOn, "aof", true, "AOF data persistence")
	flag.StringVar(&c.Port, "port", "6379", "Port number on Godix server")
	flag.Parse()
}
