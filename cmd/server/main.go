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

	listener, err := auth.ListenTLS(":8443", "certs/server-cert.pem", "certs/server-key.pem", "certs/ca.pem")
	if err != nil {
		log.Fatal("Failed to start TLS listener:", err)
	}
	defer listener.Close()

	log.Println("[server] Listening on :8443")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("[server] Accept error:", err)
			continue
		}
		log.Println("[server] Client connected")

		go vpn.StartForwarding(conn, ifce)	}
}
