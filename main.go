package main

import (
	"fmt"
	"github.com/mariosplit/unique-ids"
	"time"
)

func main() {
	// Generate unique matter numbers
	for i := 0; i < 5; i++ {
		matterNumber := uniqueids.GenerateMatterNumber()
		fmt.Println("Matter Number:", matterNumber)
	}

	fmt.Println()

	// Generate unique invoice numbers
	for i := 0; i < 5; i++ {
		invoiceNumber := uniqueids.GenerateInvoiceNumber("INV")
		fmt.Println("Invoice Number:", invoiceNumber)
	}

	fmt.Println()

	// Parse a unique ID
	uniqueID := "10699800"
	utcTime, localTime, sequenceNum, err := uniqueids.ParseUniqueID(uniqueID)
	if err != nil {
		fmt.Println("Error parsing unique ID:", err)
	} else {
		fmt.Printf("Unique ID: %s\n", uniqueID)
		fmt.Printf("UTC Time: %s\n", utcTime)
		fmt.Printf("Local Time: %s\n", localTime)
		fmt.Printf("Sequence Number: %d\n", sequenceNum)
	}

	fmt.Println()

	// Simulate generating unique IDs over time
	for i := 0; i < 10; i++ {
		matterNumber := uniqueids.GenerateMatterNumber()
		fmt.Printf("Matter Number: %s - Generated at: %s\n", matterNumber, time.Now().Format("15:04:05.000"))
		time.Sleep(100 * time.Millisecond)
	}
}
