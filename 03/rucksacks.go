package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func sharedRunes(s1, s2 string) []rune {
	var matches []rune
	for r := range s1 {
		if strings.ContainsRune(s2, rune(s1[r])) {
			alreadythere := false
			for m := range matches {
				if matches[m] == rune(s1[r]) {
					alreadythere = true
				}
			}
			if !alreadythere {
				matches = append(matches, rune(s1[r]))
			}
		}
	}
	return matches
}

const az = "abcdefghijklmnopqrstuvwxyz"
const AZ = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func gimmePrio(r rune) int {
	i := strings.Index(az, string(r))
	if i < 0 {
		return strings.Index(AZ, string(r)) + 27
	}
	return i + 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	prioSum := 0
	badgeSum := 0
	group := make([]string, 0, 3)
	line := 0
	for scanner.Scan() {
		err := scanner.Err()
		if err == io.EOF {
			return
		}
		value := scanner.Text()
		line += 1
		if len(value)%2 != 0 {
			panic(errors.New("Compartments not the same size"))
		}
		if value == "" {
			continue
		}
		group = append(group, value)
		if line == 3 {
			line = 0
			if len(group) != 3 {
				panic(errors.New(fmt.Sprintf("You did something wrong: %d", len(group))))
			}
			sh2 := sharedRunes(string(sharedRunes(group[0], group[1])), group[2])
			if len(sh2) != 1 {
				panic(errors.New(fmt.Sprintf("Group is invalid, shared badges: %d", len(sh2))))
			}
			badgeSum += gimmePrio(sh2[0])
			group = make([]string, 0, 3)
		}
		firstHalf := value[0 : len(value)/2]
		secondHalf := value[len(value)/2 : len(value)]
		shared := sharedRunes(firstHalf, secondHalf)
		fmt.Printf("%s -- %s\n", firstHalf, secondHalf)
		for i := range shared {
			prioSum += gimmePrio(shared[i])
		}
	}
	fmt.Printf("%d\n", prioSum)
	fmt.Printf("%d\n", badgeSum)
}
