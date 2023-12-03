package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func twoDigitNumber(str string) int {
	var start int
	var end int

	for start_pos := 0; start_pos < len(str); start_pos++ {
		if '0' <= str[start_pos] && '9' >= str[start_pos] {
			start = int(str[start_pos] - '0')
			break
		}
	}
	for end_pos := len(str) - 1; end_pos >= 0; end_pos-- {
		if '0' <= str[end_pos] && '9' >= str[end_pos] {
			end = int(str[end_pos] - '0')
			break
		}
	}
	return start*10 + end
}

func challenge1() {
	file, _ := os.Open("input_day1.txt")
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		num := twoDigitNumber(line)
		sum += num
	}
	fmt.Println(sum)
}

var m map[string]int = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func challenge2() {
	sum := 0

	front := "one|two|three|four|five|six|seven|eight|nine"
	frontExp := regexp.MustCompile("\\d|" + front)
	backExp := regexp.MustCompile("\\d|" + reverse(front))
	file, _ := os.Open("input_day1.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		frontFound := frontExp.FindString(line)
		sum += m[frontFound] * 10
		backFound := backExp.FindString(reverse(line))
		sum += m[reverse(backFound)]
	}
	fmt.Println(sum)
}

func reverse(str string) string {
	res := ""
	for _, char := range str {
		res = string(char) + res
	}
	return res
}
