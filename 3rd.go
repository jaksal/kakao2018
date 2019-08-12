package kakao2018

import (
	"container/list"
	"strings"
)

func k3rd(cacheSize int, citis []string) int {
	var result int

	cache := list.New()
	for _, c := range citis {

		// find
		hit := false
		for ci := cache.Front(); ci != nil; ci = ci.Next() {
			if r := strings.Compare(ci.Value.(string), strings.ToUpper(c)); r == 0 {
				hit = true
				break
			}
		}

		if hit {
			result++
		} else {
			result += 5
			cache.PushBack(strings.ToUpper(c))
			if cache.Len() > cacheSize {
				cache.Remove(cache.Front())
			}
		}
	}

	return result
}
