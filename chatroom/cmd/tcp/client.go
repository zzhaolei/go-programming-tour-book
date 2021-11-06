package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":2020")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = conn.Close()
	}()

	go func() {
		_, _ = io.Copy(os.Stdout, conn)
		log.Println("Server disconnected.")
		os.Exit(2)
	}()

	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
