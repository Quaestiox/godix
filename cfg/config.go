package cfg

import (
	"flag"
	"github.com/Quaestiox/godix/persistence"
)

type Config struct {
	AofOn   bool
	Aof     *persistence.AOF
	AofPath string
}

func (c *Config) Init() {
	c.AofPath = "./godix.aof"
	flag.BoolVar(&c.AofOn, "aof", true, "AOF data persistence")
	flag.Parse()
}
