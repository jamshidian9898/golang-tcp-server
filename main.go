package main

import (
	"errors"
	"io"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8880")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for i := 0; i < 4; i++ {
		go ListenForConnection(listener)
	}
	c := make(chan struct{})
	<-c
}

func ListenForConnection(listener net.Listener) error {
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			HandleConnection(conn)
			conn.Close()
		}()
	}
}

func HandleConnection(conn net.Conn) {
	<-time.After(40 * time.Microsecond)

	incomingMessage := make([]byte, 1024)
	_, err := conn.Read(incomingMessage)
	if err != nil {
		if errors.Is(err, io.EOF) {
			// panic(err)
			return
		}
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

	if _, err := conn.Write([]byte(httpResponse)); err != nil {
		panic(err)
	}

}
