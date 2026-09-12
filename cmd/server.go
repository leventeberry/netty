package main

import (
	"fmt"
	"net"
	"os"
)

func server(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Server started on [ %s ]\n", addr)
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return err
		}
		go func(conn net.Conn) {
			for {
				defer conn.Close()
				buffer := make([]byte, 4096)
				n, err := conn.Read(buffer)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Println(string(buffer[:n]))
			}
		} (conn)
	}

	return nil
}