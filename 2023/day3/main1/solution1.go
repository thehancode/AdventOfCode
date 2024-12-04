package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	numberMap, symbolMap := createNumberMap(lines)
	for key, value := range numberMap {
		fmt.Printf("Key: (%d, %d), Value: %+v\n", key.first, key.second, value)
	}
	for key, value := range symbolMap {
		fmt.Printf("Key: (%d, %d), Value: %+v\n", key.first, key.second, value)
	}
	directions := []intPair{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1} /* {0, 0}, */, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	total := 0
	counted := make(map[int]bool)
	for pair, _ := range symbolMap {
		fmt.Println("")
		for _, direction := range directions {
			key := intPair{first: pair.first + direction.first, second: pair.second + direction.second}
			fmt.Println("checking point %d %d", key.first, key.second)
			num := numberMap[key]
			if !counted[num.ID] {
				counted[num.ID] = true
				numValue, _ := strconv.Atoi(num.Value)
				total += numValue
				fmt.Println("adding  : %d", numValue)
			}

		}
	}
	fmt.Println(total)

}

func createNumberMap(schematic []string) (map[intPair]Number, map[intPair]int) {
	numberMap := make(map[intPair]Number)
	symbolMap := make(map[intPair]int)
	currentId := 0
	symbolId := 0
	for j := 0; j < len(schematic); j++ {
		for i := 0; i < len(schematic[j]); i++ {
			char := schematic[j][i]
			fmt.Println(fmt.Sprintf("%d %d %s", i, j, char))
			if char == '.' {
				continue
			}
			if char >= '0' && char <= '9' {
				currentNumber := ""
				numberLen := 0
				currentChar := schematic[j][i+numberLen]
				for currentChar >= '0' && currentChar <= '9' && i+numberLen < len(schematic[i]) {
					currentNumber += string(currentChar)
					numberLen += 1
					if i+numberLen < len(schematic[i]) {
						currentChar = schematic[j][i+numberLen]
					}
				}

				fmt.Println(fmt.Sprintf("%d %d %s %d %s", i, j, char, numberLen, currentNumber))
				for k := i; k < i+numberLen; k++ {
					key := intPair{first: j, second: k}
					value := Number{ID: currentId, Value: currentNumber}
					numberMap[key] = value
				}
				currentId += 1
				i = i + numberLen - 1
			} else {
				fmt.Println(symbolId)
				key := intPair{first: j, second: i}
				symbolMap[key] = symbolId
				symbolId += 1
			}

		}
	}
	return numberMap, symbolMap
}
