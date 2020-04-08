package cache

import (
	"fmt"
	"strings"
	"unicode"
)

func prepareKey(kType KeyType, name string) string {
	trimmer := strings.TrimFunc(name, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && unicode.IsSpace(r)
	})
	return fmt.Sprintf("%s_%s", kType, strings.ToUpper(trimmer))
}
