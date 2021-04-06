package main

import (
	"crypto/tls"
	"log"
	"time"

	"github.com/kevacoin-project/go-electrum/v2"
)

func main() {
	server := electrum.NewServer()
	conf := &tls.Config{}
	if err := server.ConnectSSL("ec0.kevacoin.org:50002", conf); err != nil {
		log.Fatal(err)
	}

	serverVer, protocolVer, err := server.ServerVersion()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server version: %s [Protocol %s]", serverVer, protocolVer)

	go func() {
		for {
			if err := server.Ping(); err != nil {
				log.Fatal(err)
			}
			time.Sleep(60 * time.Second)
		}
	}()
}
