package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	// username, orderNumber := "Ali", "0550554"
	file, err := os.OpenFile("user.csv", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	data := []string{"Ali", "30", "student"}
	wErr := writer.Write(data)

	if wErr != nil {
		log.Fatal(err)
	}

}
