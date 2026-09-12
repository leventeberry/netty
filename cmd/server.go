package main

import (
	"fmt"
	"net"
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
			request := make([]byte, 0)
			for {
				defer conn.Close()
				buffer := make([]byte, 8)
				n, err := conn.Read(buffer)
				if err != nil {
					fmt.Println(err)
				}
				request = append(request, buffer[:n]...)
				requestLine := ParseRequestLine(string(request))
				if requestLine != nil {
					fmt.Println(requestLine)
					break
				}
			}
		} (conn)
	}
}