package main

import (
	"bytes"
	"net"
	"strconv"
)

func generateMagicPacket(macAddress string) ([]byte, error) {
	macBytes, err := net.ParseMAC(macAddress)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	// Write 6 bytes of 0xFF
	for i := 0; i < 6; i++ {
		buf.WriteByte(0xFF)
	}

	// Write 16 copies of the MAC address
	for i := 0; i < 16; i++ {
		buf.Write(macBytes)
	}

	return buf.Bytes(), nil
}

func sendMagicPacket(macAddress string, broadcastAddress string, port int) error {
	packet, err := generateMagicPacket(macAddress)
	if err != nil {
		return err
	}

	addr, err := net.ResolveUDPAddr("udp", broadcastAddress+":"+strconv.Itoa(port))

	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(packet)
	if err != nil {
		return err
	}

	return nil
}

func Wake(macAddress string, broadcastAddress string, port int) error {
	return sendMagicPacket(macAddress, broadcastAddress, port)
}

func WakeDefault(macAddress string) error {
	return sendMagicPacket(macAddress, "255.255.255.255", 9)
}
