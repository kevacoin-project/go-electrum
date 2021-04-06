package electrum

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"testing"
)

func TestGetKeyValues(t *testing.T) {
	server := NewServer()
	conf := &tls.Config{}
	if err := server.ConnectSSL("ec0.kevacoin.org:50002", conf); err != nil {
		log.Fatal(err)
	}

	serverVer, protocolVer, err := server.ServerVersion()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server version: %s [Protocol %s]", serverVer, protocolVer)

	keyValues, err := server.GetKeyValues("4dabe718159c7eb854c95342b99bb216977f7dd4e9fc0a16f7e97c9eafb9235e", -1)
	if err != nil {
		log.Fatal(err)
	}

	jsonKeyValues, err := json.Marshal(keyValues)
	log.Printf("get_key_values result is: %s", string(jsonKeyValues))
}
