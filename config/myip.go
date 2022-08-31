package config

import (
	"log"
	"net"
)

func MyIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Println("Error al obtener IP, servidor no iniciado")
		log.Panic(err)
	}

	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
