package main

import (
	"leetcode"
	"slices"
	"strings"
)

// https://leetcode.cn/problems/synonymous-sentences/
func generateSentences(synonyms [][]string, text string) []string {
	ans := make([]string, 0)
	sentence := strings.Split(text, " ")
	curr := make([]string, 0)

	parents := make(map[string]string)
	ranks := make(map[string]int)

	findParent := func(a string) string {
		for a != parents[a] {
			parents[a] = parents[parents[a]]
			a = parents[a]
		}
		return a
	}

	merge := func(a string, b string) {
		pa := findParent(a)
		pb := findParent(b)
		if pa == pb {
			return
		}
		if ranks[pa] > ranks[pb] {
			parents[pb] = pa
		} else if ranks[pb] > ranks[pa] {
			parents[pa] = pb
		} else {
			parents[pb] = pa
			ranks[pa]++
		}
	}

	for _, pair := range synonyms {
		w1 := pair[0]
		w2 := pair[1]
		if _, ok := parents[w1]; !ok {
			parents[w1] = w1
			ranks[w1] = 0
		}
		if _, ok := parents[w2]; !ok {
			parents[w2] = w2
			ranks[w1] = 0
		}
		merge(w1, w2)
	}

	word2Conn := make(map[string][]string)
	for wd, pa := range parents {
		p := findParent(pa)
		word2Conn[p] = append(word2Conn[p], wd)
	}
	for wd, conn := range word2Conn {
		slices.Sort(conn)
		for _, w := range conn {
			word2Conn[w] = conn
		}
		word2Conn[wd] = conn
	}

	var dfs func(int)
	dfs = func(i int) {
		if i == len(sentence) {
			ans = append(ans, strings.Join(curr, " "))
			return
		}
		word := sentence[i]
		conn, ok := word2Conn[word]
		if !ok {
			curr = append(curr, word)
			dfs(i + 1)
			curr = curr[:len(curr)-1]
			return
		}
		for _, subNode := range conn {
			curr = append(curr, subNode)
			dfs(i + 1)
			curr = curr[:len(curr)-1]
		}
	}

	dfs(0)
	return ans
}

//func check(synonymsStr string, text string, expect []string) {
//	synonyms := leetcode.Make2DCharArray(synonymsStr)
//	output := generateSentences(synonyms, text)
//	leetcode.Test(slices.Equal(output, expect), "synonyms=%v\n text=%s\n output=%v\n expect=%v", synonyms, text, output, expect)
//}

func check(synonyms [][]string, text string, expect []string) {
	output := generateSentences(synonyms, text)
	leetcode.Test(slices.Equal(output, expect), "synonyms=%v\n text=%s\n output=%v\n expect=%v", synonyms, text, output, expect)
}

func main() {
	//check([][]string{{"happy", "joy"}, {"sad", "sorrow"}, {"joy", "cheerful"}}, "I am happy today but was sad yesterday", []string{"I am cheerful today but was sad yesterday", "I am cheerful today but was sorrow yesterday", "I am happy today but was sad yesterday", "I am happy today but was sorrow yesterday", "I am joy today but was sad yesterday", "I am joy today but was sorrow yesterday"})
	//check("[[\"a\",\"b\"},{\"b\",\"c\"},{\"d\",\"e\"},{\"c\",\"d\"]]", "a b", []string{"a a", "a b", "a c", "a d", "a e", "b a", "b b", "b c", "b d", "b e", "c a", "c b", "c c", "c d", "c e", "d a", "d b", "d c", "d d", "d e", "e a", "e b", "e c", "e d", "e e"})
	check([][]string{{"a", "b"}, {"b", "c"}, {"d", "e"}, {"c", "d"}}, "a b", []string{"a a", "a b", "a c", "a d", "a e", "b a", "b b", "b c", "b d", "b e", "c a", "c b", "c c", "c d", "c e", "d a", "d b", "d c", "d d", "d e", "e a", "e b", "e c", "e d", "e e"})
}
