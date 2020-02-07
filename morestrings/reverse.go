// Package morestring implements additional functionality to manipulate UTF-8
// encoded strings, beyond what is provide in the standard "strings" packageStorage.configure({
package morestrings

func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
