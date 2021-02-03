package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("--------------------------------------")
	fmt.Println("TCP Client")
	fmt.Printf("--------------------------------------\n\n")

	var Connect int

	fmt.Printf("Enter the Port number to connect: ")
	fmt.Scanf("%d", &Connect)

	CONNECT := "127.0.0.1" + ":" + strconv.Itoa(Connect)

	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Printf("\n\nTCP client exiting...")
			return
		}
	}
}
