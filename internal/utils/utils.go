package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func HandleGenericErr(context string, err error) {
	if err == nil {
		return
	}

	// User-friendly error message about what went wrong
	fmt.Printf("Message: %s\n", context)

	// Print extra debug information for devs
	_, envPresent := os.LookupEnv("FTC_DEV")
	if envPresent {
		panic(err)
	}

	os.Exit(1)
}

// readString is a wrapper around the bufio readString function that strips the delimiter and trims the string.
// this will keep my main cleaner.
func ReadString(reader *bufio.Reader, delim byte) (string, error) {
	str, err := reader.ReadString(delim)

	if len(str) > 0 && str[len(str)-1] == delim {
		str = str[:len(str)-1]
	}

	return strings.TrimSpace(str), err
}
