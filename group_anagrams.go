package main

func groupAnagrams(words []string) [][]string {

	items := make(map[[26]int][]string)
	for _, word := range words {
		var key [26]int
		for i := 0; i < len(word); i++ {
			key[word[i]-'a']++
		}
		items[key] = append(items[key], word)
	}
	var valids [][]string
	for _, word := range items {
		valids = append(valids, word)
	}

	return valids
}
