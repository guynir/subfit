package subtitles

import (
	"errors"
	"fmt"
	"strings"
	"subfit/formatutils"
	"subfit/linesitr"
)

// Parse a subtitle block from a given Lines object. The subtitles include index code, time frame
// and actual text.
func parseSubtitle(lines *linesitr.Lines) (*Subtitle, error) {
	// Assuming the first value in a block is the index of the title.
	currentLine, hasCurrent := lines.Current()

	// Skip empty lines.
	for currentLine = strings.TrimSpace(currentLine); len(currentLine) == 0; {
		if currentLine, hasCurrent = lines.Next(); !hasCurrent {
			// If we've reached the end of the file and could not find anything -- we exit gracefully with
			// no subtitle.
			return nil, nil
		}
	}

	index, indexErr := toInt(currentLine, "")
	if indexErr != nil {
		return nil, fmt.Errorf("error on line %d -- invalid subtitle index", lines.LineIndex)
	}

	currentLine, hasCurrent = lines.Next()
	if !hasCurrent {
		return nil, errors.New("invalid subtitle block (premature end of file)")
	}

	startTime, endTime, timeParseErr := parseTimeframeLine(currentLine)
	if timeParseErr != nil {
		return nil, fmt.Errorf("error on line %d -- invalid time frame expression", lines.LineIndex)
	}

	var text []string

	currentLine, hasCurrent = lines.Next()
	if !hasCurrent {
		return nil, errors.New("invalid subtitle block (premature end of file)")
	}

	// Read following (non-empty) lines as subtitle texts.
	// TODO: We should limit the number of lines to 3-4 lines at most. It is unreasonable to have too many text lines.
	for hasCurrent && !formatutils.IsEmpty(currentLine) {
		text = append(text, currentLine)
		currentLine, hasCurrent = lines.Next()
	}

	// If we couldn't read any text, we've reached the end of file prematurely.
	if len(text) == 0 {
		return nil, errors.New("invalid subtitle block (premature end of file)")
	}

	return &Subtitle{Index: index, StartTime: startTime, EndTime: endTime, Text: text}, nil
}
