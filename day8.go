package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func challenge15() {
	f, _ := os.Open("input_day8.txt")
	scanner := bufio.NewScanner(f)

	type lr struct {
		left  string
		right string
	}

	m := make(map[string]lr)

	scanner.Scan()
	directions := scanner.Text()
	scanner.Scan()

	for scanner.Scan() {
		text := scanner.Text()
		splits := strings.Split(text, "=")
		d := strings.Split(splits[1][2:len(splits[1])-1], ", ")
		m[strings.Trim(splits[0], " ")] = lr{left: d[0], right: d[1]}
	}
	step := 0
	count := 0
	pos := "AAA"
	for {
		dir := directions[count]
		if dir == 'L' {
			pos = m[pos].left
		} else {
			pos = m[pos].right
		}
		step++
		count++
		if pos == "ZZZ" {
			fmt.Println(step)
			return
		}
		if count == len(directions) {
			count = 0
		}
	}

}
