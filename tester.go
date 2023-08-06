package main

import (
	"bufio"
	"fmt"
	"simple-redis-driver/pkg/client"
	"simple-redis-driver/pkg/connection"
)

func main() {

	conn := connection.RedisConnection("localhost", "6379")
	fmt.Println(conn)
	fmt.Fprintf(conn, "PING\r\n")
	readString, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return
	}
	fmt.Fprintf(conn, "SET mykey deneme2\r\n")
	readString2, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return
	}
	fmt.Println(readString2 + readString)
	client.Demo2()
}
