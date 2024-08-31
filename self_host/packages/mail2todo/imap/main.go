package main

import (
	"log"

	"github.com/emersion/go-imap/v2/imapclient"
)

func main() {
	// Connect to the server
	c, err := imapclient.DialTLS("imap.gmail.com:993", nil)
	if err != nil {
		log.Fatalf("failed to dial IMAP server: %v", err)
	} else {
		log.Println("connected to IMAP server")
	}
	defer c.Close()
}
