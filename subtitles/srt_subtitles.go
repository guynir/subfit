package subtitles

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"subfit/formatutils"
	"subfit/linesitr"
	"time"
)

const newLine = linesitr.LinuxNewLine

// A SRTSubtitlesFile contains subtitles from .SRT file.
type SRTSubtitlesFile struct {
	filename         string
	originalContents string
	lastError        error

	subtitles []Subtitle
}

// Load contents of SRT file into memory.
func (subtitles *SRTSubtitlesFile) Load() error {
	content, readErr := os.ReadFile(subtitles.filename)
	if readErr != nil {
		return readErr
	}

	// Keep on the original contents. If we call "Reset", original contents are re-loaded.
	subtitles.originalContents = string(content)
	return nil
}

// Save writes data to original SRT file.
func (subtitles *SRTSubtitlesFile) Save() error {
	return subtitles.SaveTo(subtitles.filename)
}

// SaveTo writes data to a file in SRT format.
func (subtitles *SRTSubtitlesFile) SaveTo(path string) error {
	buf := strings.Builder{}

	for _, st := range subtitles.subtitles {
		buf.WriteString(strconv.Itoa(st.Index))
		buf.WriteString(newLine)
		buf.WriteString(durationToString(st.StartTime))
		buf.WriteString(" --> ")
		buf.WriteString(durationToString(st.EndTime))
		buf.WriteString(newLine)
		for _, text := range st.Text {
			buf.WriteString(text)
			buf.WriteString(newLine)
		}
		buf.WriteString(newLine)
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = file.WriteString(buf.String())
	closeErr := file.Close()

	if err != nil {
		return err
	} else {
		return closeErr
	}
}

func (subtitles *SRTSubtitlesFile) Reset() error {
	// Reset subtitles list contents.
	subtitles.subtitles = []Subtitle{}
	var err error

	// Parse contents.
	subtitles.subtitles, err = parseContents(subtitles.originalContents)
	if err != nil {
		err = fmt.Errorf("failed to parse SRT contents (%s)", err.Error())
	}

	return nil
}

// Adjust the subtitles by a given offset of seconds. The seconds may contain a decimal
// fraction representing milliseconds.
func (subtitles *SRTSubtitlesFile) Adjust(offset float32) {
	adjustDuration := time.Duration(offset * float32(time.Second))

	for idx := 0; idx < len(subtitles.subtitles); idx++ {
		subtitles.subtitles[idx].StartTime += adjustDuration
		subtitles.subtitles[idx].EndTime += adjustDuration
	}
}

// Parse the contents maintained in an SRT subtitle struct and generate
// list of child 'Subtitle' objects.
func parseContents(contents string) ([]Subtitle, error) {
	var subtitlesList []Subtitle
	var err error
	if !formatutils.IsEmpty(contents) {
		var lines = linesitr.Parse(contents)
		for lines.HasNext() {
			var subtitle *Subtitle
			subtitle, err = parseSubtitle(lines)
			if err != nil {
				break
			}

			if subtitle != nil {
				subtitlesList = append(subtitlesList, *subtitle)
			}
		}
	}

	return subtitlesList, nil
}

// New creates a new SRTSubtitlesFile instance and initialize it (load contents from file).
func New(filename string) (Subtitles, error) {
	newSubtitles := new(SRTSubtitlesFile)
	newSubtitles.filename = filename
	err := newSubtitles.Load()
	if err != nil {
		return nil, err
	}

	newSubtitles.subtitles, err = parseContents(newSubtitles.originalContents)
	if err != nil {
		newSubtitles = nil
	}
	return newSubtitles, err
}

// Convert a duration object to HH:mm:SS,sss format.
func durationToString(d time.Duration) string {
	hours, minutes, seconds, millis := formatutils.DurationToComponents(d)
	return fmt.Sprintf("%02d:%02d:%02d,%03d", hours, minutes, seconds, millis)
}
