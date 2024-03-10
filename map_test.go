package underscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
)

func TestMap(t *testing.T) {
	nums := []int{1, 2, 3}
	f := func(n int) int {
		return n * n
	}
	want := []int{1, 4, 9}
	assert.Equal(t, want, u.Map(nums, f))
}

func benchmarkMapN(n int, b *testing.B) {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = u.Map(nums, func(num int) int {
			return num
		})
	}
}

func BenchmarkMap100(b *testing.B)   { benchmarkMapN(100, b) }
func BenchmarkMap1000(b *testing.B)  { benchmarkMapN(1000, b) }
func BenchmarkMap10000(b *testing.B) { benchmarkMapN(10000, b) }
