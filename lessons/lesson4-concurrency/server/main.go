package main

import (
	"bufio"
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

type Handler func(net.Conn)

func Handle(c net.Conn) {
	s := bufio.NewScanner(c)
	for s.Scan() {
		c.Write(s.Bytes())
		c.Write([]byte("\n"))
	}
	log.Println("client closed the connection")
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
