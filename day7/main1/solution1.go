package main

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/maps"
	"os"
	"reflect"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	value string
	bid   int
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
	var hands []Hand
	for _, line := range lines {
		parsedHand, err := parseLine(line)
		if err != nil {
			fmt.Printf("Error parsing line '%s': \n", line)
			continue
		}
		hands = append(hands, parsedHand)
	}
	slices.SortFunc(hands, compareHands)
	fmt.Printf("Sorted %#v \n", hands)

	ans := 0
	for i, hand := range hands {
		ans += (i + 1) * hand.bid
	}
	fmt.Printf("ans: %d", ans)
}

func compareHands(handA, handB Hand) int {
	typeA := getHandType(handA)
	fmt.Printf("%d %s\n", typeA, handA.value)
	typeB := getHandType(handB)
	fmt.Printf("%d %s\n", typeB, handB.value)
	if typeA > typeB {
		return 1
	}
	if typeA < typeB {
		return -1
	}
	return compareEqual(handA.value, handB.value)
}
func getHandType(hand Hand) int {

	handCount := countLetters(hand.value)
	jCount := handCount[rune('J')]

	if jCount == 5 {
		return 7
	}
	delete(handCount, rune('J'))
	sortedValues := maps.Values(handCount)
	sort.Ints(sortedValues)
	sortedValues[len(sortedValues)-1] += jCount

	if reflect.DeepEqual(sortedValues, []int{5}) {
		return 7
	}
	if reflect.DeepEqual(sortedValues, []int{1, 4}) {
		return 6
	}
	if reflect.DeepEqual(sortedValues, []int{2, 3}) {
		return 5
	}
	if reflect.DeepEqual(sortedValues, []int{1, 1, 3}) {
		return 4
	}
	if reflect.DeepEqual(sortedValues, []int{1, 2, 2}) {
		return 3
	}
	if reflect.DeepEqual(sortedValues, []int{1, 1, 1, 2}) {
		return 2
	}
	if reflect.DeepEqual(sortedValues, []int{1, 1, 1, 1, 1}) {
		return 1
	}
	fmt.Printf("Count %#v \n", handCount)
	fmt.Printf("Sorted %#v \n", sortedValues)
	fmt.Printf("Error not type match %s \n", hand.value)

	return 0
}

func countLetters(s string) map[rune]int {
	count := make(map[rune]int)
	for _, r := range s {
		count[r]++
	}
	return count
}

func compareEqual(handA, handB string) int {
	var cardStrength = map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 1,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}
	for i := range handA {
		aStrength, aOk := cardStrength[rune(handA[i])]
		bStrength, bOk := cardStrength[rune(handB[i])]
		if !aOk || !bOk {
			fmt.Println("Invalid card found")
			return -1
		}
		if aStrength > bStrength {
			return 1
		} else if aStrength < bStrength {
			return -1
		}
	}
	return 0
}
func parseLine(line string) (Hand, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return Hand{}, fmt.Errorf("invalid line format")

	}

	bid, err := strconv.Atoi(parts[1])
	if err != nil {
		return Hand{}, fmt.Errorf("invalid bid value")
	}

	return Hand{
		value: parts[0],
		bid:   bid,
	}, nil
}
