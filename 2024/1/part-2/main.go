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
	listB := map[int]int{}
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

		listA = append(listA, itemA)
		listB[itemB] += 1
	}

	score := 0
	for i := 0; i < len(listA); i += 1 {
		itemA := listA[i]
		itemACountInB := listB[itemA]
		finalItemA := itemA * itemACountInB
		score += finalItemA
	}

	fmt.Printf("The similarity score is %d", score)
}
