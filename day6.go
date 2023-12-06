package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func challenge11() {
	f, _ := os.Open("input_day6.txt")
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	t := scanner.Text()
	scanner.Scan()
	d := scanner.Text()
	times := convert(strings.Fields(strings.Split(t, ":")[1]))
	distances := convert(strings.Fields(strings.Split(d, ":")[1]))
	fmt.Println(times, distances)
	result := 1
	for i := 0; i < len(times); i++ {
		time := times[i]
		distance := distances[i]
		s := findStartPossible(time, distance)
		e := findEndPossible(time, distance)
		result *= (e - s + 1)
	}
	fmt.Println(result)
}
func challenge12() {

	f, _ := os.Open("input_day6.txt")
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	t := scanner.Text()
	scanner.Scan()
	d := scanner.Text()
	time, _ := strconv.ParseInt(strings.Join(strings.Fields(strings.Split(t, ":")[1]), ""), 10, 64)
	distance, _ := strconv.ParseInt(strings.Join(strings.Fields(strings.Split(d, ":")[1]), ""), 10, 64)
	fmt.Println(time, distance)
	s := findStartPossible(int(time), int(distance))
	e := findEndPossible(int(time), int(distance))
	result := (e - s + 1)
	fmt.Println(result)
}
func findStartPossible(time, distance int) int {
	for i := 1; i <= time; i++ {
		if (i)*(time-i) <= distance {
			continue
		} else {
			return i
		}
	}
	return -1
}
func findEndPossible(time, distance int) int {
	for i := time; i >= 1; i-- {
		if i*(time-i) <= distance {
			continue
		} else {
			return i
		}
	}
	return -1
}
