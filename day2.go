package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(str string) (int, int, int, int) {
	green := 0
	blue := 0
	red := 0
	a := strings.Split(str, ":")
	gameInfo := a[0]
	ballInfo := a[1]
	gameNumber, _ := strconv.ParseInt(strings.Split(gameInfo, " ")[1], 10, 32)
	for _, st := range strings.Split(ballInfo, ";") {
		for _, s := range strings.Split(strings.Trim(st, " "), ",") {
			bInfo := strings.Split(strings.Trim(s, " "), " ")
			switch bInfo[1] {
			case "red":
				val, _ := strconv.ParseInt(bInfo[0], 10, 32)
				red = max(red, int(val))
			case "green":
				val, _ := strconv.ParseInt(bInfo[0], 10, 32)
				green = max(green, int(val))
			case "blue":
				val, _ := strconv.ParseInt(bInfo[0], 10, 32)
				blue = max(blue, int(val))
			}
		}
	}
	return int(gameNumber), red, green, blue
}

func challenge3() {
	file, _ := os.Open("input_day2.txt")
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		gameNumber, red, green, blue := parse(line)
		if red <= 12 && green <= 13 && blue <= 14 {
			sum += gameNumber
		}
	}
	fmt.Println(sum)
}
func challenge4() {

	file, _ := os.Open("input_day2.txt")
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		_, red, green, blue := parse(line)
		power := red * green * blue
		sum += power
	}
	fmt.Println(sum)
}
