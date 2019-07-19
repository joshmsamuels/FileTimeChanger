package file

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewData(t *testing.T) {
	tests := []struct {
		filepath string
		timeStr  string
	}{
		{
			"test/data/blankFile.txt",
			"02 Jan 2006 15:04 MST",
		},
		{
			"",
			"02 Jan 2006 15:04 MST",
		},
	}

	for _, test := range tests {
		parsedTime, err := time.Parse(TimeFormat, test.timeStr)
		if err != nil {
			t.Errorf("Error parsing time: (%v)", err)
		}

		expectedResult := Data{}
		expectedResult.SetFilepath(test.filepath)
		expectedResult.SetModifiedTime(parsedTime)

		assert.Equal(t, &expectedResult, NewData(test.filepath, parsedTime))
	}
}
