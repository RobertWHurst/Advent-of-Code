package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const listsSourceFilePath = "numbers.txt"
const lineSep = "\n"
const listSep = "   "

func main() {
	listSrc, err := os.ReadFile(listsSourceFilePath)
	if err != nil {
		panic("Failed to read source file")
	}

	listAAndBLines := strings.Split(string(listSrc), lineSep)

	var listA []int
	var listB []int
	for _, listAAndBLine := range listAAndBLines {
		if listAAndBLine == "" {
			break
		}

		listAAndBItems := strings.Split(listAAndBLine, listSep)
		listAAndBItemsLen := len(listAAndBItems)
		if listAAndBItemsLen != 2 {
			panic(fmt.Sprintf(
				"each line must contain a number for list A and a second number for "+
					"list B separated by '%s' but %d were found", lineSep, listAAndBItemsLen))
		}

		itemAStr := listAAndBItems[0]
		itemBStr := listAAndBItems[1]

		itemA, err := strconv.Atoi(itemAStr)
		if err != nil {
			panic("Failed to parse item from list A '" + itemAStr + "'")
		}
		if itemA < 0 {
			panic("Item from list A is unexpectedly negative '" + itemAStr + "'")
		}

		itemB, err := strconv.Atoi(itemBStr)
		if err != nil {
			panic("Failed to parse item from list B '" + itemBStr + "'")
		}
		if itemB < 0 {
			panic("Item from list B is unexpectedly negative '" + itemBStr + "'")
		}

		listA = bisectInsertSorted(listA, itemA)
		listB = bisectInsertSorted(listB, itemB)
	}

	totalDist := 0
	for i := 0; i < len(listA); i += 1 {
		itemA := listA[i]
		itemB := listB[i]
		if itemA > itemB {
			totalDist += itemA - itemB
		} else {
			totalDist += itemB - itemA
		}
	}

	fmt.Printf("The total distance between the lists is %d", totalDist)
}

func bisectInsertSorted(sortedList []int, item int) []int {
	maxIndex := len(sortedList) - 1
	if maxIndex == -1 {
		return []int{item}
	}

	minItem := sortedList[0]
	maxItem := sortedList[maxIndex]
	if item < minItem {
		return append([]int{item}, sortedList...)
	}
	if item > maxItem {
		return append(sortedList, item)
	}

	offset := maxIndex
	cursor := 0
	multiplicate := 1
	for {
		offset /= 2
		cursor += (offset * multiplicate)

		if sortedList[cursor] == item {
			return append(sortedList[:cursor], append([]int{item}, sortedList[cursor:]...)...)
		}

		if offset == 0 {
			for ; ; cursor += multiplicate {
				switch {
				case cursor > maxIndex:
					return append(sortedList, item)
				case cursor < 1:
					return append([]int{item}, sortedList...)
				case multiplicate == 1 && sortedList[cursor+1] > item:
					return append(sortedList[:cursor+1], append([]int{item}, sortedList[cursor+1:]...)...)
				case multiplicate == -1 && sortedList[cursor-1] < item:
					return append(sortedList[:cursor], append([]int{item}, sortedList[cursor:]...)...)
				}
			}
		}

		if sortedList[cursor] > item {
			multiplicate = -1
		} else {
			multiplicate = 1
		}
	}
}
