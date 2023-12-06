package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Number struct {
	value   int // value
	row     int // row number
	columns int // column start
	columne int // column end
	visited bool
}
type Symbol struct {
	s rune
	r int
	c int
}

func parseLine(line string, lineNumber int) ([]Number, []Symbol) {
	start := false
	current := ""
	s := -1
	var result []Number
	var symbols []Symbol
	for i, char := range line {
		if unicode.IsDigit(char) {
			if !start {
				start = true
				s = i
			}
			current = current + string(char)
		} else {
			if char != '.' {
				symbols = append(symbols, Symbol{s: char, r: lineNumber, c: int(i)})
			}
			start = false
			if current != "" {
				num, _ := strconv.ParseInt(current, 10, 64)

				result = append(result, Number{value: int(num), row: lineNumber, columns: int(s), columne: int(i - 1), visited: false})
				current = ""
			}

		}
	}
	if current != "" {
		num, _ := strconv.ParseInt(current, 10, 64)

		result = append(result, Number{value: int(num), row: lineNumber, columns: int(s), columne: int(s) + len(current), visited: false})
		current = ""
	}
	return result, symbols
}

func challenge5() {
	f, err := os.Open("input_day3.txt")
	if err != nil {
		fmt.Println("error occurred: ", err.Error())
		os.Exit(1)
	}
	content, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("error occurred: ", err.Error())
		os.Exit(1)
	}
	lines := strings.Split(string(content), "\n")
	var numbers []Number
	var symbols []Symbol
	for index, line := range lines {
		res, syms := parseLine(line, int(index))
		numbers = append(numbers, res...)
		symbols = append(symbols, syms...)
	}
	sum := 0
	for _, s := range symbols {
		for _, n := range numbers {
			if !n.visited {
				if math.Abs(float64(n.row-s.r)) <= 1 && n.columns-1 <= s.c && n.columne+1 >= s.c {
					sum += n.value
				}
			}
		}
	}
	fmt.Println(sum)
}
func challenge6() {

	f, err := os.Open("input_day3.txt")
	if err != nil {
		fmt.Println("error occurred: ", err.Error())
		os.Exit(1)
	}
	content, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("error occurred: ", err.Error())
		os.Exit(1)
	}
	lines := strings.Split(string(content), "\n")
	var numbers []Number
	var symbols []Symbol
	for index, line := range lines {
		res, syms := parseLine(line, int(index))
		numbers = append(numbers, res...)
		symbols = append(symbols, syms...)
	}
	sum := 0
	for _, s := range symbols {
		if s.s == '*' {
			count := 0
			var num1, num2 int
			for _, n := range numbers {
				if math.Abs(float64(n.row-s.r)) <= 1 && n.columns-1 <= s.c && n.columne+1 >= s.c {
					count++
					if count == 1 {
						num1 = n.value
					} else if count == 2 {
						num2 = n.value
					} else if count > 2 {
						break
					}
				}
			}
			if count == 2 {
				sum += num1 * num2
			}
		}
	}
	fmt.Println(sum)
}
