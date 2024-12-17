package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("dottxt.txt")
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

	// Create a map with the KEY as the value in the left or right and the count as the VALUE
	// fmt.Print(leftSlice)
	// for i, _ := range leftSlice {
	// 	 fmt.Print(leftSlice[i], " ")
	// 	 fmt.Print(rightSlice[i], "\t\t")
	// }
	// var leftFrequency map[int]int
	var sum int
	for i := range leftSlice {
		fmt.Print(i)
		for j := range rightSlice {
			if j < len(rightSlice)-1 && rightSlice[j] == rightSlice[j+1] {
				fmt.Print(rightSlice[j] == rightSlice[j+1])
				sum++
			}
		}
	}
	fmt.Println(sum, " ")

	// fmt.Print(leftSlice)
	var sum1 = 0
	for i := range rightSlice {

		if i < len(leftSlice)-1 && leftSlice[i] == leftSlice[i+1] {
			// fmt.Print(leftSlice[i] == leftSlice[i+1])
			sum1++
		}
	}
	// fmt.Print(" ", sum1)

}
