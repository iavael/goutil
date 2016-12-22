package goutil

import (
	"bytes"
)

// Concat slice of strings in one string
func strfmtBase(args []string) *bytes.Buffer {
	result := bytes.NewBufferString("")
	for _, arg := range args {
		result.WriteString(arg)
	}
	return result
}

func StrCat(args ...string) string {
	result := strfmtBase(args)
	return result.String()
}

func StrCatln(args ...string) string {
	result := strfmtBase(args)
	result.WriteRune('\n')
	return result.String()
}

func StrCatS(args []string) string {
	result := strfmtBase(args)
	return result.String()
}

func StrCatSln(args []string) string {
	result := strfmtBase(args)
	result.WriteRune('\n')
	return result.String()
}
