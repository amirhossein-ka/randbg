package main

import (
	"github.com/amirhossein-ka/randbg/args"
	"github.com/amirhossein-ka/randbg/config"
	"log"
	"os"
)

var (
	cfg = &config.Config{}
	arg = &args.Args{}
)

func init() {
	log.SetOutput(os.Stdout)

	if err := arg.Init(os.Args); err != nil {
		log.Fatalln(err)
	}
	if err := config.Parse(arg.DaemonArgs.ConfigPath, cfg); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	if err := arg.RunApp(cfg); err != nil {
		log.Fatalln(err)
	}
}
