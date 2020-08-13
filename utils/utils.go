package utils

import (
	"bytes"
	"regexp"
	"strings"
	"unicode"
)

const (
	cameling string = `[\p{L}\p{N}]+`
)

var (
	rxCameling = regexp.MustCompile(cameling)
)

// ToCamelCase converts from underscore separated form to camel case form.
func toCamelCase(s string) string {
	byteSrc := []byte(s)
	chunks := rxCameling.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		chunks[idx] = bytes.Title(val)
	}
	return string(bytes.Join(chunks, nil))
}

// ToSnakeCase converts from camel case form to underscore separated form.
func toSnakeCase(s string) string {
	s = toCamelCase(s)
	runes := []rune(s)
	length := len(runes)
	var out []rune
	for i := 0; i < length; i++ {
		out = append(out, unicode.ToLower(runes[i]))
		if i+1 < length && (unicode.IsUpper(runes[i+1]) && unicode.IsLower(runes[i])) {
			out = append(out, '_')
		}
	}

	return string(out)
}

func ToLowerFirstCamelCase(s string) string {
	if s == "" {
		return s
	}
	if len(s) == 1 {
		return strings.ToLower(s)
	}

	return strings.ToLower(string(s[0])) + toCamelCase(s)[1:]
}

func checkFolder() {

}