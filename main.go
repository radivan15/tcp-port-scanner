package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addr := "127.0.0.1:8000"

	// DialTimeout: coba konek, menyerah setelah 2 detik.
	// Mengembalikan dua nilai: koneksi (conn) dan error (err).
	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		fmt.Println("gagal:", err)
		return // keluar dari main
	}

	defer conn.Close()

	fmt.Println("Open:", addr)
}
