// File: array_test.go
// Created Time: 2022-12-29
// Author: GuoWei (gongguowei01@gmail.com), cathay (cathaycchen@gmail.com)

package array

/**
我们将 Go 中的 Slice 切片看作 Array array。因为这样可以
降低理解成本，利于我们将关注点放在数据结构与算法上。
*/

import (
	"fmt"
	"testing"
)

/* Driver Code */
func TestArray(t *testing.T) {
	/* 初始化数组 */
	var arr []int
	fmt.Println("array arr =", arr) //	array arr = []
	nums := []int{1, 3, 2, 5, 4}
	fmt.Println("array nums =", nums) // array nums = [1 3 2 5 4]

	/* 随机访问 */
	randomNum := randomAccess(nums)
	fmt.Println("在 nums 中获取随机元素", randomNum) // 3

	/* 长度扩展 */
	nums = extend(nums, 3)
	fmt.Println("将数组长度扩展至 8 ，得到 nums =", nums) // nums = [1 3 2 5 4 0 0 0]

	/* 插入元素 */
	insert(nums, 6, 3)
	fmt.Println("在索引 3 处插入数字 6 ，得到 nums =", nums) // nums = [1 3 2 6 5 4 0 0]

	/* 删除元素 */
	remove(nums, 2)
	fmt.Println("删除索引 2 处的元素，得到 nums =", nums) // nums = [1 3 6 5 4 0 0 0]

	/* 遍历数组 */
	traverse(nums)

	/* 查找元素 */
	index := find(nums, 3)
	fmt.Println("在 nums 中查找元素 3 ，得到索引 =", index) // 1
}
