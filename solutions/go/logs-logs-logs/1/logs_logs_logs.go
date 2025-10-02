package logs

import "unicode/utf8"

type application map[rune]string

var apps = application{
    '❗': "recommendation",
    '🔍': "search",
    '☀': "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
        app, ok := apps[r]
        if ok {
            return app
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    runes := []rune(log)
	for idx, og := range runes {
        if og == oldRune {
            runes[idx] = newRune
        }
    }
    return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	numberOfRunes := utf8.RuneCountInString(log)
    return numberOfRunes <= limit
}
