package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func handleSignals() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP) // TODO Add other signals

	for {
		select {
		case s := <-sig:
			switch s {
			case syscall.SIGINT, syscall.SIGTERM:
				log.Println("Received kill signal, stopping...")
				stop()
			case syscall.SIGHUP:
				log.Println("Reloading config", configPath)
				if err := parseAll(&cfg); err != nil {
					panic(err)
				}
			}
		}
	}
}

func stop() {

	os.Exit(0)
}
