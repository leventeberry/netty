package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	// Bind a tcp listener to a network process	
	l, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go func(conn net.Conn) {
			for {	
				defer conn.Close()
				buffer := make([]byte, 4096)
				n, err := conn.Read(buffer)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(string(buffer[:n]))
			}
			
		} (conn)
	}
}
