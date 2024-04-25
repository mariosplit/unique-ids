package uniqueids

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
)

var (
	lastTimestamp  int64
	sequenceNumber int64
	mutex          sync.Mutex
	epoch          time.Time = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC) // Hard-coded epoch
)

func Epoch() time.Time {

	fmt.Println("Epoch is: ", epoch)
	return epoch
}

// GenerateUniqueNumber generates a unique identifier using a timestamp and sequence number.
func GenerateUniqueNumber() string {
	mutex.Lock()
	defer mutex.Unlock()

	now := time.Now().Unix() // Get current time in seconds since epoch
	if now == lastTimestamp {
		sequenceNumber++
		if sequenceNumber > 99 {
			// Force wait for the next second if sequence exceeds 99
			time.Sleep(time.Second) // Explicitly wait for one second
			now = time.Now().Unix() // Update 'now' to the current time after sleep
			sequenceNumber = 0      // Reset sequence number
		}
	} else {
		sequenceNumber = 0 // Reset sequence number if new second
	}
	lastTimestamp = now // Update the last timestamp

	return fmt.Sprintf("%d%02d", now, sequenceNumber)
}

// ParseUniqueID parses the unique ID and converts it back to UTC and local time.
func ParseUniqueID(uniqueID string) (utcTime time.Time, localTime time.Time, sequenceNum int, err error) {
	// Remove non-numeric prefix by finding the first numeric character
	start := strings.IndexFunc(uniqueID, func(r rune) bool {
		return unicode.IsDigit(r)
	})
	if start == -1 {
		err = fmt.Errorf("no numeric part found in unique ID")
		return
	}

	numericID := uniqueID[start:]

	if len(numericID) < 12 {
		err = fmt.Errorf("numeric part of the unique ID is too short")
		return
	}

	timestampStr := numericID[:len(numericID)-2]
	sequenceStr := numericID[len(numericID)-2:]

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

	// Calculate times based on the parsed timestamp
	//epoch := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	utcTime = epoch.Add(time.Duration(timestamp) * time.Second).UTC()
	localTime = utcTime.Local()
	return
}

// GenerateMatterNumber generates a unique matter number.
func GenerateMatterNumber() string {
	return GenerateUniqueNumber()
}

// GenerateInvoiceNumber generates a unique invoice number with a specified prefix.
func GenerateInvoiceNumber(prefix string) string {
	return fmt.Sprintf("%s%s", prefix, GenerateUniqueNumber())
}
