package main

import (
	"bytes"
	_ "fmt"
	"io/ioutil"
)

func main() {

	data, err := ioutil.ReadFile("mobydick.txt")
	if err != nil {
		err.Error()
	}
	var oneDim []byte
	var sortedSlice [][]byte

	for i := 0; i < len(data); i++ {
		if data[i] >= 97 && data[i] <= 122 || data[i] >= 65 && data[i] <= 90 {
			oneDim = append(oneDim, data[i])
			if i == len(data)-1 {
				sortedSlice = append(sortedSlice, oneDim)
			}
		} else {
			var newd []byte
			if len(oneDim) > 0 {
				sortedSlice = append(sortedSlice, oneDim)
			}
			oneDim = newd
		}
	}
	size := len(sortedSlice)
	//Slice for checked words, reading and counting already checked words cause huge overhead
	var usedWords [][]byte
	var occurrenceSlice []uint

	for i := 0; i < size; i++ {
		var counter uint
		if usedWords != nil {
			if used(usedWords, sortedSlice[i]) {
				continue
			}
		}
		for j := 0; j < size; j++ {
			if bytes.Compare(sortedSlice[i], sortedSlice[j]) == 0 {
				counter++
			}
		}
		usedWords = append(usedWords, sortedSlice[i])
		occurrenceSlice = append(occurrenceSlice, counter)
	}

	size = len(occurrenceSlice)

	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
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

	for i := 0; i < size; i++ {
		if i == 26 {
			break
		}
		print(string(usedWords[size-i-1]) + " ")
		println(occurrenceSlice[size-i-1])
	}
}

//function for searching the slice of bytes in the slice of slice of bytes
func used(arr [][]byte, word []byte) bool {
	for i := 0; i < len(arr); i++ {
		if bytes.Compare(arr[i], word) == 0 {
			return true
		}
	}
	return false
}
