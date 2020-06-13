package sum

import (
	"fmt"
	"runtime"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	nums := []int{4, 5, 6, 7, 8, 9, 10}
	target := 10 // nums[1] + nums[7]
	sum := TwoSum1(nums, target)
	fmt.Println("sum = ", sum)
	fmt.Println("Last GC was:", stats.LastGC)
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{4, 5, 6, 7, 8, 9, 10}
	target := 10 // nums[1] + nums[7]
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		// BenchmarkTwoSum-8   	41733980	        24.1 ns/op	      16 B/op	       1 allocs/op
		// BenchmarkTwoSum-8   	16765234	        65.4 ns/op	      16 B/op	       1 allocs/op
		TwoSum2(nums, target)
	}
}
