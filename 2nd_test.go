package kakao2018

import "testing"
import "github.com/stretchr/testify/assert"

func Test2nd(t *testing.T) {
	t.Log("start kakao 2nd")

	tc := []struct {
		in  string
		out int
	}{
		{"1S2D*3T", 37},
		{"1D2S#10S", 9},
		{"1D2S0T", 3},
		{"1S*2T*3S", 23},
		{"1D#2S*3S", 5},
		{"1T2D3D#	", -4},
		{"1D2S3T*", 59},
	}

	for _, tt := range tc {
		t.Run(tt.in, func(t *testing.T) {
			result := k2nd(tt.in)
			assert.Equal(t, tt.out, result)
		})
	}
}
