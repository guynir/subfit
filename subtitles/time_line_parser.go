package subtitles

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

const InvalidDurationValue = time.Duration(-1 << 63)

// Holds a regular expression to parse timeframe expression.
var timeFrameRegex = createExpression()

func parseTimeframeLine(line string) (time.Duration, time.Duration, error) {
	var err error = nil
	var startTime = InvalidDurationValue
	var endTime = InvalidDurationValue

	// Parse time frame line into array list containing hour/minute/second/millis of both start and end times.
	var match = timeFrameRegex.FindStringSubmatch(line)
	if match == nil {
		err = errors.New("invalid time frame expression")
	} else {
		startTime, err = createDuration(match[1], match[2], match[3], match[4])
		if err == nil {
			endTime, err = createDuration(match[5], match[6], match[7], match[8])
		}

		// If any error occurred while parsing start/end times, provide a generic
		// error response.
		if err != nil {
			err = errors.New("invalid time frame expression")
		}
	}

	return startTime, endTime, err
}

// Converts hour/minutes/seconds/milliseconds to a time.Duration object.
//
// Accepts hours, minutes, second and millis as string and convert them into integer values,
// internally.
// All errors are accumulated and returned (if there were errors) vai the second returned value.
func createDuration(hour, minute, second, millis string) (time.Duration, error) {
	hourInt, errHour := toInt(hour, "invalid hours value")
	minuteInt, errMinute := toInt(minute, "invalid minutes value")
	secondInt, errSecond := toInt(second, "invalid seconds value")
	millisInt, errMillis := toInt(millis, "invalid milliseconds value")

	var durationMicro time.Duration
	var err error = errors.Join(errHour, errMinute, errSecond, errMillis)
	if err == nil {
		durationMicro = time.Hour*time.Duration(hourInt) +
			time.Minute*time.Duration(minuteInt) +
			time.Second*time.Duration(secondInt) +
			time.Millisecond*time.Duration(millisInt)
	} else {
		durationMicro = InvalidDurationValue
	}

	return durationMicro, err
}

// Converts a string value to an integer. If operation fails, return error object with predefined
// message.
func toInt(value string, errorMessage string) (int, error) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		err = errors.New(errorMessage)
		intValue = 0
	}
	return intValue, err
}

func createExpression() *regexp.Regexp {
	expression := "" +
		"^(?P<StartHour>\\d{2}):(?P<StartMinute>\\d{2}):(?P<StartSecond>\\d{2}),(?P<StartMilli>\\d+)" +
		" --> " +
		"(?P<EndHour>\\d{2}):(?P<EndMinute>\\d{2}):(?P<EndSecond>\\d{2}),(?P<EndMillis>\\d+)"

	re, _ := regexp.Compile(expression)
	return re
}
