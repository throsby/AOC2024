package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read each line of the file
	num := 0
	var leftSlice []int
	var rightSlice []int

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Print(num, "\t")
		var splitLine = strings.Split(line, " ")
		var left, err1 = strconv.Atoi(splitLine[0])
		if err1 != nil {
			fmt.Println("Error with left string to int conversion:", err1)
			return
		}
		var right, err2 = strconv.Atoi(splitLine[3])
		if err2 != nil {
			fmt.Println("Error with right string to int conversion:", err2)
			return
		}

		leftSlice = append(leftSlice, left)
		rightSlice = append(rightSlice, right)
		num++
	}

	slices.Sort(leftSlice)
	slices.Sort(rightSlice)

	// fmt.Print(leftSlice, "\n\n", rightSlice, "\n\n")

	var sum = 0
	for i := 0; i < len(leftSlice); i++ {
		diff := rightSlice[i] - leftSlice[i]
		diff64 := float64(diff)
		// Convert it back to int, don't sum floats
		sum += int(math.Abs(diff64))
		// fmt.Println(sum)
	}
	fmt.Print(sum)
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
