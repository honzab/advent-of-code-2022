package main

import "bufio"
import "fmt"
import "os"
import "io"
import "strconv"

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
		fmt.Printf("%r\n", value)
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

	max_value := uint64(0)
	max_index := 0
	for i := range elfs {
		if elfs[i] >= max_value {
			max_index = i
			max_value = elfs[i]
		}
	}
	fmt.Printf("%d elf carrying %d\n", max_index+1, max_value)
}
