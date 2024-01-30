package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Create a bufio.Reader that reads from standard input
    scanner := bufio.NewScanner(os.Stdin)

    // Create a slice to hold the lines
    var lines []string

    // Use scanner.Scan() in a loop to read lines until an error or EOF
    for scanner.Scan() {
        line := scanner.Text() // Get the text of the scanned line
        lines = append(lines, line) // Append the line to the slice
    }

    // Check for errors during scanning (excluding io.EOF)
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }

    // Print the colleocted lines
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

    firstDigit := getFirstDigit(input)
    lastDigit := getLastDigit(input)
    return firstDigit*10 + lastDigit, nil
}


    
