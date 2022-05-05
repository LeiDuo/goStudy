package main

import (
	"fmt"
	"net"
)

func HandleErr(err error) {
	_ = fmt.Errorf(err.Error())
}

func main() {
	udp, err := net.ListenUDP("udp",
		&net.UDPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: 30000,
		})
	if err != nil {
		HandleErr(err)
		return
	}
	defer func(udp *net.UDPConn) {
		err := udp.Close()
		if err != nil {
			HandleErr(err)
			return
		}
	}(udp)
	for {
		var bytes [1024]byte
		n, u, err := udp.ReadFromUDP(bytes[:])
		if err != nil {
			HandleErr(err)
			return
		}
		fmt.Printf("data:%s,addr:%v,count:%v\n", string(bytes[:n]), u, n)
		_, err = udp.WriteToUDP(bytes[:n], u)
		if err != nil {
			HandleErr(err)
			continue
		}
	}

}
