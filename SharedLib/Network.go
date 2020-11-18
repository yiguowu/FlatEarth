package SharedLib

import (
	"github.com/chyeh/pubip"
	"net"
)

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		PanicOnError(err, WARNING)
	}
	defer func() {
		err = conn.Close()
		PanicOnError(err, WARNING)
	}()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func GetPublicIP() net.IP {
	ip, err := pubip.Get()
	if err != nil {
		PanicOnError(err, WARNING)
	}
	return ip
}
