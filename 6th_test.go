package kakao2018

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test6nd(t *testing.T) {
	t.Log("start kakao 6nd")

	tc := []struct {
		m, n  int
		board []string
		out   int
	}{
		{4, 5, []string{"CCBDE", "AAADE", "AAABF", "CCBBF"}, 14},
		{6, 6, []string{"TTTANT", "RRFACC", "RRRFCC", "TRRRAA", "TTMMMF", "TMMTTJ"}, 15},
	}

	for i, tt := range tc {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := k6nd(tt.m, tt.n, tt.board)
			assert.Equal(t, tt.out, result)
		})
	}
}
