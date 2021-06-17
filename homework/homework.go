package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	resolved := 0
	problems := strings.Split(input[:len(input)-1], ";")

	for _, problem := range problems {
		divided := strings.Split(problem, "-")
		if len(divided) == 1 {
			resolved++
		} else if len(divided) > 1 {
			start, _ := strconv.Atoi(divided[0])
			end, _ := strconv.Atoi(divided[1])
			resolved += end - start + 1
		}
	}

	fmt.Println(resolved)
}
