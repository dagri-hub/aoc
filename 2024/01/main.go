package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

	separator := "   "
	listA := []int{}
	listB := []int{}

	for inputLineScanner.Scan() {
		lists := strings.SplitN(inputLineScanner.Text(), separator, 2)
		astr := lists[0]
		bstr := lists[1]

		if a, err := strconv.Atoi(astr); err == nil {
			listA = append(listA, a)
		}
		if b, err := strconv.Atoi(bstr); err == nil {
			listB = append(listB, b)
		}
	}

	part1(listA, listB)
	fmt.Println()
	part2(listA, listB)
}

func part1(listA, listB []int) {
	sort.Ints(listA)
	sort.Ints(listB)

	length := len(listA)
	diff := 0
	sum := 0
	for element := range length {
		diff = listA[element] - listB[element]
		if diff < 0 {
			diff = diff * -1
		}
		//fmt.Println(diff)
		sum = sum + diff
	}
	fmt.Println("Antwort für Teil 1 ist:", sum)
	// 3574690 ist richtig!
}

func part2(listA, listB []int) {
	fmt.Println(listA, listB)
	length := len(listA)
	sum := 0
	multi := 0
	for elementA := range length {
		sim := 0
		for elementB := range length {
			//fmt.Println(sim, listA[elementA], listB[elementB])
			if listA[elementA] == listB[elementB] {
				sim++
			}
		}
		multi = listA[elementA] * sim
		fmt.Printf("%v * %v = %v\n", sim, listA[elementA], multi)
		sum = sum + multi
	}
	fmt.Println("Antwort für Teil 2 ist:", sum)
	// ___ ist richtig!
}
