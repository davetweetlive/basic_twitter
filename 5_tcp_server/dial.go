package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Println(err)
	}

	bs, _ := ioutil.ReadAll(conn)
	fmt.Println(string(bs))
}
