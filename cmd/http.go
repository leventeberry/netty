package main

import (
	"fmt"
	"strings"
)
type RequestLine struct {
	Method string
	Uri string
	Version string
}

func ParseRequestLine(request string) *RequestLine {
	fmt.Println("Processing request line...")

	if request == "" {
		return nil
	}

	lines := strings.Split(request, "\r\n")
	for i, line := range lines {
		if i == 0 {
			parts := strings.Split(line, " ")
			if len(parts) == 3 && parts[2] == "HTTP/1.1" {
				return &RequestLine{
					Method: parts[0],
					Uri: parts[1],
					Version: parts[2],
				}
			} else {
				return nil
			}
		}
	}
	return nil
}