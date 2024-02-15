package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

type Pair struct {
	Left  string
	Right string
}
type Direction struct {
	EntryDirection [][2]int
	OutDirection   [][2]int
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
	bends := map[rune]Direction{
		'L': {
			EntryDirection: [][2]int{{1, 0}, {0, -1}},
			OutDirection:   [][2]int{{0, 1}, {-1, 0}},
		},
		'J': {
			EntryDirection: [][2]int{{1, 0}, {0, 1}},
			OutDirection:   [][2]int{{0, -1}, {-1, 0}},
		},
		'7': {
			EntryDirection: [][2]int{{-1, 0}, {0, 1}},
			OutDirection:   [][2]int{{0, -1}, {1, 0}},
		},
		'F': {
			EntryDirection: [][2]int{{-1, 0}, {0, -1}},
			OutDirection:   [][2]int{{0, 1}, {1, 0}},
		},
	}

	fmt.Printf("lines : %#v \n ", lines)
	//for character in lines

	for i := 0; i < len(lines); i++ {

		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] == 'S' {
				direction := [2]int{0, 1}
				actual := rune(' ')
				steps := 0
				ipipe, jpipe := i, j
				for actual != 'S' {
					fmt.Printf("actual idx : %d %d \n", ipipe, jpipe)

					ipipe, jpipe = ipipe+direction[0], jpipe+direction[1]
					steps++
					actual = rune(lines[ipipe][jpipe])
					fmt.Printf("actual : %s \n", string(actual))
					if actual == '|' || actual == '-' || actual == 'S' {
						continue
					}

					if reflect.DeepEqual(direction, bends[actual].EntryDirection[0]) {
						fmt.Printf("dir0 : %#v \n ", direction)
						direction = bends[actual].OutDirection[0]
						fmt.Printf("dir : %#v \n ", direction)

					}
					if reflect.DeepEqual(direction, bends[actual].EntryDirection[1]) {
						fmt.Printf("dir0 : %#v \n ", direction)
						direction = bends[actual].OutDirection[1]
						fmt.Printf("dir : %#v \n ", direction)
					}
				}
				fmt.Printf("\n steps : %d ", steps)

				fmt.Printf("\n steps : %d ", steps/2)
			}

		}
	}

}
