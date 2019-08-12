package kakao2018

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test3rd(t *testing.T) {
	t.Log("start kakao 3rd")

	tc := []struct {
		s   int
		in  []string
		out int
	}{
		{3, []string{"Jeju", "Pangyo", "Seoul", "NewYork", "LA", "Jeju", "Pangyo", "Seoul", "NewYork", "LA"}, 50},
		{3, []string{"Jeju", "Pangyo", "Seoul", "Jeju", "Pangyo", "Seoul", "Jeju", "Pangyo", "Seoul"}, 21},
		{2, []string{"Jeju", "Pangyo", "Seoul", "NewYork", "LA", "SanFrancisco", "Seoul", "Rome", "Paris", "Jeju", "NewYork", "Rome"}, 60},
		{5, []string{"Jeju", "Pangyo", "Seoul", "NewYork", "LA", "SanFrancisco", "Seoul", "Rome", "Paris", "Jeju", "NewYork", "Rome"}, 52},
		{2, []string{"Jeju", "Pangyo", "NewYork", "newyork"}, 16},
		{0, []string{"Jeju", "Pangyo", "Seoul", "NewYork", "LA"}, 25},
	}

	for i, tt := range tc {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := k3rd(tt.s, tt.in)
			assert.Equal(t, tt.out, result)
		})
	}
}
