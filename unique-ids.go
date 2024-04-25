package uniqueids

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	lastTimestamp  float64
	sequenceNumber int64
	mutex          sync.Mutex
	epoch          time.Time = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC) // Hard-coded epoch
)

// generateUniqueNumber generates a unique identifier using a timestamp and sequence number.
func generateUniqueNumber() string {
	mutex.Lock()
	defer mutex.Unlock()

	now := float64(time.Since(epoch).Nanoseconds()) / 1e9 // Get the number of seconds since the hard-coded epoch as float64
	if now == lastTimestamp {
		sequenceNumber++
		if sequenceNumber > 99 { // Limit the sequence number to 2 digits
			// If the sequence number exceeds 99, wait for the next second
			for now == lastTimestamp {
				now = float64(time.Since(epoch).Nanoseconds()) / 1e9
			}
			sequenceNumber = 0
		}
	} else {
		sequenceNumber = 0
		lastTimestamp = now
	}

	return fmt.Sprintf("%d%02d", int64(now), sequenceNumber)
}

// parseUniqueID parses the unique ID and converts it back to UTC and local time.
func parseUniqueID(uniqueID string) (utcTime time.Time, localTime time.Time, sequenceNum int, err error) {
	if len(uniqueID) < 12 {
		err = fmt.Errorf("invalid unique ID length")
		return
	}

	timestampStr := uniqueID[:len(uniqueID)-2]
	sequenceStr := uniqueID[len(uniqueID)-2:]

	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		err = fmt.Errorf("invalid timestamp format: %v", err)
		return
	}
	sequenceNum, err = strconv.Atoi(sequenceStr)
	if err != nil {
		err = fmt.Errorf("invalid sequence number format: %v", err)
		return
	}

	utcTime = epoch.Add(time.Duration(timestamp) * time.Second).UTC()
	localTime = utcTime.Local()
	return
}

// GenerateMatterNumber generates a unique matter number.
func GenerateMatterNumber() string {
	return generateUniqueNumber()
}

// GenerateInvoiceNumber generates a unique invoice number with a specified prefix.
func GenerateInvoiceNumber(prefix string) string {
	return fmt.Sprintf("%s%s", prefix, generateUniqueNumber())
}
