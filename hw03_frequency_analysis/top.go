package hw03frequencyanalysis

import (
	"sort"
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

	words := strings.Fields(s)

	m := make(map[string]int)
	for _, word := range words {
		if word == "" {
			continue
		}
		m[word] = m[word] + 1
	}
	sl := wordsToSlice(m)

	sort.Slice(sl, func(i, j int) bool {
		if sl[i].fr == sl[j].fr {
			return sl[i].w < sl[j].w
		} else {
			return sl[i].fr > sl[j].fr
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
