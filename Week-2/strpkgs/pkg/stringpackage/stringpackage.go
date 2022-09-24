package stringpackage

import (
	"strings"
)

func Concat(x, y string) string {
	return x + y
}

func ReplaceIt(x, old, new string, n int) string {
	return strings.Replace(x, old, new, n)
}
