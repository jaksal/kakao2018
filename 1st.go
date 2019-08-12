package kakao2018

import (
	"fmt"
	"strings"
)

var re = strings.NewReplacer("1", "#", "0", " ")

func printbit(n int, d int) string {
	return re.Replace(fmt.Sprintf(fmt.Sprintf("%%0%db", n), d))
}

func k1st(n int, arr1, arr2 []int) []string {
	var result []string
	for i := 0; i < n; i++ {
		fmt.Printf("a : %s\n", printbit(n, arr1[i]))
		fmt.Printf("b : %s\n", printbit(n, arr2[i]))
		fmt.Printf("r : %s\n", printbit(n, arr1[i]|arr2[i]))
		result = append(result, printbit(n, arr1[i]|arr2[i]))
	}
	return result
}
