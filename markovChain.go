package main

import (
	"fmt"
	"strconv"
	"strings"
)

type MarkovChain struct {
	value     string
	count     int
	prob      float64
	nextChain []MarkovChain
}

func addWordToChain(chain []MarkovChain, word string) ([]MarkovChain, int) {
	index := -1
	for i := 0; i < len(chain); i++ {
		if chain[i].value == word {
			index = i
		}
	}
	if index >= 0 {
		chain[index].count++
	} else {
		var tempChain MarkovChain
		tempChain.value = word
		tempChain.count = 1

		chain = append(chain, tempChain)
		index = len(chain) - 1
	}
	return chain, index
}

func generateMarkovChain(words []string) []MarkovChain {
	var chain []MarkovChain
	var index int

	for i := 0; i < len(words)-1; i++ {

		chain, index = addWordToChain(chain, words[i])
		if index < len(words){
			chain[index].nextChain, _ = addWordToChain(chain[index].nextChain, words[i+1])
		}
		printLoading(i, len(words))
	}

	for i := 0; i < len(chain); i++{
		chain[i].prob = float64(chain[i].count)/float64(len(words))*100
		for j := 0; j < len(chain[i].nextChain); j++ {
			chain[i].nextChain[j].prob = float64(chain[i].nextChain[j].count)/float64(len(words))*100
		}
	}

	return chain
}

func printLoading(n int, total int) {
	var bar []string
	tantPerFourty := int((float64(n) / float64(total)) * 40)
	tantPerCent := int((float64(n) / float64(total)) * 100)
	for i := 0; i < tantPerFourty; i++ {
		bar = append(bar, "█")
	}
	progressBar := strings.Join(bar, "")
	fmt.Printf("\r " + progressBar + " - " + strconv.Itoa(tantPerCent) + "")
}

func getNextWord(chain []MarkovChain, word string)string{
	index := -1

	for i := 0; i < len(chain); i++ {
		if chain[i].value == word{
			index = i
		}
	}

	if index < 0 {
		return "word doesn't exist in chain"
	}
	var nextWord MarkovChain
	if chain[index].nextChain != nil{
		nextWord = chain[index].nextChain[0]
		for i := 0; i < len(chain[index].nextChain); i++ {
			if chain[index].nextChain[i].prob > nextWord.prob {
				nextWord = chain[index].nextChain[i]
			}
		}
	}

	return nextWord.value
}

func Train(text []string)[]MarkovChain{
	chain := generateMarkovChain(text)
	return chain
}

func GenerateMarkovText(chain []MarkovChain, startWord string, number int) string{
	var generatedText []string
	word := startWord

	generatedText = append(generatedText, word)

	for i := 0; i < number; i++ {
		word = getNextWord(chain, word)
		if word == "word doesn't exist in chain" {
			return "word doesn't exist in chain"
		}
		generatedText = append(generatedText, word)
	}

	return strings.Join(append(generatedText, "."), " ")
}




