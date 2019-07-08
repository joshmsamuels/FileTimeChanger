package validator

import (
	"FileTimeChanger/internal/file"
	"os"
	"time"
)

// InputValidator is used as the function prototype of all methods that can be
// used to validate input.
type InputValidator func(string) bool

// ValidateFilepath ensures the path passed in is reachable and is not a
// directory.
// ValidateFilepath returns false when there is an error getting the file
// information or if the path is a directory and returns true otherwise.
func ValidateFilepath(path string) bool {
	fi, err := os.Stat(path)
	if err != nil || fi.IsDir() {
		return false
	}

	return true
}

// ValidateDate parses the date string with the application's time format to
// ensure the date is valid.
// ValidateDate returns false when there is an error parsing the date and
// returns true otherwise.
func ValidateDate(date string) bool {
	if _, err := time.Parse(file.TimeFormat, date); err != nil {
		return false
	}

	return true
}
