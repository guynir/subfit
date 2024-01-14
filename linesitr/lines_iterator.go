// Package linesitr provides a simple line iterator over a list of lines.
package linesitr

import (
	"strings"
	"subfit/formatutils"
)

// Various newlines markers.
const WindowsNewLine string = "\r\n"
const WindowsVariantNewLine string = "\n\r"
const LinuxNewLine = "\n"

// Lines contains list of liens and a pointer to the current line.
type Lines struct {
	lines     []string
	LineIndex int
}

// Length return the number of lines in object.
func (lines *Lines) Length() int {
	return len(lines.lines)
}

// The HasNext function provides indication if there are more lines to read.
func (lines *Lines) HasNext() bool {
	return len(lines.lines) > 0 && lines.LineIndex+1 < len(lines.lines)
}

// Next function advances the line pointer to the next available line and returns it.
// The function returns a string and flag indicating whether there was a line to read or not.
func (lines *Lines) Next() (string, bool) {
	var str = ""
	var available = lines.HasNext()
	if available {
		lines.LineIndex += 1
		str = lines.lines[lines.LineIndex]
	}

	return str, available
}

// Current provides a peek for the current line. It does not change any state.
// Returns the current line, if available and flag indicating whether there is a current line or not.
func (lines *Lines) Current() (string, bool) {
	var index = lines.LineIndex

	// In order for us to read the "current line", we first need to make sure we called "Next" at least once,
	// and we have not consumed all lines.
	var hasCurrent = index >= 0 && index < len(lines.lines)
	var line string
	if hasCurrent {
		line = lines.lines[lines.LineIndex]
	}

	return line, hasCurrent
}

// HasAtLeast check if there are at least 'desiredCount' lines available for reading.
func (lines *Lines) HasAtLeast(desiredCount int) bool {
	return lines.LineIndex+desiredCount < len(lines.lines)
}

// SkipEmptyLines read lines until the firsts non-empty line.
// A non-empty line is a line that contains any character other than white space characters.
func (lines *Lines) SkipEmptyLines() {
	// If we don't have anything more to read -- exit.
	if !lines.HasNext() {
		return
	}

	var currentLine string
	var hasLine bool

	// If we haven't started reading from lines -- start now. Otherwise, look at the current location.
	if lines.LineIndex < 0 {
		currentLine, hasLine = lines.Next()
	} else {
		currentLine, hasLine = lines.Current()
	}

	// While we encounter empty lines -- skip to the next line.
	for formatutils.IsEmpty(currentLine) && hasLine {
		currentLine, hasLine = lines.Next()
	}
}

// New Constructs a new instance of Lines.
func New(lines []string) *Lines {
	return &Lines{lines: lines, LineIndex: -1}
}

// Parse a given string contents to lines and return a new 'Lines' object.
func Parse(contents string) *Lines {
	normalizedContents := strings.ReplaceAll(contents, WindowsNewLine, LinuxNewLine)
	normalizedContents = strings.ReplaceAll(normalizedContents, WindowsVariantNewLine, LinuxNewLine)
	lines := strings.Split(normalizedContents, LinuxNewLine)

	return New(lines)
}
