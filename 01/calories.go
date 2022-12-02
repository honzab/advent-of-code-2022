package main

import "bufio"
import "fmt"
import "os"
import "io"
import "strconv"
import "sort"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	elfs := make([]uint64, 1000)
	index := 0

	for scanner.Scan() {
		err := scanner.Err()
		if err == io.EOF {
			return
		}
		value := scanner.Text()
		if value == "" {
			index = index + 1
			continue
		}
		ii, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		elfs[index] += ii
	}

	sort.Slice(elfs, func(i, j int) bool {
		return elfs[i] > elfs[j]
	})
	fmt.Printf("%d\n", elfs[0])
	fmt.Printf("%d\n", elfs[0]+elfs[1]+elfs[2])
}
