package vpn

import (
	"log"

	"github.com/songgao/water"
)

func CreateTUN(name string) (*water.Interface, error) {
	config := water.Config{
		DeviceType: water.TUN,
	}
	config.Name = name

	ifce, err := water.New(config)
	if err != nil {
		return nil, err
	}

	log.Println("[vpn] TUN interface created:", ifce.Name())
	return ifce, nil
}
