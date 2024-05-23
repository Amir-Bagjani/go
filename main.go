package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("users.json", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := map[string]interface{}{
		"username": "olodocoder",
		"twitter":  "@olodocoder",
		"email":    "hello@olodocoder.com",
		"website":  "https://dev.to/olodocoder",
		"location": "Lagos, Nigeria",
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}
