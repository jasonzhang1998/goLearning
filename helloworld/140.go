package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello")
	s := "catsanddog"
	dict := []string{"cat", "cats", "and", "sand", "dog"}
	res := wordBreak(s, dict)
	for _, sentence := range res {
		fmt.Println(sentence)
	}
}

func wordBreak(s string, wordDict []string) (sentences []string) {
	wordSet := map[string]struct{}{}
	for _, w := range wordDict {
		wordSet[w] = struct{}{}
	}

	n := len(s)
	dp := make([][][]string, n)
	var backtrack func(index int) [][]string
	backtrack = func(index int) [][]string {
		if dp[index] != nil {
			return dp[index]
		}

		wordslist := [][]string{}
		for i := index + 1; i < n; i++ {
			word := s[index:i]
			if _, has := wordSet[word]; has {
				for _, nextwords := range backtrack(i) {
					wordslist = append(wordslist, append([]string{word}, nextwords...))
				}
			}
		}

		word := s[index:]
		if _, has := wordSet[word]; has {
			wordslist = append(wordslist, []string{word})
		}
		dp[index] = wordslist

		return wordslist
	}

	for _, words := range backtrack(0) {
		sentences = append(sentences, strings.Join(words, " "))
	}

	return
}
