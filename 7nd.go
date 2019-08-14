package kakao2018

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parse7time(str string) (int64, int64) {
	tmp := strings.Split(str, " ")

	t, _ := time.Parse("2006-01-02 15:04:05.000", tmp[0]+" "+tmp[1])
	// fmt.Println("in :", str)
	end := t.UnixNano() / int64(time.Millisecond)

	// fmt.Println("time :", t, end)

	s, _ := strconv.ParseFloat(tmp[2][:len(tmp[2])-1], 64)

	// fmt.Println("du :", tmp[2], tmp[2][:len(tmp[2])-1], s)
	return end - int64(s*1000.0) + 1, end
}

type lt struct {
	start, end int64
}

func (l *lt) printtime() string {
	///start := time.Unix(0, l.start*int64(time.Millisecond))
	//end := time.Unix(0, l.end*int64(time.Millisecond))
	return fmt.Sprintf("%d <==> %d", l.start, l.end)
}

func (l *lt) between(s, e int64) bool {
	fmt.Printf("check %d <==> %d : %s result=%t\n", s, e, l.printtime(), l.start > e || l.end < s)
	if l.start > e || l.end < s {
		return false
	}
	return true
}

func k7th(in []string) int {
	var tt []*lt
	for _, s := range in {
		var t lt
		t.start, t.end = parse7time(s)
		// fmt.Println(s, t.printtime())
		tt = append(tt, &t)
	}

	var result int
	for i := 0; i < len(tt); i++ {
		fmt.Println("====================================================")

		{
			r := 1

			s := tt[i].start - 1000 + 1
			e := tt[i].start

			for j := 0; j < len(tt); j++ {
				if i == j {
					continue
				}
				if tt[j].between(s, e) {
					r++
					fmt.Printf("%s : %s check ok\n", tt[i].printtime(), tt[j].printtime())
				}

			}

			if result < r {
				result = r
			}

		}

		{
			r := 1

			s := tt[i].start
			e := tt[i].start + 1000 - 1

			for j := 0; j < len(tt); j++ {
				if i == j {
					continue
				}
				if tt[j].between(s, e) {
					r++
					fmt.Printf("%s : %s check ok\n", tt[i].printtime(), tt[j].printtime())
				}

			}

			if result < r {
				result = r
			}

		}

		{
			r := 1

			s := tt[i].end - 1000 + 1
			e := tt[i].end

			for j := 0; j < len(tt); j++ {
				if i == j {
					continue
				}
				if tt[j].between(s, e) {
					r++
					fmt.Printf("%s : %s check ok\n", tt[i].printtime(), tt[j].printtime())
				}

			}

			if result < r {
				result = r
			}

		}

		{
			r := 1

			s := tt[i].end
			e := tt[i].end + 1000 - 1

			for j := 0; j < len(tt); j++ {
				if i == j {
					continue
				}
				if tt[j].between(s, e) {
					r++
					fmt.Printf("%s : %s check ok\n", tt[i].printtime(), tt[j].printtime())
				}

			}

			if result < r {
				result = r
			}

		}

	}

	return result
}
