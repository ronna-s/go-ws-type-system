package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	const addr = ":9090"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		time.Sleep(10 * time.Second)
		log.Println("closing the listener...")
		l.Close()
	}()
	log.Print(Serve(l, Handle))
}

type Handler func(writer io.ReadWriter)

func Handle(c io.ReadWriter) {
	s := bufio.NewScanner(c)
	// reads a single line from the connection and writes it back
	for s.Scan() {
		b := s.Bytes()
		log.Println("received:", string(b))
		_, err := c.Write(append(b, []byte("\n")...))
		if err != nil {
			log.Fatal("failed to write to the connection")
		}
	}
}

func Serve(l net.Listener, h Handler) error {
	var wg sync.WaitGroup
	var conn net.Conn
	var err error
	for {
		conn, err = l.Accept()
		if err != nil {
			break
		}
		wg.Add(1)
		go func(c net.Conn) {
			defer wg.Done()
			h(c)
		}(conn)
	}
	wg.Wait()
	return err
}
