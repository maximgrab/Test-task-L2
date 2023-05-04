package main

import (
	"sort"
	"strings"
)

func searchAnagrams(words []string) *map[string][]string {
	sort.Strings(words)
	words = uniqeStrings(words)
	result := make(map[string][]string)
	for i := 0; i < len(words)-1; i++ {
		pivot := words[i]
		anagrams := make([]string, 0)
		for j := i + 1; j < len(words); j++ {
			a := words[j]
			if isAnagram(pivot, a) {
				anagrams = append(anagrams, a)
			}
		}
		if len(anagrams) > 1 {
			result[pivot] = anagrams
		}
	}
	return &result
}
func uniqeStrings(strs []string) []string {
	stringSetMap := make(map[string]bool)
	for _, s := range strs {
		stringSetMap[strings.ToLower(s)] = true
	}
	StringSetString := make([]string, 0)
	for key := range stringSetMap {
		StringSetString = append(StringSetString, key)
	}
	return StringSetString
}
func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	as := strings.Split(a, "")
	bs := strings.Split(b, "")
	sort.Strings(as)
	sort.Strings(bs)
	bo := true
	for i := 0; i < len(as)-1; i++ {
		bo = bo && (as[i] == bs[i])
	}
	return bo
}

func main() {

}
