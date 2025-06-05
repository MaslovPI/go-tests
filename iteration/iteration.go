package iteration

import "strings"

func Repeat(character string, repeatCount uint) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func RepeatWithSeparation(character, separator string, repeatCount uint) string {
	var repeated strings.Builder
	for i := range repeatCount {
		repeated.WriteString(character)
		if i != repeatCount-1 {
			repeated.WriteString(separator)
		}
	}
	return repeated.String()
}
