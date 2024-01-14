package subtitles

import (
	"github.com/stretchr/testify/assert"
	"subfit/linesitr"
	"testing"
	"time"
)

// Test should parse a single subtitle with on problems.
func TestShouldParseSubtitleSuccessfully(t *testing.T) {
	const sample = `
0
00:01:34,234 --> 00:01:38,001
One upon a time, in a far-far land
There was a king that ...

`
	lines := linesitr.Parse(sample)
	subtitle, err := parseSubtitle(lines)
	assert.Nilf(t, err, "Failed to parse subtitle.")

	startTime, err := time.ParseDuration("1m34s234ms")
	endTime, err := time.ParseDuration("1m38s001ms")

	assert.Equal(t, subtitle.Index, 0)
	assert.Equal(t, subtitle.StartTime, startTime)
	assert.Equal(t, subtitle.EndTime, endTime)
	assert.Equal(t, subtitle.Text, []string{"One upon a time, in a far-far land", "There was a king that ..."})
}

// Test should fail on invalid title index.
func TestShouldFailOnInvalidSubtitleIndex(t *testing.T) {
	const sample = `
A
00:01:34,234 --> 00:01:38,001
One upon a time, in a far-far land
`
	lines := linesitr.Parse(sample)
	_, err := parseSubtitle(lines)
	assert.ErrorContains(t, err, "invalid subtitle index")
}

// Test should fail on invalid title index.
func TestShouldFailOnInvalidTimeframeIndex(t *testing.T) {
	const sample = `
1
00:01:34 --> 00:01:38
One upon a time, in a far-far land
`
	lines := linesitr.Parse(sample)
	_, err := parseSubtitle(lines)
	assert.ErrorContains(t, err, "invalid time frame")
}

// Test should fail on invalid title index.
func TestShouldFailPartialSubtitle(t *testing.T) {
	const sample = "1"
	lines := linesitr.Parse(sample)
	_, err := parseSubtitle(lines)
	assert.ErrorContains(t, err, "premature")
}
