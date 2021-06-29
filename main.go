package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){

	cmdReader := bufio.NewReader(os.Stdin)
	options :=`
	1 - Generate text
	2 - Exit
			`

	for {
		fmt.Println(options)

		option, _ := cmdReader.ReadString('\n')

		option = strings.TrimSpace(option)

		switch option{
			case "1":
				content, err := ReadText("example.pdf")
				if err != nil {
					fmt.Printf("Error occurred with readText: %v\n", err)
				}
				chain := Train(content)
				start(chain)
				//printChain(chain)
				break
			case "2":
				return
				break
			default:
				fmt.Println("Wrong option")

		}

	}
}

func start(chain []MarkovChain){
	cmdReader := bufio.NewReader(os.Stdin)

	fmt.Print("\nEnter word from which will start the generation - ")
	word, _ := cmdReader.ReadString('\n')
	word = strings.TrimSpace(word)

	fmt.Print("Enter the count of word in the chain - ")
	answer, _ := cmdReader.ReadString('\n')
	answer = strings.TrimSpace(answer)

	count, _ := strconv.Atoi(answer)

	text := GenerateMarkovText(chain, word, count)
	fmt.Println("Result - ", text)
}

func printChain(chain []MarkovChain){
	for i := 0; i < len(chain); i++ {
		fmt.Println(i+1," - ",chain[i].value, " prob - ", chain[i].prob)
		if chain[i].nextChain != nil {
			printChain(chain[i].nextChain)
		}
	}
}
