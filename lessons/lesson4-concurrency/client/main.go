package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"time"
)

var addr = flag.String("addr", ":9090", "address to listen")

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if _, err = conn.Write([]byte("hello\n")); err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(conn)
	s, err := r.ReadString('\n')
	log.Print(s)
	if err != nil {
		log.Fatal(err)
	}
	if _, err = conn.Write([]byte("world\n")); err != nil {
		log.Fatal(err)
	}
	s, err = r.ReadString('\n')
	log.Print(s)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Minute)
}
