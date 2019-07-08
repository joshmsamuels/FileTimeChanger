package file

import (
	"errors"
	"fmt"
	"os"
	"time"
)

const (
	// TimeFormat is the custom format I created based on RFC1123.
	// I removed the day of the week from my time format since golang's time
	// parser ignores the day of the week after validating it is valid (it does
	// not even ensure the day of the week corresponds to the rest of the date).
	// I removed the seconds from the time format since on my computer seconds
	// are not displayed, however, I am willing to add it back if other machines
	// dispay the seconds of the last time a file was modified.
	// TODO: Find out why timesone must be MST.
	TimeFormat = "02 Jan 2006 15:04 MST"
)

type Data struct {
	// filepath represents the path to the file.
	filepath string

	// modifiedTime represents the new modified time.
	modifiedTime time.Time
}

// NewData creates a file data object with a file path and time modified.
// NewData returns a pointer to the data object that was created.
func NewData(filepath string, modifiedTime time.Time) *Data {
	createdData := Data{}

	createdData.SetFilepath(filepath)
	createdData.SetModifiedTime(modifiedTime)

	return &createdData
}

// NewDataBlank creates a file data object with none of its data set.
// NewDataBlank returns a pointer to the data object that was created.
func NewDataBlank() *Data {
	createdData := Data{}
	return &createdData
}

// GetFilepath returns the filepath stored in the data object.
func (d *Data) GetFilepath() string {
	return d.filepath
}

// GetModifiedTime returns the modified time stored in the data object.
func (d *Data) GetModifiedTime() time.Time {
	return d.modifiedTime
}

// SetFilepath updates the data filepath to the value of "filepath".
func (d *Data) SetFilepath(filepath string) {
	d.filepath = filepath
}

// SetModifiedTime updates the data modified time to the value of "modifiedTime".
func (d *Data) SetModifiedTime(modifiedTime time.Time) {
	d.modifiedTime = modifiedTime
}

// ChangeFileTime consumes modifiedTime and filpath in Data.
// ChangeFileTime works by changing the time modified of the filepath stored in data to
// the modified time stored in data.
// ChangeFileTime returns the first error that occurred.
func (d *Data) ChangeFileTime() error {
	// TODO: Try to keep aTime the same if possible.
	aTime := time.Now()

	err := os.Chtimes(d.GetFilepath(), aTime, d.GetModifiedTime())
	if err != nil {
		return errors.New(fmt.Sprintf("Error changing the time of the file %s to %s: Error: %s", d.GetFilepath(), d.GetModifiedTime(), err))
	}

	return nil
}
