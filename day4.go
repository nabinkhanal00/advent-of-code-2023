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
func findMatching(a, b []int) int {

	wmap := make(map[int]bool)
	for _, w := range a {
		wmap[w] = true
	}
	val := 0
	for _, v := range b {
		if _, ok := wmap[v]; ok {
			val++
		}
	}
	return val
}
func challenge8() {

	f, err := os.Open("input_day4.txt")
	if err != nil {
		fmt.Println("error occurred: ", err.Error())
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	sum := 0
	res := make(map[int]int)
	cur := 1
	for scanner.Scan() {
		res[cur]++
		line := scanner.Text()
		winning, present := parseCard(line)
		count := findMatching(winning, present)
		for i := 0; i < count; i++ {
			res[cur+i+1] += res[cur]
		}
		sum += res[cur]
		cur++
	}
	fmt.Println(sum)
}
