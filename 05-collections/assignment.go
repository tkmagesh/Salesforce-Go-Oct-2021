package main

import "strings"

func main() {
	str := "Veniam ipsum laborum cupidatat et proident. Consectetur ipsum et sunt esse. Deserunt adipisicing reprehenderit ad est sit esse. Aliqua eiusmod consectetur culpa officia cillum. Eiusmod ea excepteur reprehenderit ullamco et minim fugiat do consectetur esse ad. Aute qui ad ea sit nisi aliquip excepteur velit pariatur mollit ex. Est officia voluptate minim non laborum. Reprehenderit culpa ullamco enim aliqua sint sit aliquip voluptate magna culpa aute elit. Eiusmod do officia dolor ullamco nisi amet qui do. Adipisicing esse aliqua aute non. Nulla enim amet sint aliquip aliquip amet dolor sit. Magna consectetur nisi nostrud amet anim proident aliquip ullamco."

	/*
		find out the "size of the word" that occurs the most (by size) and print the size and the number of times it occurs
	*/
	/*  Hint: use strings.Split() to split the string to words */
	words := strings.Split(str, " ")
	wordCountBySize := getSordCountBySize(words)
	maxWordSize, maxOccurance := getMaxWordSizeByOccurance(wordCountBySize)
	println(maxWordSize, maxOccurance)
}

func getSordCountBySize(words []string) map[int]int {
	wordCountBySize := make(map[int]int)
	for _, word := range words {
		wordCountBySize[len(word)]++
	}
	return wordCountBySize
}

func getMaxWordSizeByOccurance(wordCountBySize map[int]int) (int, int) {
	maxWordSize := 0
	maxOccurance := 0
	for wordSize, occurance := range wordCountBySize {
		if occurance > maxOccurance {
			maxWordSize = wordSize
			maxOccurance = occurance
		}
	}
	return maxWordSize, maxOccurance
}
