package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jacobsa/go-serial/serial"
	s "go.bug.st/serial.v1"
)

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func getFirstPort() (string, error) {
	ports, err := s.GetPortsList()
	if err != nil {
		return "", err
	}

	if len(ports) == 0 {
		return "", errors.New("no ports found")
	}

	return ports[0], nil
}

func getBaudOrDefault(input string) uint {
	if input == "" {
		return 115200
	}

	baud, err := strconv.Atoi(input)
	if err != nil {
		return 115200
	}

	if !serial.IsStandardBaudRate(uint(baud)) {
		fmt.Println("invalid baud rate, using 115200 instead")
		return 115200
	}

	return uint(baud)
}

func main() {
	args := os.Args[1:]
	var err error
	var port string
	var baud uint

	switch len(args) {
	case 0:
		port, err = getFirstPort()
		if err != nil {
			log.Fatalf("error getting port: %v", err)
			return
		}

		baud = 115200
	case 1:
		port = args[0]
		baud = 115200
	case 2:
		port = args[0]
		baud = getBaudOrDefault(args[1])
	default:
		log.Fatalln("incorrect number of args")
		return
	}

	options := serial.OpenOptions{
		PortName:        port,
		BaudRate:        baud,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}
	conn, err := serial.Open(options)
	if err != nil {
		log.Fatalf("error opening serial port: %v", err)
		return
	}
	defer conn.Close()

	var buf []byte

	for {
		buf = make([]byte, 1)
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatalf("error reading serial port: %v", err)
			return
		}

		fmt.Printf("%s", buf[:n])
	}
}
