package linesitr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test should successfully parse texts successfully.
func TestShouldParseContentsSuccessfully(t *testing.T) {
	assert.Equal(t, 1, Parse("").Length())
	assert.Equal(t, 2, Parse("One upon a time\n").Length())
	assert.Equal(t, 3, Parse("One upon a time\r\nIn a far-far land\nThere was a mighty king").Length())
}

// Test should successfully indicate if there are more lines to read before initial read and after first read.
func TestShouldIndicateThereAreMoreLinesToRead(t *testing.T) {
	lines := Parse("Once upon a time.\n")

	// Try 'HasNext' before starting reading lines.
	assert.True(t, lines.HasNext())

	// Read the first line.
	lines.Next()

	// Try 'HasNext' after the first read.
	assert.True(t, lines.HasNext())
}

// Test should read all available lines.
func TestShouldReadAllLines(t *testing.T) {
	lines := Parse("Red\nGreen\nBlue")

	var colors []string
	for lines.HasNext() {
		colorName, _ := lines.Next()
		colors = append(colors, colorName)
	}

	assert.Equal(t, colors, []string{"Red", "Green", "Blue"})
}
