package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
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

	timesLine := strings.TrimPrefix(lines[0], "Time:")
	times := strings.Fields(strings.TrimSpace(timesLine))
	distancesLine := strings.TrimPrefix(lines[1], "Distance:")
	distances := strings.Fields(strings.TrimSpace(distancesLine))
	ans := 1
	for i, time := range times {
		distance := distances[i]
		intTime, _ := strconv.Atoi(time)
		intDistance, _ := strconv.Atoi(distance)
		fmt.Printf("pair %d %d\n", intTime, intDistance)

		countWins := countWaysToWin(intTime, intDistance)
		fmt.Printf("count %d \n", countWins)
		if countWins > 0 {
			ans *= countWins
		}
	}
	fmt.Printf("%d \n", ans)

}

func countWaysToWin(time int, distance int) int {
	distance++
	discriminant := time*time - 4*distance
	fmt.Printf("discriminant %d \n", discriminant)

	if discriminant < 0 {
		return 0
	}
	if discriminant == 0 {
		s := quadraticFormula(-1, time, -distance, true)
		return countIntSolution(s)
	}
	if discriminant > 0 {
		s1 := quadraticFormula(-1, time, -distance, true)
		s2 := quadraticFormula(-1, time, -distance, false)
		return countIntSolutionTwo(s1, s2)
	}
	return 0
}
func quadraticFormula(ia, ib, ic int, positive bool) float64 {
	a := float64(ia)
	b := float64(ib)
	c := float64(ic)
	discriminant := b*b - 4*a*c
	if positive {
		return (-b + math.Sqrt(discriminant)) / (2 * a)
	}
	return (-b - math.Sqrt(discriminant)) / (2 * a)
}
func countIntSolution(sol float64) int {
	fmt.Printf("sol %d \n", sol)
	if math.Abs(sol-math.Round(sol)) < 0.000000001 {
		return 1
	}
	return 0
}

func countIntSolutionTwo(sol1, sol2 float64) int {
	var hisol, losol float64
	if sol1 < sol2 {
		hisol = sol2
		losol = sol1
	} else {
		hisol = sol1
		losol = sol2
	}
	fmt.Printf("sols %f %f\n", sol1, sol2)
	loSolInt := int(math.Ceil(losol))
	hiSolInt := int(math.Floor(hisol))
	fmt.Printf("sols %d %d\n", hiSolInt, loSolInt)
	return hiSolInt - loSolInt + 1

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
