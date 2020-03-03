package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		// handle error
	}
	fmt.Fprintln(conn, "Hello, World!\n")
	msgRcv, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Msg Received:", msgRcv)

}
