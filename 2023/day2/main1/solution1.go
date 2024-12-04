package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	fmt.Println("\nCollected Lines:")
	hypothesis := map[string]int{
		"blue":  14,
		"red":   12,
		"green": 13,
	}
	for _, line := range lines {
		gameNumber, colorMaps, err := parseGameData(line)
		if err != nil {
			return
		}
		if checkCounts(colorMaps, hypothesis) {
			total += gameNumber
			fmt.Println(line)
			fmt.Println(gameNumber)
		}
	}
	fmt.Println(total)
}

func checkCounts(colorMaps []map[string]int, hypothesis map[string]int) bool {
	for _, colorMap := range colorMaps {
		for color, qty := range colorMap {
			if hypoQty, exists := hypothesis[color]; !exists || qty > hypoQty {
				return false
			}
		}
	}
	return true
}

func parseGameData(input string) (int, []map[string]int, error) {
	parts := strings.SplitN(input, ":", 2)
	if len(parts) != 2 {
		return 0, nil, fmt.Errorf("invalid input format")
	}

	gameNumberStr := strings.TrimSpace(strings.TrimPrefix(parts[0], "Game"))
	gameNumber, err := strconv.Atoi(gameNumberStr)
	if err != nil {
		return 0, nil, fmt.Errorf("error converting game number to int: %v", err)
	}

	data := strings.TrimSpace(parts[1])
	colorBlocks := strings.Split(data, ";")

	var result []map[string]int

	for _, block := range colorBlocks {
		colorQuantities := strings.Split(strings.TrimSpace(block), ",")
		colorMap := make(map[string]int)

		for _, cq := range colorQuantities {
			parts := strings.Split(strings.TrimSpace(cq), " ")
			if len(parts) == 2 {
				quantity, err := strconv.Atoi(parts[0])
				if err != nil {
					return 0, nil, fmt.Errorf("error converting quantity to int: %v", err)
				}
				colorMap[parts[1]] = quantity
			}
		}

		// Ensuring all colors are present in each map
		colors := []string{"blue", "red", "green"}
		for _, color := range colors {
			if _, found := colorMap[color]; !found {
				colorMap[color] = 0
			}
		}
		result = append(result, colorMap)
	}

	return gameNumber, result, nil
}
