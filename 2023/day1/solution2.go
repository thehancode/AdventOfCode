package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    
    scanner := bufio.NewScanner(os.Stdin)
    var lines []string
    
    for scanner.Scan() {
        line := scanner.Text() 
        lines = append(lines, line) 
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
    
    total := 0 
    fmt.Println("\nCollected Lines:")
    for _, line := range lines {
        fmt.Println(line)
        number, err := findCalibrationNumber(line)
        total += number
        if err != nil {
      fmt.Println("Error")
    }
        fmt.Println(number)
    }
  fmt.Println(total)
}

func getFirstDigit(input string) int {
    for _, char := range input {
        if char >= '0' && char <= '9' {
            return int(char - '0')
        }
    }
    return 0 
}

func getLastDigit(input string) int {
    for i := len(input) - 1; i >= 0; i-- {
        if input[i] >= '0' && input[i] <= '9' {
            return int(input[i] - '0')
        }
    }
    return 0 
}

func findCalibrationNumber (input string) (int, error) {
    processedString := literalsToDigit(input)
    fmt.Println(processedString)
    firstDigit := getFirstDigit(processedString)
    lastDigit := getLastDigit(processedString)
    return firstDigit*10 + lastDigit, nil
}

func getFirstLiteralIndex(input string, digits [9]string) (int, int) {
    for j := 0; j < len(input); j++ {
        for i := 0; i < len(digits); i++ {
            matchFound := true
            for k := 0; k < len(digits[i]); k++ {
                if j+k >= len(input) || input[j+k] != digits[i][k] {
                    matchFound = false
                    break
                }
            }
            if matchFound {
                return j, i
            }
        }
    }
    return -1, -1
}

func getLastLiteralIndex (input string, digits [9]string ) (int , int) {
    for j := len(input); j > 0 ; j-- {
        for i := 0; i < len(digits); i++ {
            matchFound := true
            for k := 0; k < len(digits[i]); k++ {
                if j+k >= len(input) || input[j+k] != digits[i][k] {
                    matchFound = false
                    break
                }
            }
            if matchFound {
                return j, i
            }
        }
    }
    return -1, -1
}
func literalsToDigit (input string) string {
        digits := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    indexFirstLiteral , firstValue := getFirstLiteralIndex(input, digits)
    indexLastLiteral , lastValue := getLastLiteralIndex(input, digits)

    runes := []rune(input)
    if indexFirstLiteral != -1 && firstValue != -1 {
      runes[indexFirstLiteral] = rune (firstValue + '1' )
  }
        
    if indexLastLiteral != -1 && lastValue != -1 {
  runes[indexLastLiteral] = rune (lastValue + '1')

  }

    return string(runes)
}
