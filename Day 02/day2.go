package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("directions.txt")
	var horizontal = 0
	var depth = 0
	var aim = 0

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")
		fmt.Printf("%s %s\n", data[0], data[1])
		switch data[0] {
		case "forward":
			result, err := strconv.Atoi(data[1])
			if err == nil {
				horizontal += result
				depth += aim * result
			}
			break
		case "up":
			result, err := strconv.Atoi(data[1])
			if err == nil {
				aim -= result
			}
			break
		case "down":
			result, err := strconv.Atoi(data[1])
			if err == nil {
				aim += result
			}
			break
		}
	}
	fmt.Printf("Final Horizontal: %d, Final Depth %d\nMultiple is %d", horizontal, depth, horizontal*depth)
}
