package kakao2018

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test5nd(t *testing.T) {
	t.Log("start kakao 5nd")

	tc := []struct {
		str1, str2 string
		out        int
	}{
		{"FRANCE", "french", 16384},
		{"handshake", "shake hands", 65536},
		{"aa1+aa2", "AAAA12", 43690},
		{"E=M*C^2", "e=m*c^2", 65536},
	}

	for i, tt := range tc {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := k5nd(tt.str1, tt.str2)
			assert.Equal(t, tt.out, result)
		})
	}
}
