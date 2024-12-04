package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Number struct {
	ID    int
	Value string
}

type intPair struct {
	first  int
	second int
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

	total := 0.0

	for _, line := range lines {
		_, myNumbers, winnerNumbers, _ := parseLine(line)
		countMyWinnerNumbers := countEqualElements(myNumbers, winnerNumbers)
		fmt.Println(countMyWinnerNumbers)
		if countMyWinnerNumbers > 0 {
			total += math.Exp2(float64(countMyWinnerNumbers - 1))
		}
		fmt.Println("total  :%d", total)
	}
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
