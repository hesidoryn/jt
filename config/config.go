package config

import "flag"

const (
	defaultPort     = "3333"
	defaultDumpFile = "dump.db"
)

type Config struct {
	Port     string
	Password string
	DumpFile string
}

func (c *Config) Parse() {
	flag.StringVar(&c.Port, "port", defaultPort, "port for bind server")
	flag.StringVar(&c.Password, "password", "", "password for server")
	flag.StringVar(&c.DumpFile, "dump", defaultDumpFile, "file with saved data")
	flag.Parse()
}
