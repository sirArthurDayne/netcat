package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

const (
	DEFAULT_PORT = 8080
	DEFAULT_HOST = "localhost"
)

var (
	host = flag.String("h", DEFAULT_HOST, "host to listen.(default=localhost)")
	port = flag.Int("p", DEFAULT_PORT, "port to listen on. (default=8080)")
)

func CopyContent(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "io.Copy:%v\n", err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		panic(err)
	}
	fmt.Println("connected")
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()
	CopyContent(conn, os.Stdin)
	conn.Close()
	fmt.Println(<-done)
}
