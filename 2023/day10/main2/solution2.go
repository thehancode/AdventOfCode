package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
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
	loop := getPipeLoop(lines, bends)
	fmt.Printf("loop :  %#v \n", loop)
	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], "S") {
			if len(lines[i]) > 30 {
				lines[i] = strings.ReplaceAll(lines[i], "S", "L")
			} else {
				lines[i] = strings.ReplaceAll(lines[i], "S", "F")
			}
		}
	}
	fmt.Printf("lines s : %#v \n ", lines)
	insiders := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if !inLoop(i, j, loop) {

				//fmt.Printf("check + :  %d %d \n", i, j)
				wallsE := countInDirection(i, j, [2]int{0, 1}, lines, loop)
				wallsW := countInDirection(i, j, [2]int{0, -1}, lines, loop)
				wallsN := countInDirection(i, j, [2]int{-1, 0}, lines, loop)
				wallsS := countInDirection(i, j, [2]int{1, 0}, lines, loop)

				if wallsE%2 != 0 || wallsW%2 != 0 || wallsN%2 != 0 || wallsS%2 != 0 {
					insiders++

					fmt.Printf("+inside +1 :  %d %d \n", i+1, j+1)
					fmt.Printf("+inside pair :  %d %d \n", i, j)

					fmt.Printf("counts EWNS :  %d %d %d %d \n", wallsE, wallsW, wallsN, wallsS)

				}
			}
		}
	}
	fmt.Printf("insiders :  %d \n", insiders)
}

func getPipeLoop(lines []string, bends map[rune]Direction) [][2]int {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] == 'S' {
				direction := [2]int{0, 1}
				actual := rune(' ')
				steps := 0
				ipipe, jpipe := i, j
				loop := append([][2]int{}, [2]int{ipipe, jpipe})
				for actual != 'S' {

					ipipe, jpipe = ipipe+direction[0], jpipe+direction[1]
					steps++
					loop = append(loop, [2]int{ipipe, jpipe})
					actual = rune(lines[ipipe][jpipe])
					if actual == '|' || actual == '-' || actual == 'S' {
						continue
					}

					if reflect.DeepEqual(direction, bends[actual].EntryDirection[0]) {
						direction = bends[actual].OutDirection[0]

					}
					if reflect.DeepEqual(direction, bends[actual].EntryDirection[1]) {
						direction = bends[actual].OutDirection[1]
					}
				}
				return loop
			}
		}
	}
	return make([][2]int, 0)
}

func countInDirection(i int, j int, direction [2]int, lines []string, loop [][2]int) int {
	actuali, actualj := i, j
	count := 0
	var actual rune
	onPipe := false
	lastCorner := ' '
	for checkIfInBounds(actuali+direction[0], actualj+direction[1], lines) {

		actuali, actualj = actuali+direction[0], actualj+direction[1]
		actual = rune(lines[actuali][actualj])
		if i == 4 && j == 4 {

			fmt.Printf("+actual :  %c \n", actual)
			fmt.Printf("+pair :  %d %d \n", actuali, actualj)
			fmt.Printf("+count :  %d \n", count)
			fmt.Printf("+onpipe :  %t \n", onPipe)
		}
		if !inLoop(actuali, actualj, loop) {
			continue
		}
		if !onPipe && (actual == '|' || actual == '-') {
			count++
			continue
		}
		incrementConditions := map[rune]rune{
			'L': '7',
			'J': 'F',
			'7': 'L',
			'F': 'J',
		}

		switch actual {
		case 'L', 'J', '7', 'F':
			if !onPipe {
				lastCorner = actual
			} else if lastCorner == incrementConditions[actual] {
				count++
			}
			onPipe = !onPipe
		}
	}

	return count
}

func inLoop(i int, j int, loop [][2]int) bool {
	for _, pair := range loop {
		if pair[0] == i && pair[1] == j {
			return true
		}
	}
	return false
}

func checkIfInBounds(i int, j int, lines []string) bool {
	return i >= 0 && i < len(lines) && j >= 0 && j < len(lines[0])
}
