package main

import (
	"fmt"
	"net"
)

func HandleErr(err error) {
	_ = fmt.Errorf(err.Error())
}

func main() {
	udp, err := net.DialUDP("udp", nil,
		&net.UDPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: 30000,
		},
	)
	if err != nil {
		HandleErr(err)
	}
	defer func(udp *net.UDPConn) {
		err := udp.Close()
		if err != nil {
			HandleErr(err)
			return
		}
	}(udp)
	write, err := udp.Write([]byte("测试我要写的道喜"))
	if err != nil {
		HandleErr(err)
	}
	fmt.Println("写入", write, "个字节")
	var bytes [1024]byte
	n, addr, err := udp.ReadFromUDP(bytes[:])
	if err != nil {
		HandleErr(err)
	}
	fmt.Printf("从%v读取%v字节:%v", n, addr, string(bytes[:n]))
}
