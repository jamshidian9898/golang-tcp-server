package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8880")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		HandelConnection(conn)
		conn.Close()
	}

}

func HandelConnection(conn net.Conn) {
	incomingMessage := make([]byte, 1024)
	_, err := conn.Read(incomingMessage)
	if err != nil {
		panic(err)
	}

	httpResponse := `HTTP/1.1 200 OK
Date: Mon, 27 Jul 2009 12:28:53 GMT
Server: Apache/2.2.14 (Win32)
Last-Modified: Wed, 22 Jul 2009 19:15:56 GMT
Content-Type: text/html
Connection: Closed

<html>
<body>
<h1>Hello, World!</h1>
</body>
</html>`

	_, err = conn.Write([]byte(httpResponse))
	if err != nil {
		panic(err)
	}

}
