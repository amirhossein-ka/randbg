package main

import (
	"github.com/amirhossein-ka/randbg/config"
	"github.com/amirhossein-ka/randbg/daemon"
	"log"
	"os"
)

var cfg = &config.Config{}

const CfgPath = "./cfg.json"

func init() {
	log.SetOutput(os.Stdout)
	if err := config.Parse(CfgPath, cfg); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	log.Printf("%#v\n", cfg)
	dmn,err := daemon.New(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	dmn.Run()
}
