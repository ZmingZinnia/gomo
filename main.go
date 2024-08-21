package main

import (
	"fmt"

	"github.com/alecthomas/kingpin/v2"
)

var (
	debug   = kingpin.Flag("debug", "Enable debug mode.").Bool()
	timeout = kingpin.Flag("timeout", "Timeout waiting for ping.").Default("5s").Envar("PING_TIMEOUT").Short('t').Duration()
	ip      = kingpin.Arg("ip", "IP address to ping.").Required().IP()
	count   = kingpin.Arg("count", "Number of packets to send").Int()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	fmt.Printf("debug: %t Would ping: %s with timeout %s and count %d\n", *debug, *ip, *timeout, *count)
}
