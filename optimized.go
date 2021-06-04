package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

//used for collecting word from bytes as slice of bytes
func WordDivider(data *[]byte, sortedSlice *[][]byte) {

	//slice for single word
	var oneDim []byte

	size := len(*data)

	for i := 0; i < size-1; i++ {
		//checking byte for letter
		if (*data)[i] >= 97 && (*data)[i] <= 122 || (*data)[i] >= 65 && (*data)[i] <= 90 {
			//formation of word by byte
			oneDim = append(oneDim, (*data)[i])
			//if no letter, jump to next iteration
			continue
		}
		if len(oneDim) > 0 {
			//empty array check
			*sortedSlice = append(*sortedSlice, oneDim)
		}
		oneDim = []byte{}
	}
}

func WordEnumerator(out io.Writer) {

	data, err := ioutil.ReadFile("mobydick.txt")

	if err != nil {
		fmt.Println(err)
	}
	//slice for storing byte words
	var sortedSlice [][]byte
	//collect words
	WordDivider(&data, &sortedSlice)
	//
	size := len(sortedSlice)

	//slice for storing grouped by occurrence, distinct words
	var usedWords [][]byte
	var occurrenceSlice []uint
	//index of word on a usedWords slice
	var index int

	for i := 0; i < size; i++ {
		if usedWords != nil {
			//is word occurred before or not
			index = isUsedCheck(&usedWords, &sortedSlice[i])
			if index == -1 {
				usedWords = append(usedWords, sortedSlice[i])
				occurrenceSlice = append(occurrenceSlice, 1)
				continue
			}
			occurrenceSlice[index] += 1
			continue
		} else {
			usedWords = append(usedWords, sortedSlice[i])
			occurrenceSlice = append(occurrenceSlice, 1)
		}
	}

	size = len(occurrenceSlice)

	//bubble sort for sorting arrays by occurrence
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size-1; j++ {
			if occurrenceSlice[j] > occurrenceSlice[j+1] {

				temp := occurrenceSlice[j]
				occurrenceSlice[j] = occurrenceSlice[j+1]
				occurrenceSlice[j+1] = temp

				byteSlice := usedWords[j]
				usedWords[j] = usedWords[j+1]
				usedWords[j+1] = byteSlice

			}
		}
	}

	//printing used words
	for i := 0; i < 25; i++ {
		print(string(usedWords[size-i-1]) + " ")
		println(occurrenceSlice[size-i-1])
	}
	fmt.Fprintln(out, "io.Writer")
}

//function for searching the slice of bytes in the slice of slice of bytes
func isUsedCheck(arr *[][]byte, word *[]byte) int {
	for i := 0; i < len(*arr); i++ {
		if bytes.Equal((*arr)[i], *word) {
			return i
		}
	}
	return -1
}
