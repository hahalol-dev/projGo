package main

import (
	"fmt"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"google.golang.org/api/option"
)

func main() {
	// Use something from gopenpgp
	key, _ := crypto.GenerateKey("Test", "test@example.com", "rsa", 1024)
	fmt.Println("Generated key:", key)

	// Use something from google.golang.org/api
	option.WithAPIKey("dummy-api-key")
	fmt.Println("API Key Option created.")
}
