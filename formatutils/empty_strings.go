package formatutils

const whiteSpaceCharacters = " \t\n\r\f\v"

// IsEmpty provides indication if the given string is empty or not.
// An empty string is one that either has zero length of contains only white-spaces (i.e.: tab, space, newline, ...).
// This is a memory and speed optimal implementation.
func IsEmpty(str string) bool {
	for idx := 0; idx < len(str); idx++ {
		ch := str[idx]
		// Check if current character is either space, tab, new line (linefeed), carriage return, form feed or
		// vertical tab.
		if ch != ' ' && ch != '\t' && ch != '\n' && ch != '\r' && ch != '\f' && ch != '\v' {
			return false
		}
	}

	return true
}
