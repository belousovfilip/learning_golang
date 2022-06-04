package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var regWord = regexp.MustCompile(`(?i)[a-zа-я]+[-a-zа-я\d]*`)

func Top10(txt string) []string {
	words := Top(txt)
	if len(words) >= 10 {
		return words[:10]
	}
	return words
}

func Top(txt string) []string {
	wordCount := map[string]int{}
	for _, v := range searchWords(txt, -1) {
		wordCount[strings.ToLower(v)]++
	}
	words := []string{}
	for key := range wordCount {
		words = append(words, key)
	}
	sortByPopularity(words, wordCount)
	return words
}

func sortByPopularity(words []string, wordCount map[string]int) {
	sort.SliceStable(words, func(i, j int) bool {
		if wordCount[words[i]] == wordCount[words[j]] {
			switch strings.Compare(words[j], words[i]) {
			case 1:
				return true
			case -1:
				return false
			case 0:
				return false
			}
		}
		return wordCount[words[i]] > wordCount[words[j]]
	})
}

func searchWords(txt string, count int) []string {
	if len(txt) == 0 {
		return []string{}
	}
	return regWord.FindAllString(txt, count)
}
