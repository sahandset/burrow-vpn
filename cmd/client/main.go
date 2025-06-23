package main

import (
	"github.com/sahandset/burrow-vpn/internal/vpn"
	"log"
)

func main() {
	ifce, err := vpn.CreateTUN("bvpn1")
	if err != nil {
		log.Fatal(err)
	}
	defer ifce.Close()
}
