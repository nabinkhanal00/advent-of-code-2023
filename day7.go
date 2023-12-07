package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type HandType int

var (
	HighCard     HandType = 1
	OnePair      HandType = 2
	TwoPair      HandType = 3
	ThreeOfAKind HandType = 4
	FullHouse    HandType = 5
	FourOfAKind  HandType = 6
	FiveOfAKind  HandType = 7
)

type Hand struct {
	cards    string
	handType HandType
}

type Group struct {
	hand Hand
	bid  int
}

func (hand *Hand) findType() {
	m := make(map[rune]int)
	for _, c := range hand.cards {
		m[c]++
	}
	var (
		twos   = 0
		threes = 0
		fours  = 0
		fives  = 0
	)
	for _, val := range m {
		switch val {
		case 2:
			twos++
		case 3:
			threes++
		case 4:
			fours++
		case 5:
			fives++
		}
	}
	if fives == 1 {
		hand.handType = FiveOfAKind
	} else if fours == 1 {
		hand.handType = FourOfAKind
	} else if threes == 1 {
		if twos == 1 {
			hand.handType = FullHouse
		} else {
			hand.handType = ThreeOfAKind
		}
	} else if twos == 2 {
		hand.handType = TwoPair
	} else if twos == 1 {
		hand.handType = OnePair
	} else {
		hand.handType = HighCard
	}
}

func challenge13() {
	f, _ := os.Open("input_day7.txt")
	scanner := bufio.NewScanner(f)

	var groups []Group
	for scanner.Scan() {
		text := scanner.Text()
		parsed := strings.Fields(text)
		bid, _ := strconv.ParseInt(parsed[1], 10, 64)
		hand := Hand{cards: parsed[0]}
		hand.findType()
		groups = append(groups, Group{hand: hand, bid: int(bid)})
	}
	slices.SortFunc[[]Group, Group](groups, comparator)
	sum := 0
	for i, g := range groups {
		sum += (i + 1) * g.bid
	}
	fmt.Println(sum)
}

var vals map[byte]int = map[byte]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

func comparator(a Group, b Group) int {
	if a.hand.handType == b.hand.handType {
		for i := 0; i < 5; i++ {
			if vals[a.hand.cards[i]] < vals[b.hand.cards[i]] {
				return -1
			} else if vals[a.hand.cards[i]] > vals[b.hand.cards[i]] {
				return 1
			}
		}
		return 0

	}
	return int(a.hand.handType) - int(b.hand.handType)
}
