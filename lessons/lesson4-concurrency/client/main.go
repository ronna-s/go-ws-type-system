package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"time"
)

var addr = flag.String("addr", ":9090", "address to listen")

func writeAndRead(r *bufio.Reader, writer io.Writer, s string) {
	if _, err := writer.Write([]byte(s + "\n")); err != nil {
		log.Fatal(err)
	}
	s2, err := r.ReadString('\n')
	log.Print(s2)
	if err != nil {
		log.Fatal(err)
	}
	if s != s2 {
		log.Fatalf("expected %s, got %s", s, s2)
	}
}

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	r := bufio.NewReader(conn)
	writeAndRead(r, conn, "hello")
	writeAndRead(r, conn, "world")
	time.Sleep(1 * time.Minute)
}
