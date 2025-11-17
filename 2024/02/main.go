package main

// This f*** does not work :-(

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
	safeReports1 := 0
	safeReports2 := 0

	for inputLineScanner.Scan() {
		reportStrSlice := strings.Split(inputLineScanner.Text(), separator)
		length := len(reportStrSlice)
		for elem := range length {
			if level, err := strconv.Atoi(reportStrSlice[elem]); err == nil {
				report = append(report, level)
			}
		}
		result1 := part1(report)
		if result1 == true {
			safeReports1++
		}
		result2 := part2(report)
		if result2 == true {
			safeReports2++
			fmt.Println(report, "is safe")
		}
		report = []int{}
	}
	fmt.Println("safeReports:", safeReports1, "/", safeReports2)
	// 686 /
}

func part1(report []int) bool {
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

func part2(report []int) bool {
	up := 0
	down := 0
	sumup := 0
	sumdown := 0
	//goingup := report[1] > report[0]
	//fmt.Println(goingup)
	reportLength := len(report)
	newreport := []int{report[0]}
	fmt.Println(report)
	fmt.Println(newreport)
	if report[0] == report[1] {
		newreport = []int{}
	}
	for l := 1; l < reportLength; l++ {
		//fmt.Println("l:", l)
		diff := report[l] - report[l-1]
		//fmt.Println("diff:", diff)
		if diff > 0 && diff >= minDiff && diff <= maxDiff {
			up++
			sumup++
		} else if diff < 0 && diff <= (-1*minDiff) && diff >= (-1*maxDiff) {
			down++
			sumdown++
		}
		//fmt.Println("up:", up)
		//fmt.Println("down:", down)
		//fmt.Println("sumup:", sumup)
		//fmt.Println("sumdown:", sumdown)
		if up > 0 && down > 0 {
			//fmt.Println("not adding the level")
			up = 0
			down = 0
		} else {
			//fmt.Println("added the level")
			newreport = append(newreport, report[l])
		}
		//fmt.Println(newreport)
		//newreportLength := len(newreport)
		//fmt.Println(newreportLength)
		//fmt.Println()
	}
	newreportLength := len(newreport)
	//fmt.Println("length:", reportLength, newreportLength)
	if up == reportLength-1 || down == reportLength-1 || sumup+sumdown == reportLength-1 || up == newreportLength-1 || down == newreportLength-1 {
		return true
	} else {
		return false
	}
}
