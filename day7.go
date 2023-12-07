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

func (hand *Hand) findType2() {
	m := make(map[rune]int)
	jokers := 0
	for _, c := range hand.cards {
		if c == 'J' {
			jokers++
			continue
		}
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
	hand.handType = correctHandtype(hand.handType, jokers)
}
func correctHandtype(handType HandType, jokers int) HandType {
	if handType == FourOfAKind {
		if jokers == 1 {
			return FiveOfAKind
		}
	} else if handType == ThreeOfAKind {
		if jokers == 2 {
			return FiveOfAKind
		} else if jokers == 1 {
			return FourOfAKind
		}
	} else if handType == TwoPair {
		if jokers == 1 {
			return FullHouse
		}
	} else if handType == OnePair {
		if jokers == 3 {
			return FiveOfAKind
		} else if jokers == 2 {
			return FourOfAKind
		} else if jokers == 1 {
			return ThreeOfAKind
		}
	} else if handType == HighCard {
		if jokers == 5 || jokers == 4 {
			return FiveOfAKind
		} else if jokers == 3 {
			return FourOfAKind
		} else if jokers == 2 {
			return ThreeOfAKind
		} else if jokers == 1 {
			return OnePair
		}
	}
	return handType
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

func challenge14() {

	f, _ := os.Open("input_day7.txt")
	scanner := bufio.NewScanner(f)

	var groups []Group
	for scanner.Scan() {
		text := scanner.Text()
		parsed := strings.Fields(text)
		bid, _ := strconv.ParseInt(parsed[1], 10, 64)
		hand := Hand{cards: parsed[0]}
		hand.findType2()
		groups = append(groups, Group{hand: hand, bid: int(bid)})
	}
	slices.SortFunc[[]Group, Group](groups, comparator2)
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

var vals2 map[byte]int = map[byte]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 0,
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

func comparator2(a Group, b Group) int {
	if a.hand.handType == b.hand.handType {
		for i := 0; i < 5; i++ {
			if vals2[a.hand.cards[i]] < vals2[b.hand.cards[i]] {
				return -1
			} else if vals2[a.hand.cards[i]] > vals2[b.hand.cards[i]] {
				return 1
			}
		}
		return 0

	}
	return int(a.hand.handType) - int(b.hand.handType)
}
