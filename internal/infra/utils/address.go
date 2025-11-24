package utils

import (
	"net"
)

func GetLocalAddress() (string, error) {
	dial, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer dial.Close()

	localAddr := dial.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}
