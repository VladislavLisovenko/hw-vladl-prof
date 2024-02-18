package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func stringIsVaid(s string, checks []*regexp.Regexp) bool {
	valid := true
	if s == "" {
		return valid
	}

	for _, re := range checks {
		if re.MatchString(s) {
			valid = false
			break
		}
	}
	return valid
}

func isDigit(c rune) bool {
	return c > 47 && c < 58
}

func isLetter(c rune) bool {
	return !isBackslash(c) && !isDigit(c)
}

func isBackslash(c rune) bool {
	return c == 92
}

func checks() []*regexp.Regexp {
	checks := []*regexp.Regexp{}

	re := regexp.MustCompile("^[0-9]")
	checks = append(checks, re)

	re = regexp.MustCompile(`[^\\][0-9]{2,}`)
	checks = append(checks, re)

	return checks
}

func Unpack(s string) (string, error) {
	checks := checks()
	if !stringIsVaid(s, checks) {
		return "", ErrInvalidString
	}
	symbols := make([]string, 0)
	repeaters := make([]int, 0)
	prevR := ``
	for _, r := range s {
		switch {
		case isLetter(r):
			if prevR == `\` {
				symbols = append(symbols, `\`+string(r))
				prevR = ``
			} else {
				symbols = append(symbols, string(r))
			}
			repeaters = append(repeaters, 1)
		case isDigit(r):
			d, _ := strconv.Atoi(string(r))
			repeaters[len(repeaters)-1] = d
		default:
			prevR = `\`
		}
	}

	sb := &strings.Builder{}
	for i, v := range symbols {
		sb.WriteString(strings.Repeat(v, repeaters[i]))
	}

	return sb.String(), nil
}
