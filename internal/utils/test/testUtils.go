package test

import (
	"testing"
	"time"
)

func EnsureParsedTime(t *testing.T, timeFormat, timeStr string) (parsed time.Time) {
	parsed, err := time.Parse(timeFormat, timeFormat)

	if err != nil {
		t.Errorf("Error parsing time: (%v)", err)
	}

	return
}
