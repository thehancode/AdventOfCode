package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x         int
	y         int
	color     string
	direction rune
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
	directions := map[string][2]int{
		"U": {-1, 0},
		"D": {1, 0},
		"L": {0, -1},
		"R": {0, 1},
	}

	nextClockwise := map[string][2]int{
		"U": {0, 1},
		"D": {0, -1},
		"L": {-1, 0},
		"R": {1, 0},
	}

	minx, miny, maxx, maxy := math.MaxInt32, math.MaxInt32, math.MinInt32, math.MinInt32
	border := make([]Point, 0)
	borderMap := make(map[int][]Point)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		steps, _ := strconv.Atoi(parts[1])
		color := parts[2]
		x, y := 0, 0

		for i := 0; i < steps; i++ {
			x += directions[direction][0]
			y += directions[direction][1]
			currentPoint := Point{x, y, color, rune(direction)}
			border = append(border, currentPoint)
			borderMap[x] = append(borderMap[x], currentPoint)
		}
		if x < minx {
			minx = x
		}
		if x > maxx {
			maxx = x
		}
		if y < miny {
			miny = y
		}
		if y > maxy {
			maxy = y
		}
	}

	countInside := 0
	for i := minx; i <= maxx; i++ {
		points := borderMap[i]
		isInside := false
		currentPointIdx := 0
		for j := miny; j <= maxy; j++ {

			if points[currentPointIdx].y == j {
				if points[currentPointIdx].direction == 'W' || points[currentPointIdx].direction == 'E'{

				}
				isInside = !isInside
				currentPointIdx++
			} else if isInside {
				countInside++
			}
		}
		fmt.Println()
	}
}

func inv(dir [2]int) [2]int {
	return [2]int{-dir[0], -dir[1]}
}
