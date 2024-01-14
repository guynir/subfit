package formatutils

import "time"

// ToDuration convert hours/minutes/seconds/millis components into a single duration object.
func ToDuration(hours int, minutes int, seconds int, millis int) time.Duration {
	return time.Hour*time.Duration(hours) +
		time.Minute*time.Duration(minutes) +
		time.Second*time.Duration(seconds) +
		time.Millisecond*time.Duration(millis)
}

// DurationToComponents parse a given duration and return hours/minutes/seconds/milliseconds components.
func DurationToComponents(d time.Duration) (int, int, int, int) {
	var hours, minutes, seconds, millis int
	hours, d = getComponentAs(d, time.Hour)
	minutes, d = getComponentAs(d, time.Minute)
	seconds, d = getComponentAs(d, time.Second)
	millis, d = getComponentAs(d, time.Millisecond)

	return hours, minutes, seconds, millis
}

// Helper function for parsing a component from a Duration object.
// Returns a new duration object without the extracted component.
func getComponentAs(d time.Duration, componentFactor time.Duration) (int, time.Duration) {
	return int(d / componentFactor), d % componentFactor
}
