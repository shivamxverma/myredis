package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
)

// Store added 

var (
	store = make(map[string]string)
	mu    sync.RWMutex
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

func WriteData(input string, conn net.Conn) (int, error) {
	trimmed := strings.TrimSpace(input)
	fields := strings.Fields(trimmed)

	if len(fields) == 0 {
		return 0, nil
	}

	command := strings.ToUpper(fields[0])

	// Logic added 

	switch command {
	case "SET":
		if len(fields) < 3 {
			return conn.Write([]byte("ERR wrong number of arguments for 'SET' command\n"))
		}
		key := fields[1]
		value := strings.Join(fields[2:], " ")

		mu.Lock()
		store[key] = value
		mu.Unlock()

		return conn.Write([]byte("OK\n"))

	case "GET":
		if len(fields) < 2 {
			return conn.Write([]byte("ERR wrong number of arguments for 'GET' command\n"))
		}
		key := fields[1]

		mu.RLock()
		val, exists := store[key]
		mu.RUnlock()

		if !exists {
			return conn.Write([]byte("(nil)\n"))
		}
		return conn.Write([]byte(val + "\n"))
	}

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

		_, err = WriteData(string(buffer[:n]),conn)

		if err != nil {
			fmt.Println("write error:", err)
			return
		}
	}
}