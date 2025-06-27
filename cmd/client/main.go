package main

import (
	"log"

	"github.com/sahandset/burrow-vpn/internal/auth"
	"github.com/sahandset/burrow-vpn/internal/vpn"
)

func main() {
	ifce, err := vpn.CreateTUN("bvpn0")
	if err != nil {
		log.Fatal("Failed to create TUN:", err)
	}
	defer ifce.Close()

	conn, err := auth.DialTLS("localhost:8443", "certs/client-cert.pem", "certs/client-key.pem", "certs/ca.pem")
	if err != nil {
		log.Fatal("TLS connection failed:", err)
	}
	defer conn.Close()

	log.Println("[client] Connected to server")

	vpn.StartForwarding(conn, ifce)
}
