/*
@Time : 2019/7/23 11:09
@Author : 一条小咸鱼
@File :
@Software: GoLand
*/
package utils

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	nums := []int{1, 2, 3}
	fmt.Printf("%p\n", &nums)
	a(nums)

	fmt.Println(nums)
}
func removeDuplicates(nums []int) int {

	length := len(nums)
	left := 0
	if length < 2 {
		return length
	}
	for right := 1; right < length; {
		if nums[left] == nums[right] {
			right++
			continue
		}
		left++
		nums[left] = nums[right]

	}

	return left + 1
}
func a(arr []int) {
	fmt.Printf("%p\n", &arr)
	arr = []int{1, 2, 3, 4}

}
func TestGetAllDir(t *testing.T) {

	var list []string
	strings, e := GetAllDir(list, "E:\\server\\app")
	if e != nil {
		t.Error(e)
		return
	}

	for _, v := range strings {
		fmt.Println(v)
	}

}
