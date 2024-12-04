package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Galaxy struct {
	x int
	y int
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
	//empty row and column map []int : bool
	emptyRow := make(map[int]bool, 0)
	emptyColumn := make(map[int]bool, 0)

	galaxys := make([]Galaxy, 0)
	for i := 0; i < len(lines); i++ {
		isEmpty := true
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] != '.' {
				isEmpty = false
			}
		}
		if isEmpty == true {
			emptyRow[i] = true
		}
	}
	for j := 0; j < len(lines[0]); j++ {
		isEmpty := true
		for i := 0; i < len(lines); i++ {
			if lines[i][j] != '.' {
				isEmpty = false
			}
		}
		if isEmpty == true {
			emptyColumn[j] = true
		}
	}
	fmt.Printf("emptyRow: %v\n", emptyRow)
	fmt.Printf("emptyColumn: %v\n", emptyColumn)
	reali, realj := 0, 0
	for i := 0; i < len(lines); i++ {
		if emptyRow[i] == true {
			reali++
		}
		realj = 0
		for j := 0; j < len(lines[0]); j++ {
			if emptyColumn[j] == true {
				realj++
			}
			if lines[i][j] == '#' {
				galaxys = append(galaxys, Galaxy{x: reali, y: realj})
			}
			realj++
		}
		reali++
	}
	ans := 0.0
	for i, galaxy := range galaxys {
		for j := i + 1; j < len(galaxys); j++ {
			disx := math.Abs(float64(galaxys[j].x - galaxy.x))
			disy := math.Abs(float64(galaxys[j].y - galaxy.y))
			fmt.Printf("galaxy1: %v, galaxy2: %v\n", galaxy, galaxys[j])
			//i j`
			fmt.Printf("gal1 idx: %v, gal2 idx: %v\n", i, j)
			fmt.Printf("disx: %v, disy: %v\n", disx, disy)
			ans += disx + disy
		}
	}
	fmt.Printf("ans: %d\n", int(ans))

}
