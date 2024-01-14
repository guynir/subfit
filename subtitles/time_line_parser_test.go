package subtitles

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//
// Unit tests for 'time_line_parser' file.
//

// Test should successfully parse a valid time frame line to start/end durations..
func TestShouldSuccessfullyParseTimeframeLine(t *testing.T) {
	var expectedStartTime, _ = time.ParseDuration("0h12m34s567ms")
	var expectedEndTime, _ = time.ParseDuration("01h43m66s899ms")
	var line string = "00:12:34,567 --> 01:43:66,899"
	startTime, endTime, _ := parseTimeframeLine(line)

	assert.Equal(t, startTime, expectedStartTime, "invalid start time")
	assert.Equal(t, endTime, expectedEndTime, "invalid end time")
}

// Test should fail on invalid timeframe line.
func TestShouldFailOnInvalidLineFormat(t *testing.T) {
	var line string = "something with garbage ....."
	_, _, err := parseTimeframeLine(line)
	assert.NotNil(t, err, "Test should have failed.")
}

// Test should fail on invalid property of either start or end times. The format of the timeframe
// line is valid, though.
func TestShouldFailOnInvalidDate(t *testing.T) {
	_, _, err := parseTimeframeLine("0a:12:34,567 --> 01:43:66,899")
	assert.NotNil(t, err, "Test should have failed.")

	_, _, err = parseTimeframeLine("00:12:34,567 --> 0:43:66,899")
	assert.NotNil(t, err, "Test should have failed.")
}
