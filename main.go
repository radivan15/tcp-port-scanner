package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {
	// flag.string mengembalikan *sting (pointer)
	host := flag.String("host", "", "target host to scan")
	port := flag.Int("port", 0, "port number to scan")

	// memproses os.Args menjadi nilai flag
	flag.Parse()

	// pastikan user mengisi kedua flag
	if *host == "" || *port == 0 {
		fmt.Println("Usage: port-scanner -host <host> -port <port>")
		return
	}

	// membuat string "host:port" — format yang dibutuhkan net.DialTimeout
	address := fmt.Sprintf("%s:%d", *host, *port)

	// net.DialTimeout mencoba koneksi TCP dengan batas waktu
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		fmt.Printf("Port %d\tCLOSED\n", *port)
		return
	}

	defer conn.Close()

	fmt.Printf("Port %d\tOPEN\n", *port)
}
