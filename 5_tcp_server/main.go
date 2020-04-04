package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	ls, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println(err)
	}

	defer ls.Close()

	for {
		conn, err := ls.Accept()
		if err != nil {
			panic(err)
		}

		io.WriteString(conn, fmt.Sprint("Hello World\n", time.Now(), "\n"))
		conn.Close()
	}
}
