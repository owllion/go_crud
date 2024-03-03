package escape

import (
	"fmt"
	"math/rand"
)

func gen8191() {
	nums := make([]int, 8191)
	for i := 0; i < 8191; i++ {
		nums[i] = i
	}
}
func gen8192() {
	nums := make([]int, 8192)
	for i := 0; i < 8192; i++ {
		nums[i] = i
	}
}
func gen8193() {
	nums := make([]int, 8193)
	for i := 0; i < 8193; i++ {
		nums[i] = i
	}
}

func genRandom(n int) {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Int()
	}
}
func getInterface(args ...interface{}) {
	res := 0
	if len(args) != 0 {
		if val, ok := args[0].(int); ok {
			res += val
		}
		return
	}
}
func getSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total

}

func closureEscape() func() int {
	x := 5
	return func() int {
		x *= 5
		return x
	}
}
func main() {
	gen8191()
	gen8192()      //64kb not escape
	gen8193()      //64.07kb,escape
	genRandom(200) //not sure the size, escape
	n := 8
	getInterface(n, n, n) //not escape
	closure := closureEscape()
	fmt.Println(closure())
}
