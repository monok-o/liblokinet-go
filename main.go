package main

import (
	"log"
	"lokitest/lokinet"
)

func main() {
	isReady := lokinet.Start()

	lokinet.Connect("exit.loki")

	if isReady {
		log.Println("our address:", lokinet.Addr())

	}
}
