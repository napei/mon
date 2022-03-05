# Mon

A tiny cross-platform serial port monitor written in Go, for making life simpler when working with embedded systems.

No more PuTTY, no more screen, just `mon`

## Installation

Install the package by running `go install github.com/napei/mon`

Alternatively, clone this repository and run `go build` to build for your platform manually.

## Usage

`$ mon [port] [baud]`

- If no port is specified, `mon` will try to open the first available serial port.
  - Specify the port in the form for your operating system (`COM3`, `/dev/ttyUSB0`, etc.)
- If no baud is specified, `mon` will use a default baud rate of `115200`
