package kakao2018

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func parsetime(str string) int {
	tm := strings.Split(str, ":")
	h, _ := strconv.Atoi(tm[0])
	m, _ := strconv.Atoi(tm[1])

	return (h-9)*60 + m
}

func parsetime2(t int) string {
	var h, m int

	h = t/60 + 9
	m = t % 60
	if m < 0 {
		h--
		m = 60 + m
	}

	return fmt.Sprintf("%02d:%02d", h, m)
}

func k4nd(n, t, m int, timetable []string) string {
	fmt.Printf("n:%d t:%d m:%d timetable:%+v\n", n, t, m, timetable)
	sort.Slice(timetable, func(i, j int) bool {
		return parsetime(timetable[i]) < parsetime(timetable[j])
	})

	bus := make(map[int][]int)

	var busorder []int
	start := 0

	for ni := 0; ni < n; ni++ {
		// 한명씩 태워보자..
		var bm []int
		var i int
		for _, tt := range timetable {
			d := parsetime(tt)
			if d <= start {
				bm = append(bm, d)
				i++
			}
			if len(bm) >= m {
				// 만차.
				break
			}
		}
		timetable = timetable[i:]

		bus[start] = bm
		busorder = append(busorder, start)

		start = start + t
	}

	// 맨마지막 시간이 없다면 전차로 이동.
	for i := len(busorder) - 1; i >= 0; i-- {
		// 버스를 거꾸로 돌면서. 빈자리 있으면.. ok
		order := busorder[i]
		fmt.Printf("order:%d bus:%+v\n", order, bus[order])
		if len(bus[order]) < m {
			return parsetime2(order)
		}

		/*
			last := bus[order][len(bus[order])-1]
			if i> 0 && last == busorder[i-1] {

			}
		*/
		// 없으면. 막차 탄사람중에 맨 마지막 시간..
		return parsetime2(bus[order][len(bus[order])-1] - 1)
	}

	// 첫차도 꽉찼다면..
	return parsetime2(0)
}
