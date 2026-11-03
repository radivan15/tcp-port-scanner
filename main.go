package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {
	// Definisikan flag: nama, nilai default, teks bantuan.
	// Hasilnya POINTER (*string / *int), bukan nilai langsung.
	host := flag.String("host", "127.0.0.1", "target host")
	port := flag.Int("port", 8000, "target port")

	// baca argumen dari cli
	flag.Parse()

	// *host = ambil isi dari pointer. Sprintf menggabungkan jadi "host:port".
	addr := fmt.Sprintf("%s:%d", *host, *port)

	// DialTimeout: coba konek, menyerah setelah 2 detik.
	// Mengembalikan dua nilai: koneksi (conn) dan error (err).
	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		fmt.Println("Failed:", err)
		return
	}

	defer conn.Close()

	fmt.Println("Open:", addr)
}
