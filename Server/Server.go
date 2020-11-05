package main

import (
	"FlatEarth/SharedLib"
	"log"
	"net"
)

func GetTCPListener(host, port string) *net.TCPListener {
	service := host + ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	SharedLib.PanicOnError(err, SharedLib.FATAL)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	SharedLib.PanicOnError(err, SharedLib.FATAL)
	return listener
}

func main() {
	params := SharedLib.ParseParameter()
	listener := GetTCPListener(*params.Host, *params.Port)
	log.Printf("Listening on host %s port %s", *params.Host, *params.Port)
	for {
		conn, err := listener.Accept()
		SharedLib.PanicOnError(err, SharedLib.WARNING)
		err = conn.Close()
		SharedLib.PanicOnError(err, SharedLib.WARNING)
	}
}
