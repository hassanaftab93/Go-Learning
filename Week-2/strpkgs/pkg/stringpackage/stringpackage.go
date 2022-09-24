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

func SpaceTrimmer(s string) string {
	return strings.TrimSpace(s)
}

func Trimmer(s, t string) string {
	return strings.Trim(s, t)
}
