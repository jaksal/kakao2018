package kakao2018

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test4th(t *testing.T) {
	t.Log("start kakao 4th")

	tc := []struct {
		n, t, m   int
		timetalbe []string
		out       string
	}{
		{1, 1, 5, []string{"08:00", "08:01", "08:02", "08:03"}, "09:00"},
		{2, 10, 2, []string{"09:10", "09:09", "08:00"}, "09:09"},
		{2, 1, 2, []string{"09:00", "09:00", "09:00", "09:00"}, "08:59"},
		{1, 1, 5, []string{"00:01", "00:01", "00:01", "00:01", "00:01"}, "00:00"},
		{1, 1, 1, []string{"23:59"}, "09:00"},
		{10, 60, 45, []string{"23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59"}, "18:00"},
	}

	for i, tt := range tc {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := k4th(tt.n, tt.t, tt.m, tt.timetalbe)
			assert.Equal(t, tt.out, result)
		})
	}
}
