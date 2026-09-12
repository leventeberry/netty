package main

import (
	"log"
)

func main() {

	if err := server(":8082"); err != nil {
		log.Fatal(err)
	}

}
