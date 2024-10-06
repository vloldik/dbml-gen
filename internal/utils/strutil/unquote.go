package strutil

import (
	"strconv"
	"strings"
)

func RemoveQuotes(s, quote string) (string, bool) {
	if strings.HasPrefix(s, quote) && strings.HasSuffix(s, quote) {
		s = strings.TrimPrefix(s, quote)
		s = strings.TrimSuffix(s, quote)
		return s, true
	}

	return s, false
}

func TryUnquote(s string) string {
	if !strings.HasPrefix(s, "\"") && !strings.HasPrefix(s, "'") && !strings.HasPrefix(s, "`") {
		return s
	}
	unquoted, err := UnquoteString(s)
	if err != nil {
		return s
	}
	return unquoted
}

func UnquoteString(s string) (string, error) {
	if len(s) < 2 {
		return s, nil
	}
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
