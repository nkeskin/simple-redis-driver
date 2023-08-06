package connection

import "net"

func RedisConnection(host string, port string) net.Conn {

	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		panic("Error while connecting to redis server")
	}
	return conn

}
