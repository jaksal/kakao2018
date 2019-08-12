package kakao2018

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

type dart struct {
	point  int
	bonus  int
	option byte
	result int
}

func k2nd(in string) int {
	var result int

	count := -1
	var sp string
	var try [3]dart
	for i := 0; i < len(in); i++ {
		t := in[i]

		switch {
		case unicode.IsDigit(rune(t)):
			if sp == "" {
				count++
			}
			sp += string(t)
		case t == 'D':
			p, _ := strconv.Atoi(sp)
			try[count].point = p
			try[count].bonus = 2
			sp = ""

		case t == 'S':
			p, _ := strconv.Atoi(sp)
			try[count].point = p
			try[count].bonus = 1
			sp = ""

		case t == 'T':
			p, _ := strconv.Atoi(sp)
			try[count].point = p
			try[count].bonus = 3
			sp = ""

		case t == '*' || t == '#':
			try[count].option = t
		}

	}
	fmt.Printf("input : %+v\n", try)
	for i := 0; i < 3; i++ {
		temp := int(math.Pow(float64(try[i].point), float64(try[i].bonus)))
		if try[i].option == '*' {
			try[i].result = temp * 2
			if i > 0 {
				try[i-1].result = try[i-1].result * 2
			}
		} else if try[i].option == '#' {
			try[i].result = temp * (-1)
		} else {
			try[i].result = temp
		}
	}
	result = try[0].result + try[1].result + try[2].result
	return result
}
