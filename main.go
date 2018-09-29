package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/remeh/sizedwaitgroup"
)

var swg = sizedwaitgroup.New(25) // 8 Being the amount of conncurent tasks allowed to run
var ip string

func main() {
	fmt.Println("Concurrent Port Scanner")

	fmt.Print("Enter Ip Address: ")
	fmt.Scan(&ip)

	for port := 1; port < 65535; port++ { // Full amount of scanable ports
		swg.Add()
		go ScanPorts(ip, port) // Go routine starts here

	}
	swg.Wait() // Wait for all go routines to finish executing

	fmt.Println("All ports have been scanned!")
	os.Exit(0)
}

func ScanPorts(ip string, port int) {
	strf := fmt.Sprintf("%s:%d", ip, port)

	_, err := net.DialTimeout("tcp", strf, 10*time.Millisecond)

	if err != nil {
		//fmt.Printf("- Port %d is closed \n", port)
	} else {
		fmt.Printf("- Port %d is open \n", port)
	}

	defer swg.Done() // Calls when all routines are called

}
