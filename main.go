package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/AVC99/Advent_Of_Code_2024/days"
)

func main() {
	fmt.Println("🎄 Hello welcome to my AVC99 of Code for 2024! 🎅")
	fmt.Println("🦌 I will be solving the puzzles in Go ✨ ")
	fmt.Println("☃️ Heres the days I have completed so far: ")
	fmt.Println("🎁 Day 1: Historian Hysteria")

	day := askForDay()
	displayDay(day)
}
func displayDay(day int) {
	switch day {
	case 1:
		fmt.Println("Running -> Day 1: Historian Hysteria")
		days.Day01()
	}
}
func askForDay() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter the day number (1-25): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.", err)
			continue
		}
		input = strings.TrimSpace(input)
		day, err := strconv.Atoi(input)
		if err != nil || day < 1 || day > 25 {
			fmt.Println("Invalid input. Please enter a number between 1 and 25.")
			continue
		}

		return day

	}
}
