package main

import (
	// "fmt"
	"net"
	"strings"
)

func main(){
	listener, _ := net.Listen("tcp", ":8000")
	defer listener.Close()

	for {
		conn, _ := listener.Accept()

		go func() {
			defer conn.Close()
			buffer := make([]byte, 1024)
			n,_ := conn.Read(buffer)
			
			request := string(buffer[:n])

			if strings.HasPrefix(request, "GET /users") {
				response := 
					"HTTP/1.1 201 OK\r\n" +
					"Content-Type: application/json\r\n" +
					"\r\n" +
					`{"message":"user returns"}`

				conn.Write([]byte(response))
			}


			// conn.Write([]byte(
			// 	"HTTP/1.1 200 OK\r\n" +
			// 	"Content-Type: text/plain\r\n" +
			// 	"\r\n" + 
			// 	"Hello from the go",
			// ))
		}()
	}
}
