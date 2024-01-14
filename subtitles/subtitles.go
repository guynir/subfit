package subtitles

import (
	"time"
)

// Subtitles interface represents a subtitles resource that can be loaded and parsed into a list of subtitles blocks.
type Subtitles interface {
	// Load subtitles from external resource.
	Load() error

	// Save subtitles back to original resource.
	Save() error

	// Reset in-memory state of subtitles to original state (as if it were loaded).
	Reset() error

	// SaveTo saves the subtitles to a target location.
	SaveTo(path string) error

	// Adjust the subtitles in a given seconds offset.
	Adjust(offset float32)
}

// Subtitle represents a single subtitle instance.
type Subtitle struct {
	Index       int
	StartTime   time.Duration
	EndTime     time.Duration
	Coordinates string
	Text        []string
}
