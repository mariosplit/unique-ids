# Unique IDs

The `uniqueids` package provides functionality for generating unique identifiers based on timestamps and sequence numbers. It allows you to generate unique matter numbers and invoice numbers with customizable prefixes.

## Features

- Generates unique identifiers using a combination of timestamp and sequence number
- Supports customizable epoch for timestamp calculation
- Provides functions for generating unique matter numbers and invoice numbers
- Allows parsing of generated unique IDs to extract timestamp and sequence number

## Installation

To use the `uniqueids` package in your Go project, you can install it using `go get`:

```shell
go get github.com/mariosplit/uniqueids
```

## Usage

### Generating Unique IDs

To generate a unique matter number, use the `GenerateMatterNumber` function:

```go
matterNumber := uniqueids.GenerateMatterNumber()
```

To generate a unique invoice number with a custom prefix, use the `GenerateInvoiceNumber` function:

```go
invoiceNumber := uniqueids.GenerateInvoiceNumber("INV")
```

### Parsing Unique IDs

To parse a generated unique ID and extract the timestamp and sequence number, use the `ParseUniqueID` function:

```go
utcTime, localTime, sequenceNum, err := uniqueids.ParseUniqueID(uniqueID)
if err != nil {
    // Handle the error
} else {
    // Use the parsed values
    fmt.Printf("UTC Time: %s\nLocal Time: %s\nSequence Number: %d\n", utcTime, localTime, sequenceNum)
}
```

The `ParseUniqueID` function returns the UTC time, local time, sequence number, and any error that occurred during parsing.

### Customizing Epoch

By default, the `uniqueids` package uses a hard-coded epoch of January 1, 2000, for timestamp calculation. If you want to use a different epoch, you can modify the `epoch` variable in the package:

```go
uniqueids.epoch = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
```

### Concurrency Safety

The `uniqueids` package is designed to be concurrency-safe. It uses a mutex to synchronize access to the shared state (last timestamp and sequence number) when generating unique IDs.

## Example

Here's a complete example of using the `uniqueids` package:

```go
package main

import (
    "fmt"
    "github.com/mariosplit/uniqueids"
)

func main() {
    // Generate unique IDs
    matterNumber := uniqueids.GenerateMatterNumber()
    invoiceNumber := uniqueids.GenerateInvoiceNumber("INV")

    fmt.Println("Matter Number:", matterNumber)
    fmt.Println("Invoice Number:", invoiceNumber)

    // Parse unique IDs
    utcTime, localTime, sequenceNum, err := uniqueids.ParseUniqueID(matterNumber)
    if err != nil {
        fmt.Println("Error parsing Matter Number:", err)
    } else {
        fmt.Printf("Matter Number - UTC Time: %s\nMatter Number - Local Time: %s\nMatter Number - Sequence Number: %d\n",
            utcTime, localTime, sequenceNum)
    }
}
```

## License

The `uniqueids` package is open-source software licensed under the [MIT License](LICENSE).
