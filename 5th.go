package kakao2018

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

// parse
func parse(str string) []string {
	var result []string

	var temp string
	for _, r := range str {
		if !unicode.IsLetter(r) {
			temp = ""
			continue
		}

		if len(temp) == 0 {
			temp = string(r)
		} else if len(temp) == 1 {
			result = append(result, temp+string(r))
			temp = string(r)
		}
	}
	return result
}

func k5th(str1, str2 string) int {

	fmt.Printf("input:%s , %s\n", str1, str2)
	arr1 := parse(str1)
	arr2 := parse(str2)
	fmt.Printf("parse %+v , %+v\n", arr1, arr2)

	if len(arr1) == 0 && len(arr2) == 0 {
		return 1 * 65536
	}

	var inter, union []string

	sort.Slice(arr1, func(i, j int) bool {
		return strings.ToUpper(arr1[i]) < strings.ToUpper(arr1[j])
	})
	sort.Slice(arr2, func(i, j int) bool {
		return strings.ToUpper(arr2[i]) < strings.ToUpper(arr2[j])
	})
	fmt.Printf("sorted parse %+v , %+v\n", arr1, arr2)

	var i, j int
	for i < len(arr1) && j < len(arr2) {
		c := strings.Compare(strings.ToUpper(arr1[i]), strings.ToUpper(arr2[j]))

		fmt.Printf("i:%d [%s] j:%d [%s] c=%d \n", i, arr1[i], j, arr2[j], c)

		if c < 0 {
			union = append(union, arr1[i])
			i++
		} else if c > 0 {
			union = append(union, arr2[j])
			j++
		} else {
			union = append(union, arr1[i])
			inter = append(inter, arr1[i])
			i++
			j++
		}
		fmt.Printf("union=%+v inter=%+v\n", union, inter)
	}
	if i < len(arr1) {
		union = append(union, arr1[i:]...)
	}
	if j < len(arr2) {
		union = append(union, arr2[j:]...)
	}
	fmt.Printf("union:%+v inter:%+v", union, inter)

	return int(float64(len(inter)) / float64(len(union)) * 65536.0)
}
