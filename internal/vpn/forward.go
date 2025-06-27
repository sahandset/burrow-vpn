package vpn

import (
	"io"
	"log"
	"net"
	"sync"

	"github.com/songgao/water"
)

func StartForwarding(conn net.Conn, ifce *water.Interface) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		buf := make([]byte, 1500)
		for {
			n, err := ifce.Read(buf)
			if err != nil {
				if err != io.EOF {
					log.Println("[forward] TUN read error:", err)
				}
				return
			}
			_, err = conn.Write(buf[:n])
			if err != nil {
				log.Println("[forward] Conn write error:", err)
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		buf := make([]byte, 1500)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				if err != io.EOF {
					log.Println("[forward] Conn read error:", err)
				}
				return
			}
			_, err = ifce.Write(buf[:n])
			if err != nil {
				log.Println("[forward] TUN write error:", err)
				return
			}
		}
	}()

	wg.Wait()
}
