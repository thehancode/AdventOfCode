package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Converter struct {
	source      string
	destination string
	maps        []Mapping
}

type Mapping struct {
	destiantionStart int
	sourceStart      int
	rangeSize        int
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

	var currentConverter *Converter
	var converters []Converter
	var seeds []int

	for _, line := range lines {
		if len(line) <= 1 {
			continue
		}
		if strings.HasPrefix(line, "seeds:") {
			seeds, _ = getSeeds(line)
			continue
		}
		if !unicode.IsDigit(rune(line[0])) {
			if currentConverter != nil {
				converters = append(converters, *currentConverter)
			}
			currentSource, currentDestination := getNames(line)

			currentConverter = &Converter{destination: currentDestination, source: currentSource}
			converters = append(converters, *currentConverter)
			continue
		}
		currentDestinationStart, currentSourceStart, currentRange := getMapping(line)
		currentMap := Mapping{destiantionStart: currentDestinationStart, sourceStart: currentSourceStart, rangeSize: currentRange}

		currentConverter.maps = append(currentConverter.maps, currentMap)

	}
	converters = append(converters, *currentConverter)
	minLocation := math.MaxInt32
	fmt.Printf("seeds  : %#v", seeds)
	fmt.Printf("conveters  : %#v", converters)

	for _, seed := range seeds {
		convertedSeed := seed
		fmt.Print(" %d", seed)
		for _, converter := range converters {
			convertedSeed = convert(convertedSeed, converter)
			fmt.Print(" %d", convertedSeed)
		}
		println(":")
		println(convertedSeed)
		if minLocation > convertedSeed {
			minLocation = convertedSeed
		}
	}
	println(minLocation)

}

func getSeeds(line string) ([]int, error) {
	var seeds []int
	numbersStr := strings.TrimPrefix(line, "seeds:")
	numbersStr = strings.TrimSpace(numbersStr)
	numberStrs := strings.Fields(numbersStr)
	for _, numStr := range numberStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("failed to convert '%s' to int: %v", numStr, err)
		}
		seeds = append(seeds, num)
	}

	return seeds, nil
}

func getNames(line string) (string, string) {
	parts := strings.Split(line, "-to-")
	if len(parts) < 2 {
		return "", ""
	}
	firstName := parts[0]
	secondParts := strings.Fields(parts[1])
	secondName := secondParts[0]
	return firstName, secondName
}

func getMapping(line string) (int, int, int) {
	var a, b, c int
	fmt.Sscanf(line, "%d %d %d", &a, &b, &c)
	return a, b, c
}

func convert(seed int, converter Converter) int {
	convertedSeed := seed
	for _, mapping := range converter.maps {
		if mapping.sourceStart <= seed && seed < mapping.sourceStart+mapping.rangeSize {
			convertedSeed = mapping.destiantionStart + seed - mapping.sourceStart
		}
	}
	return convertedSeed
}
