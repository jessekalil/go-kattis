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

	values := strings.Split(input[:len(input)-1], " ")

	first, _ := strconv.Atoi(values[1])
	second, _ := strconv.Atoi(values[0])

	fmt.Println(first + second)
}
