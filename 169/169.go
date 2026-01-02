package main

import "sort"

func main() {

	nums := []int{3, 2, 3}

	ans := majorityElement(nums)
	print(ans)

}

func majorityElement(nums []int) int {

	temp := map[int]int{}

	for _, val := range nums {

		value, isExist := temp[val]
		if isExist {
			temp[val] = value + 1
		} else {
			temp[val] = 1
		}
	}

	max := 0

	for key, value := range temp {
		if value > max {
			max = value
			if max > len(nums)/2 {
				return key
			}
		}
	}
	return max
}

// sort
func majorityElement3(nums []int) int {
	// 先排序，中间位置的数字就是多数元素
	sort.Ints(nums)
	return nums[(len(nums)-1)/2]
}

// 摩爾投票法(Moore voting)
// 這題可以用一種解法，叫做摩爾投票法(Moore voting)，也叫多數投票法。
// 它的原理是，相鄰的元素做比較，若不相等，將它們從陣列中剔除，如此一來最後剩下的，就會是最多數的，因為有超過半數的保證。
// 作法：
// 宣告變數majority為陣列的第一個元素，計數器count初始為0，
// 遍歷陣列，元素相等count+1，不相等count-1，
// 當count=0表示最多元素將會替換，遍歷完後，最後majority中的元素，將是最多數
func majorityElement2(nums []int) int {
	result, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			result = v
		}
		if v != result {
			count--
		} else {
			count++
		}
	}
	return result
}
