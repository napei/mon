package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jacobsa/go-serial/serial"
	"github.com/jessevdk/go-flags"
	s "go.bug.st/serial.v1"
)

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func getFirstPort() (*string, error) {
	ports, err := s.GetPortsList()
	if err != nil {
		return nil, err
	}

	if len(ports) == 0 {
		return nil, nil
	}

	return &ports[0], nil
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

type Opts struct {
	Port     *string `short:"p" long:"port" description:"serial port to use" default-mask:"first available port"`
	Baud     *uint   `short:"b" long:"baud" description:"baud rate to use" default:"115200"`
	DataBits *uint   `short:"d" long:"databits" description:"data bits to use" default:"8"`
	StopBits *uint   `short:"s" long:"stopbits" description:"stop bits to use" default:"1"`
}

func main() {
	var opts Opts
	if _, err := flags.Parse(&opts); err != nil {
		log.Fatal(err)
	}

	var err error

	if opts.Port == nil {
		p, err := getFirstPort()
		if err != nil {
			log.Fatalf("error getting port: %v", err)
			return
		}

		if p == nil {
			log.Fatal("No available serial ports were detected")
			return
		}

		opts.Port = p
	}


	options := serial.OpenOptions{
		PortName:        *opts.Port,
		BaudRate:        *opts.Baud,
		DataBits:        *opts.DataBits,
		StopBits:        *opts.StopBits,
		MinimumReadSize: 1,
	}
	conn, err := serial.Open(options)
	if err != nil {
		log.Fatalf("error opening serial port: %v", err)
		return
	}
	defer conn.Close()

	fmt.Println("Reading serial port", *opts.Port)

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
