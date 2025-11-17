package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const minDiff = 1
const maxDiff = 3
const separator = " "

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	inputLineScanner := bufio.NewScanner(f)
	if err := inputLineScanner.Err(); err != nil {
		log.Fatal(err)
	}

	report := []int{}
	safeReports := 0

	for inputLineScanner.Scan() {
		reportStrSlice := strings.Split(inputLineScanner.Text(), separator)
		length := len(reportStrSlice)
		for elem := range length {
			if level, err := strconv.Atoi(reportStrSlice[elem]); err == nil {
				report = append(report, level)
			}
		}
		result := testReport(report)
		if result == true {
			safeReports++
		}
		report = []int{}
	}
	fmt.Println("safeReports:", safeReports)
	// 686
}

func testReport(report []int) bool {
	up := 0
	down := 0
	reportLength := len(report) - 1
	for l := range reportLength {
		diff := report[l+1] - report[l]
		if diff > 0 && diff >= minDiff && diff <= maxDiff {
			up++
		} else if diff < 0 && diff <= (-1*minDiff) && diff >= (-1*maxDiff) {
			down++
		}
	}
	if up == reportLength || down == reportLength {
		return true
	} else {
		return false
	}
}
