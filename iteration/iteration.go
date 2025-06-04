package iteration

import "strings"

func Repeat(character string, repeatCount uint) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(character)
	}
	return repeated.String()
}

// - example test
// - look through strings lib and write examples and tests
