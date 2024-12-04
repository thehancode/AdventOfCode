package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Card struct {
	cardId      int
	winingCount int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return
	}

	total := 0
	var matchingCountCards []int

	for _, line := range lines {
		_, myNumbers, winnerNumbers, _ := parseLine(line)
		countMyWinnerNumbers := countEqualElements(myNumbers, winnerNumbers)
		matchingCountCards = append(matchingCountCards, countMyWinnerNumbers)
	}
	tot := 0
	for _, matchCount := range matchingCountCards {
		tot += int(math.Exp2(float64(matchCount - 1)))
	}

	fmt.Printf("tot : %d", tot)
	fmt.Printf("%#v", matchingCountCards)
	totalCountCards := make([]int, len(matchingCountCards))
	for i := range totalCountCards {
		totalCountCards[i] = 1
	}
	for i := len(matchingCountCards) - 1; i >= 0; i-- {
		totalCountCards[i] = 1
		for j := 1; j < matchingCountCards[i]+1; j++ {

			totalCountCards[i] += totalCountCards[i+j]
		}
	}
	fmt.Printf("Result: %#v", totalCountCards)
	total = sumIntSlice(totalCountCards)
	fmt.Printf("Total : %d", total)
}
func sumIntSlice(slice []int) int {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return sum
}

func parseLine(line string) (cardNumber string, array1, array2 []string, err error) {
	parts := strings.Split(line, "|")
	if len(parts) != 2 {
		return "", nil, nil, fmt.Errorf("invalid format, expected '|' separator")
	}

	cardParts := strings.SplitN(parts[0], ":", 2)
	if len(cardParts) != 2 {
		return "", nil, nil, fmt.Errorf("invalid card number format")
	}
	cardNumber = strings.TrimSpace(strings.Split(cardParts[0], " ")[1])
	array1 = strings.Fields(strings.TrimSpace(cardParts[1]))
	array2 = strings.Fields(strings.TrimSpace(parts[1]))
	return cardNumber, array1, array2, nil
}

func countEqualElements(array1, array2 []string) int {
	elementMap := make(map[string]bool)
	for _, elem := range array2 {
		elementMap[elem] = true
	}

	count := 0
	for _, elem := range array1 {
		if _, exists := elementMap[elem]; exists {
			count++
		}
	}

	return count
}
