package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	connection, err := net.Dial("tcp", "localhost:9000")
	logFatal(err)

	defer connection.Close()

	fmt.Println("Enter your name:")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	logFatal(err)

	username = strings.Trim(username, "\r\n")
	welcomeMSg := fmt.Sprintf("Welcome %s ! write Below to send Messages :-.", username)
	fmt.Println(welcomeMSg)

	Write(connection, username)

}

func Write(connection net.Conn, username string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		message = fmt.Sprintf("%s:- %s\n", username, strings.Trim(message, "\r\n"))
		connection.Write([]byte(message))

	}
}