package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Pair struct {
	Left  string
	Right string
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
	nodes, starters := parseLines(lines)
	instructions := lines[0]
	currents := starters
	var step rune
	steps := 0
	idx := 0
	lengths := make([]int, len(starters))

	fmt.Printf(" star  : %#v  \n", currents)
	for notAllEmpty(currents) {
		//fmt.Printf(" curernts : %#v  \n", currents)
		step = rune(instructions[idx])
		for i, current := range currents {
			if current == "" {
				continue
			}
			if step == 'L' {
				currents[i] = nodes[current].Left
			}
			if step == 'R' {
				currents[i] = nodes[current].Right
			}
			if currents[i][len(currents[i])-1] == 'Z' {
				lengths[i] = steps + 1
				currents[i] = ""
			}

		}
		idx = (idx + 1) % len(instructions)
		steps++
	}

	fmt.Printf(" lengths : %#v  \n", lengths)

	lcm := lcm(lengths)
	fmt.Printf("ans: %d \n", lcm)
	fmt.Printf("steps: %d \n", lcm)
}

func notAllEmpty(nodes []string) bool {
	for _, node := range nodes {
		if node != "" {
			return true
		}
	}
	return false

}

func parseLines(lines []string) (map[string]Pair, []string) {
	pairs := make(map[string]Pair)
	var starters []string
	re := regexp.MustCompile(`(\w+)\s*=\s*\((\w+),\s*(\w+)\)`)
	for i, line := range lines {
		if i == 0 {
			continue
		}
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			if matches[1][len(matches[1])-1] == 'A' {
				starters = append(starters, matches[1])
			}
			pairs[matches[1]] = Pair{
				Left:  matches[2],
				Right: matches[3],
			}
		}
	}
	return pairs, starters
}
func lcm(nums []int) int64 {
	lcm := int64(1)
	for _, num := range nums {
		num64 := int64(num)
		lcm = lcm * (num64 / gcd(lcm, num64))
	}
	return lcm
}
func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
