package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

// scanPort mencoba koneksi TCP ke satu port, return true jika open
func scanPort(host string, port int, timeout time.Duration) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func main() {
	host := flag.String("host", "", "target host to scan (e.g. scanme.nmap.org)")
	port := flag.Int("port", 0, "single port to scan (e.g. 80)")
	start := flag.Int("start", 0, "start of port range (e.g. 1)")
	end := flag.Int("end", 0, "end of port range (e.g. 1024)")

	flag.Parse()

	if *host == "" {
		fmt.Println("Usage:")
		fmt.Println("  port-scanner -host <host> -port <port>")
		fmt.Println("  port-scanner -host <host> -start <start> -end <end>")
		return
	}

	// mode 1: scan single port (sama seperti Stage 1)
	if *port != 0 {
		if scanPort(*host, *port, 3*time.Second) {
			fmt.Printf("Port %d\tOPEN\n", *port)
		} else {
			fmt.Printf("Port %d\tCLOSED\n", *port)
		}
		return
	}

	// mode 2: scan range — validasi dulu
	if *start == 0 || *end == 0 || *start > *end {
		fmt.Println("Error: provide valid -start and -end (start must be <= end)")
		return
	}

	fmt.Printf("Scanning %s port %d-%d...\n", *host, *start, *end)

	openCount := 0
	// for loop di Go — satu-satunya cara looping
	// i mulai dari *start, naik 1 setiap iterasi, sampai *end
	for i := *start; i <= *end; i++ {
		if scanPort(*host, i, 3*time.Second) {
			fmt.Printf("Port %d\tOPEN\n", i)
			openCount++
		}
	}

	total := *end - *start + 1
	fmt.Printf("\nTotal: %d open out of %d ports\n", openCount, total)

}
