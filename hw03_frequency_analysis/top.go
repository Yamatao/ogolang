package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
	"unicode"
)

type countedItems struct {
	v      []string
	counts map[string]int
}

func (v countedItems) Len() int {
	return len(v.v)
}

func (v countedItems) Less(i, j int) bool {
	return v.counts[v.v[i]] > v.counts[v.v[j]] // inverted
}

func (v countedItems) Swap(i, j int) {
	v.v[i], v.v[j] = v.v[j], v.v[i]
}

// Top10 returns the most frequent words in the text, but not more than 10.
func Top10(src string) []string {
	counts := make(map[string]int, 32)
	lastIsLetter := false
	var wordb strings.Builder
	wordb.Grow(32)
	const inwordChar = '-'

	// split the text to words and put them to the `counts` map
	for _, r := range src {
		isLetter := unicode.IsLetter(r)
		if isLetter {
			wordb.WriteRune(unicode.ToLower(r))
		} else if lastIsLetter {
			if r != inwordChar {
				// add the word to the `counts` map
				counts[wordb.String()]++
				wordb.Reset()
			} else {
				wordb.WriteRune(r)
				isLetter = true // '-' is considered as word's part
			}
		}
		lastIsLetter = isLetter
	}
	if wordb.Len() > 0 {
		counts[wordb.String()]++
	}

	// sort to have the most popular items first
	result := make([]string, 0, len(counts))
	for word := range counts {
		result = append(result, word)
	}
	items := countedItems{result, counts}
	sort.Sort(items)

	// trim the rest but 10
	if len(result) > 10 {
		result = result[:10]
	}
	return result
}
