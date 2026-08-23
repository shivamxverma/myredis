package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":6380")

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		fmt.Println("Client connected:", conn.RemoteAddr())
		go handleConnection(conn)
	}
}

func ReadData(buffer []byte, conn net.Conn) (int, error) {
	return conn.Read(buffer)
}

func WriteData(conn net.Conn) (int, error) {
	return conn.Write([]byte("Hello from server\n"))
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		n, err := ReadData(buffer, conn)

		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected:", conn.RemoteAddr())
			} else {
				fmt.Println("read error:", err)
			}
			return
		}

		fmt.Println("Received:", string(buffer[:n]))

		_, err = WriteData(conn)

		if err != nil {
			fmt.Println("write error:", err)
			return
		}
	}
}