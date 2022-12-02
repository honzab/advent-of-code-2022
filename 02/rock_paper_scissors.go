package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// A X rock
// B Y paper
// C Z scissors

func winPoints(a, b string) int {
	if (a == "A" && b == "X") || (a == "B" && b == "Y") || (a == "C" && b == "Z") {
		return 3
	}
	if (a == "C" && b == "X") || (a == "A" && b == "Y") || (a == "B" && b == "Z") {
		return 6
	}
	return 0
}

func usagePoints(a string) (int, error) {
	if a == "X" {
		return 1, nil
	} else if a == "Y" {
		return 2, nil
	} else if a == "Z" {
		return 3, nil
	} else {
		return 0, errors.New(fmt.Sprintf("Invalid character %s", a))
	}
}

func whatToDo(a, b string) (string, error) {
	if b == "X" {
		// Need to lose
		if a == "A" {
			return "Z", nil
		} else if a == "B" {
			return "X", nil
		} else if a == "C" {
			return "Y", nil
		}
	} else if b == "Y" {
		// Need to draw
		if a == "A" {
			return "X", nil
		} else if a == "B" {
			return "Y", nil
		} else if a == "C" {
			return "Z", nil
		}
	} else if b == "Z" {
		// Need to win
		if a == "A" {
			return "Y", nil
		} else if a == "B" {
			return "Z", nil
		} else if a == "C" {
			return "X", nil
		}
	}
	return "", errors.New(fmt.Sprintf("Unkown chars: %s, %s", a, b))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	totalScore := 0
	totalScore2 := 0
	for scanner.Scan() {
		err := scanner.Err()
		if err == io.EOF {
			return
		}
		value := scanner.Text()
		values := strings.Split(value, " ")
		if len(values) != 2 {
			panic("Invalid line!")
		}
		a, b := values[0], values[1]

		up, err := usagePoints(b)
		if err != nil {
			panic(err)
		}
		totalScore += winPoints(a, b) + up
		fmt.Printf("%s: %d + %d\n", value, winPoints(a, b), up)

		b, err = whatToDo(a, b)
		if err != nil {
			panic(err)
		}
		up, err = usagePoints(b)
		if err != nil {
			panic(err)
		}
		totalScore2 += winPoints(a, b) + up
	}
	fmt.Printf("%d\n", totalScore)
	fmt.Printf("%d\n", totalScore2)
}
