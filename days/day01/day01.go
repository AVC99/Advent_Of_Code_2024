package day01

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Day01() {
	leftRow, rightRow := readFileAndGetRows()
	if rightRow == nil || leftRow == nil {
		return
	}

	totalDistance := calculateTotalDistance(leftRow, rightRow)
	fmt.Println("Total distance is: ", totalDistance)

}
func calculateTotalDistance(leftRow []int, rightRow []int) int {
	fmt.Println("Sorting rows...")
	sort.Ints(rightRow)
	sort.Ints(leftRow)
	totalDistance := 0
	for i := 0; i < len(leftRow); i++ {
		fmt.Println(leftRow[i], rightRow[i])
		rowDistance := rightRow[i] - leftRow[i]
		if rowDistance < 0 {
			rowDistance = -rowDistance
		}
		totalDistance += rowDistance
	}
	return totalDistance
}

func readFileAndGetRows() ([]int, []int) {
	file, err := os.Open("days/day01.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return nil, nil
	}
	defer file.Close()
	var leftRowRead, rightRowRead string
	var leftRow, rightRow []int
	for {
		_, err := fmt.Fscanf(file, "%s %s\n", &leftRowRead, &rightRowRead)
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				fmt.Println("Error reading file", err)
				return nil, nil
			}
		}
		fmt.Println(leftRowRead, rightRowRead)
		firstRowReadInt, err := strconv.Atoi(leftRowRead)
		if err != nil {
			fmt.Println("Error converting string to int", err)
			return nil, nil
		}
		secondRowReadInt, err := strconv.Atoi(rightRowRead)
		if err != nil {
			fmt.Println("Error converting string to int", err)
			return nil, nil
		}
		leftRow = append(leftRow, firstRowReadInt)
		rightRow = append(rightRow, secondRowReadInt)
	}

	fmt.Println("Finished reading file :)")
	return leftRow, rightRow

}