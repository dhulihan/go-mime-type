package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// https://cs.opensource.google/go/go/+/refs/tags/go1.20.3:src/net/http/sniff.go;l=13
const BYTES_TO_READ = 512

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("please provide file path")
		os.Exit(1)
	}

	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("could not open file \"%s\": %s", path, err)
	}

	// read as bytes
	defer file.Close()

	// read minimal number of bytes to determine content type
	buf := make([]byte, BYTES_TO_READ)
	if _, err := io.ReadFull(file, buf); err != nil {
		log.Fatal(err)
	}

	mimeType := http.DetectContentType(buf)

	fmt.Println(mimeType)
}
