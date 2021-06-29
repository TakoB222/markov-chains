package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadText(filePath string)([]string, error){
	rawContent, err := ioutil.ReadFile(filePath)
	if err != nil{
		return nil, err
	}

	contentWords :=  strings.Split(strings.ToLower(strings.Replace(string(rawContent), "\n", " ", -1)), " ")
	fmt.Println("Total words: ", len(contentWords))

	return contentWords, nil
}
