package main

import (
	"bytes"
	_ "fmt"
	"io"
	"io/ioutil"
)

/* Написал как Case Sensitive
 */
func ReadFile(io.Writer) {
	data, err := ioutil.ReadFile("mobydick.txt")
	if err != nil {
		err.Error()
	}
	var oned []byte
	var sortedSlice [][]byte
	for i := 0; i < len(data); i++ {
		if data[i] >= 97 && data[i] <= 122 || data[i] >= 65 && data[i] <= 90 {
			oned = append(oned, data[i])
			if i == len(data)-1 {
				sortedSlice = append(sortedSlice, oned)
			}
		} else {
			var newd []byte
			if len(oned) > 0 {
				sortedSlice = append(sortedSlice, oned)
			}
			oned = newd
		}
	}
	size := len(sortedSlice)
	//Slice for checked words, reading and counting already checked words cause huge overhead
	var usedWords [][]byte
	var occurrenceSlice []uint
	var index int
	usedWords = append(usedWords, sortedSlice[0])
	occurrenceSlice = append(occurrenceSlice, 1)
	for i := 1; i < size; i++ {
		index = isUsed(&usedWords, &sortedSlice[i])
		if index == -1 {
			usedWords = append(usedWords, sortedSlice[i])
			occurrenceSlice = append(occurrenceSlice, 1)
		} else {
			occurrenceSlice[index] += 1
		}
	}
	size = len(occurrenceSlice)
	sort(&usedWords, &occurrenceSlice, 0, size-1)
	//print
	for i := 0; i < size; i++ {
		if i == 26 {
			break
		}
		print(string(usedWords[size-i-1]) + " ")
		println(occurrenceSlice[size-i-1])
	}
}
func sort(cache *[][]byte, occurrences *[]uint, first int, last int) {
	left, right := first, last
	pivot := (*occurrences)[(left+right)/2]
	for left <= right {
		for (*occurrences)[left] < pivot {
			left++
		}
		for (*occurrences)[right] > pivot {
			right--
		}
		if left <= right {
			(*occurrences)[left], (*occurrences)[right] = (*occurrences)[right], (*occurrences)[left]
			(*cache)[left], (*cache)[right] = (*cache)[right], (*cache)[left]
			left++
			right--
		}
	}
	if first < right {
		sort(cache, occurrences, first, right)
	}
	if left < last {
		sort(cache, occurrences, left, last)
	}
}

//function for searching the slice of bytes in the slice of slice of bytes
func isUsed(arr *[][]byte, word *[]byte) int {
	for i := 0; i < len(*arr); i++ {
		if bytes.Equal((*arr)[i], *word) == true {
			return i
		}
	}
	return -1
}
