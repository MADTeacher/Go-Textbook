package main

import (
	"fmt"
	"log"
	"net"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	s, _ := net.ResolveUDPAddr("udp4", "localhost:7890")
	c, err := net.DialUDP("udp4", nil, s)
	checkErr(err)

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()
	count := 0
	for {
		data := []byte(fmt.Sprintf("%d)Привет!\n", count))
		fmt.Print("->", string(data))
		_, err = c.Write(data)

		checkErr(err)

		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		checkErr(err)

		fmt.Printf("<-: %s", string(buffer[0:n]))
		count++
		if count >= 5 {
			data := []byte("STOP")
			fmt.Println("->", string(data))
			_, err = c.Write(data)
			checkErr(err)

			fmt.Println("Finished!")
			return
		}
	}
}
