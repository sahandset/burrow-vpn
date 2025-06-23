package vpn

import (
	"fmt"
	"github.com/songgao/water"
)

func CreateTUN(name string) (*water.Interface, error) {
	config := water.Config{
		DeviceType: water.TUN,
	}
	config.Name = config.Name

	ifce, err := water.New(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create TUN device: %w", err)
	}

	fmt.Printf("TUN device %s created\n", ifce.Name())
	return ifce, nil
}
