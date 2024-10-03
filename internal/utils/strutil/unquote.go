package strutil

import (
	"strconv"
	"strings"
)

func UnquoteString(s string) (string, error) {
	quote := s[0]
	s = s[1 : len(s)-1]
	out := ""
	for s != "" {
		if strings.HasPrefix(s, `\`) {
			s = removeNewline(s)
		}
		value, _, tail, err := strconv.UnquoteChar(s, quote)
		if err != nil {
			return "", err
		}
		s = tail
		out += string(value)
	}
	return out, nil
}

func removeNewline(s string) string {
	if s[1:3] == "\r\n" {
		return s[3:]
	}
	return s
}
