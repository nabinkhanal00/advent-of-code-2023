package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type aocmap struct {
	dest  int
	start int
	rang  int
}
type aocseed struct {
	start int
	rang  int
}

func convert(input []string) []int {
	output := []int{}
	for _, inp := range input {
		res, _ := strconv.ParseInt(inp, 10, 64)
		output = append(output, int(res))
	}
	return output
}
func challenge9() {

	f, err := os.Open("input_day5.txt")
	if err != nil {
		fmt.Println("error occurred: ", err.Error())
		os.Exit(1)
	}
	seeds := []int{}
	seedToSoil := []aocmap{}
	soilToFertilizer := []aocmap{}
	fertilizerToWater := []aocmap{}
	waterToLight := []aocmap{}
	lightToTemperature := []aocmap{}
	temperatureToHumidity := []aocmap{}
	humidityToLocation := []aocmap{}

	scanner := bufio.NewScanner(f)
	var currentMap *[]aocmap
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else if strings.Contains(text, "seeds") {
			seeds = convert(strings.Fields(text[6:]))
		} else if strings.Contains(text, "seed-to-soil") {
			currentMap = &seedToSoil
		} else if strings.Contains(text, "soil-to-fertilizer") {
			currentMap = &soilToFertilizer
		} else if strings.Contains(text, "fertilizer-to-water") {
			currentMap = &fertilizerToWater
		} else if strings.Contains(text, "water-to-light") {
			currentMap = &waterToLight
		} else if strings.Contains(text, "light-to-temperature") {
			currentMap = &lightToTemperature
		} else if strings.Contains(text, "temperature-to-humidity") {
			currentMap = &temperatureToHumidity
		} else if strings.Contains(text, "humidity-to-location") {
			currentMap = &humidityToLocation
		} else {
			res := convert(strings.Fields(text))
			*currentMap = append(*currentMap, aocmap{dest: res[0], start: res[1], rang: res[2]})

		}
	}
	result := math.MaxInt
	// fmt.Println(seeds)
	for _, val := range seeds {
		res := getValue(seedToSoil, val)
		// fmt.Print(res, " ")
		res = getValue(soilToFertilizer, res)
		// fmt.Print(res, " ")
		res = getValue(fertilizerToWater, res)
		// fmt.Print(res, " ")
		res = getValue(waterToLight, res)
		// fmt.Print(res, " ")
		res = getValue(lightToTemperature, res)
		// fmt.Print(res, " ")
		res = getValue(temperatureToHumidity, res)
		// fmt.Print(res, " ")
		res = getValue(humidityToLocation, res)
		// fmt.Println(res)
		result = min(result, res)
	}
	fmt.Println(result)
}
func getValue(m []aocmap, key int) int {
	for _, s := range m {
		val := key - s.start
		if val >= 0 && val < s.rang {
			return s.dest + val
		}
	}
	return key
}

func challenge10() {

	f, err := os.Open("input_day5.txt")
	if err != nil {
		fmt.Println("error occurred: ", err.Error())
		os.Exit(1)
	}
	seeds := []aocseed{}
	seedToSoil := []aocmap{}
	soilToFertilizer := []aocmap{}
	fertilizerToWater := []aocmap{}
	waterToLight := []aocmap{}
	lightToTemperature := []aocmap{}
	temperatureToHumidity := []aocmap{}
	humidityToLocation := []aocmap{}

	scanner := bufio.NewScanner(f)
	var currentMap *[]aocmap
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else if strings.Contains(text, "seeds") {

			seedsRange := convert(strings.Fields(text[6:]))
			length := len(seedsRange)
			i := 0
			for i < length {
				seeds = append(seeds, aocseed{start: seedsRange[i], rang: seedsRange[i+1]})
				i += 2
			}
		} else if strings.Contains(text, "seed-to-soil") {
			currentMap = &seedToSoil
		} else if strings.Contains(text, "soil-to-fertilizer") {
			currentMap = &soilToFertilizer
		} else if strings.Contains(text, "fertilizer-to-water") {
			currentMap = &fertilizerToWater
		} else if strings.Contains(text, "water-to-light") {
			currentMap = &waterToLight
		} else if strings.Contains(text, "light-to-temperature") {
			currentMap = &lightToTemperature
		} else if strings.Contains(text, "temperature-to-humidity") {
			currentMap = &temperatureToHumidity
		} else if strings.Contains(text, "humidity-to-location") {
			currentMap = &humidityToLocation
		} else {
			res := convert(strings.Fields(text))
			*currentMap = append(*currentMap, aocmap{dest: res[0], start: res[1], rang: res[2]})

		}
	}
	// fmt.Println(seeds)
	c := make(chan int)
	length := len(seeds)
	for _, v := range seeds {
		go func(s int, r int, c chan int) {

			result := math.MaxInt
			for i := 0; i < r; i++ {
				val := s + i
				res := getValue(seedToSoil, val)
				// fmt.Print(res, " ")
				res = getValue(soilToFertilizer, res)
				// fmt.Print(res, " ")
				res = getValue(fertilizerToWater, res)
				// fmt.Print(res, " ")
				res = getValue(waterToLight, res)
				// fmt.Print(res, " ")
				res = getValue(lightToTemperature, res)
				// fmt.Print(res, " ")
				res = getValue(temperatureToHumidity, res)
				// fmt.Print(res, " ")
				res = getValue(humidityToLocation, res)
				// fmt.Println(res)
				result = min(result, res)
			}
			c <- result
		}(v.start, v.rang, c)
	}
	result := math.MaxInt
	for i := 0; i < length; i++ {
		result = min(result, <-c)
	}
	fmt.Println(result)
}
