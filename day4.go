package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseCard(line string) ([]int, []int) {
	cards := strings.Split(line, ":")[1]
	wp := strings.Split(cards, "|")
	w := strings.Fields(wp[0])
	p := strings.Fields(wp[1])
	var winning []int
	var present []int
	for _, wstr := range w {
		val, _ := strconv.ParseInt(wstr, 10, 64)
		winning = append(winning, int(val))
	}
	for _, pstr := range p {
		val, _ := strconv.ParseInt(pstr, 10, 64)
		present = append(present, int(val))
	}
	return winning, present
}
func challenge7() {

	f, err := os.Open("input_day4.txt")
	if err != nil {
		fmt.Println("error occurred: ", err.Error())
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		winning, present := parseCard(line)
		wmap := make(map[int]bool)
		for _, w := range winning {
			wmap[w] = true
		}
		val := 0
		for _, v := range present {
			if _, ok := wmap[v]; ok {
				if val == 0 {
					val = 1
				} else {
					val *= 2
				}
			}
		}
		sum += val
	}
	fmt.Println(sum)
}
