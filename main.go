package main

import (
	"crypto/rand"
	"fmt"
	"kevinsheeran/walrus-backend/router/walrus"
	"log"
)

const (
	address    = "127.0.0.1:31415"
	epochs     = "1"
	aggregator = "https://aggregator-devnet.walrus.space"
	publisher  = "https://publisher-devnet.walrus.space"
)

func main() {

	randomData := make([]byte, 1024*1024)
	if _, err := rand.Read(randomData); err != nil {
		log.Fatalf("failed to read random data: %v", err)
	}
	blobId, err := walrus.UploadBlob(address, epochs, randomData)
	if err != nil {
		log.Fatalf("failed to upload blob: %v", err)
	}

	fmt.Printf("Blob ID: %s\n", blobId)
}
