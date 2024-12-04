package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	ans := 0
	for _, line := range lines {
		fmt.Printf("line  : %s", line)

		var diferences [][]int
		values := strings.Split(line, " ")
		hist := make([]int, len(values))
		for i, value := range values {
			hist[i], _ = strconv.Atoi(value)
		}
		fmt.Printf("value  : %#v\n", values)
		fmt.Printf("hist  : %#v\n", hist)
		current := hist
		for !allZero(current) {
			tmp := make([]int, len(current))
			copy(tmp, current)
			diferences = append(diferences, tmp)
			current = getDifferences(current)
			//	diferences = append(diferences, current)
			//	current = getDiferences(current)
		}
		fmt.Printf("diferences : %#v\n", diferences)
		for i := len(diferences) - 2; i >= 0; i-- {

			diference := diferences[i][len(diferences[i])-1]
			nextDiference := diferences[i+1][len(diferences[i+1])-1]
			diferences[i] = append(diferences[i], diference+nextDiference)
		}
		fmt.Printf("diferences : %#v\n", diferences)
		prediction := diferences[0][len(diferences[0])-1]
		fmt.Printf("\n pred : %d \n", prediction)
		ans += prediction
	}
	fmt.Printf("\nans : %d", ans)
}

func getDifferences(nums []int) []int {
	var differences []int
	for i, _ := range nums {
		if i == len(nums)-1 {
			continue
		}
		differences = append(differences, nums[i+1]-nums[i])
	}
	return differences
}
func allZero(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}
