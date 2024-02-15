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
	nodes := parseLines(lines)
	instructions := lines[0]
	current := "AAA"
	var step rune
	steps := 0
	idx := 0
	for current != "ZZZ" {
		step = rune(instructions[idx])
		if step == 'L' {
			current = nodes[current].Left
		}
		if step == 'R' {
			current = nodes[current].Right
		}
		idx = (idx + 1) % len(instructions)
		steps++
	}
	fmt.Printf("ans: %d", steps)
}

func parseLines(lines []string) map[string]Pair {
	pairs := make(map[string]Pair)
	re := regexp.MustCompile(`(\w+)\s*=\s*\((\w+),\s*(\w+)\)`)
	for i, line := range lines {
		if i == 0 {
			continue
		}
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			pairs[matches[1]] = Pair{
				Left:  matches[2],
				Right: matches[3],
			}
		}
	}
	return pairs
}
