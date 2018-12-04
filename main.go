package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/hypebeast/go-osc/osc"
)

var options struct {
	listen  string
	forward string
}

func exit(status int, t string, args ...interface{}) {
	if !strings.HasSuffix(t, "\n") {
		t += "\n"
	}

	if status != 0 {
		fmt.Fprintf(os.Stderr, t, args...)
	} else {
		fmt.Fprintf(os.Stdout, t, args...)
	}
	os.Exit(status)
}

func main() {
	flag.StringVar(&options.listen, "listen", "0.0.0.0:9220", "listen address for receiving OSC messages")
	flag.StringVar(&options.forward, "forward", "127.0.0.1:9225", "forward address for sending TCP data")
	flag.Parse()

	fmt.Printf("dialing %s\n", options.forward)
	conn, err := net.Dial("tcp", options.forward)
	if err != nil {
		exit(1, "unable to dial upstream server: %v", err)
	}
	defer conn.Close()
	fmt.Println("connected to upstream tcp handler")

	server := &osc.Server{Addr: options.listen}
	server.Handle("/example", func(msg *osc.Message) {
		fmt.Printf("received OSC message: %v\n", *msg)
		if len(msg.Arguments) == 0 {
			return
		}
		_, err := fmt.Fprintf(conn, "%d\n", msg.Arguments[0])
		if err != nil {
			exit(1, "ERROR: unable to forward message to upstream: %v\n", err)
		}
	})
	fmt.Printf("listening on %s\n", options.listen)
	server.ListenAndServe()
}
