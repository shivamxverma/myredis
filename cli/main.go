package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6380")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)


	for {

		input,err := reader.ReadString('\n');

		if err != nil {
			fmt.Println("Input Error:", err)
			return
		}

		_, err = conn.Write([]byte(input))

		if err != nil {
			fmt.Println("write error:", err)
			return
		}

		buffer := make([]byte, 1024)

		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("read error:", err)
			return
		}

		fmt.Print("Server: ")
		fmt.Print(string(buffer[:n]))
	}
}