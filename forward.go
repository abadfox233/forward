package main

import (
	"fmt"
	"io"
	"net"
	"time"

	"golang.org/x/sync/errgroup"
)

func Start() {

	startServer("0.0.0.0:8430", "127.0.0.1:8022")

}

func startServer(linstenAddr, forwardAddr string) {

	listener, err := net.Listen("tcp", linstenAddr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil && err.(net.Error).Timeout() {
			time.Sleep(1 * time.Second)
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("accept client from %s\n", conn.RemoteAddr().String())
		go handleConn(conn, forwardAddr)
	}

}

func handleConn(conn net.Conn, forwardAddr string) {

	defer conn.Close()
	forwordConn, err := net.DialTimeout("tcp", forwardAddr, 5*time.Second)
	fmt.Printf("dial forword addr %s\n", forwardAddr)
	if err != nil {
		fmt.Printf("dial forword addr %s error: %s\n", forwardAddr, err)
		return
	}

	defer forwordConn.Close()
	group := errgroup.Group{}

	group.Go(func() error {
		_, err := io.Copy(forwordConn, conn)
		if err != nil {
			fmt.Printf("copy conn to forwordConn error: %s\n", err)
		}
		return err
	})

	group.Go(func() error {
		_, err := io.Copy(conn, forwordConn)
		if err != nil {
			fmt.Printf("copy forwordConn to conn error: %s\n", err)
		}
		return err
	})

	if err := group.Wait(); err != nil {
		fmt.Printf("group wait error: %s\n", err)
	}
	fmt.Printf("conn %s close\n", conn.RemoteAddr().String())

}
