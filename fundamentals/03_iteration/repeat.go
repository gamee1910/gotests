// Package iteration for tdd
package iteration

import "strings"

const repeatedCount = 5

func Repeat(character string) string {
	var repeated strings.Builder
	/*
		//!NOTE: OLD Version
			for i := 0; i < repeatedCount; i++ {
				repeated.WriteString(character)
			}
	*/
	for range repeatedCount {
		repeated.WriteString(character)
	}
	return repeated.String()
}
