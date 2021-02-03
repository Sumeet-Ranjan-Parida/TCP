package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("--------------------------------------")
	fmt.Println("TCP Server")
	fmt.Printf("--------------------------------------\n\n")

	var Port int

	fmt.Printf("Enter your preferred port number: ")
	fmt.Scanf("%d", &Port)

	PORT := ":" + strconv.Itoa(Port)

	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Oops... Something went wrong.")
		return
	}
	defer l.Close()

	fmt.Printf("\n\nStarting Server...")
	fmt.Printf("\nListening on port: %d\n\n", Port)

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Printf("\n\nBye Bye, TCP Server has stopped.")
			return
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}
