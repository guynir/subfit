package formatutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// Test should parse a duration into hours/minutes/seconds/milliseconds components.
func TestShouldParseDurationToComponents(t *testing.T) {
	// Sample values.
	var hour, minute, second, millis = 3, 29, 33, 34

	// Convert time to duration.
	var duration time.Duration = ToDuration(hour, minute, second, millis)

	// Parse duration back to i
	parsedHours, parsedMinutes, parsedSeconds, parsedMillis := DurationToComponents(duration)

	// Assert that all components are parsed correctly.
	assert.Equal(t, hour, parsedHours)
	assert.Equal(t, minute, parsedMinutes)
	assert.Equal(t, second, parsedSeconds)
	assert.Equal(t, millis, parsedMillis)
}
