package kakao2018

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test7th(t *testing.T) {
	t.Log("start kakao 7th")

	tc := []struct {
		in  []string
		out int
	}{
		{[]string{"2016-09-15 01:00:04.001 2.0s", "2016-09-15 01:00:07.000 2s"}, 1},
		{[]string{"2016-09-15 01:00:04.002 2.0s", "2016-09-15 01:00:07.000 2s"}, 2},
		{[]string{"2016-09-15 20:59:57.421 0.351s", "2016-09-15 20:59:58.233 1.181s", "2016-09-15 20:59:58.299 0.8s", "2016-09-15 20:59:58.688 1.041s", "2016-09-15 20:59:59.591 1.412s", "2016-09-15 21:00:00.464 1.466s", "2016-09-15 21:00:00.741 1.581s", "2016-09-15 21:00:00.748 2.31s", "2016-09-15 21:00:00.966 0.381s", "2016-09-15 21:00:02.066 2.62s"}, 7},
	}

	for i, tt := range tc {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := k7th(tt.in)
			assert.Equal(t, tt.out, result)
		})
	}
}
