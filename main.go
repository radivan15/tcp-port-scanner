package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"time"
)

// scanPort mengecek satu port TCP dan mengembalikan statusnya.
// Parameter: host, port, timeout. Return: string ("open" atau "closed").
func scanPort(host string, port int, timeout time.Duration) string {
	addr := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {

		var netErr net.Error

		// errors.As: "apakah di dalam err ada net.Error?"
		// Kalau ada, isi ke netErr dan hasilnya true.
		if errors.As(err, &netErr) && netErr.Timeout() {
			return "filtered"
		}
		return "closed"
	}

	defer conn.Close()

	return "open"

}

func main() {
	host := flag.String("host", "127.0.0.1", "target host")
	port := flag.Int("port", 8000, "target port")
	flag.Parse()

	// Panggil fungsi. *host dan *port: ambil nilai dari pointer.
	status := scanPort(*host, *port, 2*time.Second)

	fmt.Printf("%s:%d %s\n", *host, *port, status)
}
