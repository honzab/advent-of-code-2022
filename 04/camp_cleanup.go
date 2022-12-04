package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func isInRangeBothWays(r1, r2 []string) bool {
	b11, _ := strconv.ParseInt(r1[0], 10, 64)
	b12, _ := strconv.ParseInt(r1[1], 10, 64)
	b21, _ := strconv.ParseInt(r2[0], 10, 64)
	b22, _ := strconv.ParseInt(r2[1], 10, 64)
	if (b11 <= b21) && (b12 >= b22) {
		return true
	}
	if (b21 <= b11) && (b22 >= b12) {
		return true
	}
	return false
}

func overlapAtAll(r1, r2 []string) bool {
	b11, _ := strconv.ParseInt(r1[0], 10, 64)
	b12, _ := strconv.ParseInt(r1[1], 10, 64)
	b21, _ := strconv.ParseInt(r2[0], 10, 64)
	b22, _ := strconv.ParseInt(r2[1], 10, 64)

	if (b12 < b21) || (b22 < b11) {
		return false
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	counter := 0
	overlapCounter := 0
	for scanner.Scan() {
		err := scanner.Err()
		if err == io.EOF {
			return
		}
		value := scanner.Text()
		if value == "" {
			continue
		}
		sections := strings.Split(value, ",")
		if len(sections) != 2 {
			panic(errors.New("Wrong amount of sections"))
		}
		range1 := strings.Split(sections[0], "-")
		if len(range1) != 2 {
			panic(errors.New("Wrong section definition"))
		}
		range2 := strings.Split(sections[1], "-")
		if len(range2) != 2 {
			panic(errors.New("Wrong section definition"))
		}
		if isInRangeBothWays(range1, range2) {
			counter += 1
			fmt.Printf("Match: %s\n", value)
		}
		if overlapAtAll(range1, range2) {
			overlapCounter += 1
			fmt.Printf("Overlap: %s\n", value)
		}
	}
	fmt.Printf("%d\n", counter)
	fmt.Printf("%d\n", overlapCounter)
}
