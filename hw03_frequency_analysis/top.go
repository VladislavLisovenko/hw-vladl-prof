package hw03frequencyanalysis

import (
	"slices"
	"strings"
)

type WordFrequency struct {
	w  string
	fr int
}

func wordsToSlice(m map[string]int) []WordFrequency {
	res := make([]WordFrequency, 0)
	for k, v := range m {
		res = append(res, WordFrequency{
			w:  k,
			fr: v,
		})
	}

	return res
}

func Top10(s string) []string {
	if s == "" {
		return []string{}
	}

	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", " ")

	words := strings.Split(s, " ")

	m := make(map[string]int)
	for _, word := range words {
		if word == "" {
			continue
		}
		m[word] = m[word] + 1
	}
	sl := wordsToSlice(m)

	slices.SortFunc[[]WordFrequency](sl, func(a, b WordFrequency) int {
		if a.fr == b.fr {
			return strings.Compare(a.w, b.w)
		}
		switch {
		case a.fr > b.fr:
			return -1
		case a.fr < b.fr:
			return 1
		default:
			return 0
		}
	})

	res := []string{}
	for i, v := range sl {
		if i == 10 {
			break
		}
		res = append(res, v.w)
	}

	return res
}
