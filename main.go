package main

import (
	"net"
	"net/netip"
	"os"
	"strings"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp4", ":1195")
	udpAddrRemote, err := net.ResolveUDPAddr("udp4", "60.204.171.251:1195")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	listen, err := net.ListenUDP("udp4", udpAddr)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println("udp server started")

	inter, err := net.InterfaceByName("en0")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	cid := strings.TrimSpace(inter.HardwareAddr.String())

	listen.WriteToUDP([]byte("a,"+cid), udpAddrRemote)

	for {
		var buf [128]byte
		len, addr, err := listen.ReadFromUDPAddrPort(buf[:])

		if err != nil {
			println(err.Error())
			continue
		}

		ip := addr.Addr().String()
		port := addr.Port()
		msg := string(buf[:len])

		println(ip, ":", port, " -> ", msg)

		if msg[0:2] == "cc" {
			handleCC(listen, msg[3:])
		} else if msg[0:3] == "ccc" {
			handleCCC(listen, addr)
		} else if msg[0:4] == "cccc" {
			handleCCC(listen, addr)
		}
	}
}
func handleCC(listen *net.UDPConn, address string) {
	udpAddr, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	listen.WriteToUDP([]byte("ccc"), udpAddr)
}
func handleCCC(listen *net.UDPConn, addr netip.AddrPort) {
	listen.WriteToUDPAddrPort([]byte("cccc"), addr)
}
func handleCCCC(listen *net.UDPConn, addr netip.AddrPort) {
	listen.WriteToUDPAddrPort([]byte("ccccc"), addr)
}
