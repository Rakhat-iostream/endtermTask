package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func WordEnumerator(out io.Writer) {

	data, err := ioutil.ReadFile("mobydick.txt")

	if err != nil {
		fmt.Println(err)
	}
	//one dimensional array for storing a single word
	var oned []byte
	//kind of sorted slice. In few words, it stores words
	var sortedSlice [][]byte

	size := len(data)

	for i := 0; i < size-1; i++ {
		//checking here whether a byte is a letter or a symbol
		if data[i] >= 97 && data[i] <= 122 || data[i] >= 65 && data[i] <= 90 {
			//and appending only symbols
			oned = append(oned, data[i])
			//if array does not find any letters it means that new word started
			continue
		}
		if len(oned) > 0 {
			//empty array check
			sortedSlice = append(sortedSlice, oned)
		}
		oned = []byte{}
	}
	size = len(sortedSlice)
	//Slice for checked words, reading and counting already checked words cause huge overhead

	var usedWords [][]byte
	var occurrenceSlice []uint
	var index int

	for i := 0; i < size; i++ {
		if usedWords != nil {
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
	fmt.Fprintln(out, "fuck")
}

//function for searching the slice of bytes in the slice of slice of bytes
func isUsedCheck(arr *[][]byte, word *[]byte) int {
	for i := 0; i < len(*arr); i++ {
		if bytes.Compare((*arr)[i], *word) == 0 {
			return i
		}
	}
	return -1
}
