package main

func twoSum(nums []int, target int) []int {

	s := set{}
	s.Init()

	for index, val := range nums {

		if s.has(target - val) {
			return []int{s.get(target - val), index}
		}

		s.add(val)
	}

	return nil

}

func twoSum2(nums []int, target int) []int {
	// Space: O(n)
	s := make(map[int]int)

	// Time: O(n)
	for idx, num := range nums {
		// Time: O(1)
		if pos, ok := s[target-num]; ok {
			return []int{pos, idx}
		}
		s[num] = idx
	}
	return []int{}
}

func main() {

	nums := []int{2, 7, 11, 15}
	target := 9

	twoSum(nums, target)

}

type set struct {
	items map[int]int
}

func (s *set) Init() *set {
	return &set{
		items: make(map[int]int),
	}
}

func (s *set) add(item int) {
	s.items[item] = item
}

func (s *set) has(i int) bool {
	_, exists := s.items[i]
	return exists
}

func (s *set) get(i int) int {
	idx, exists := s.items[i]
	if !exists {
		return -1
	}
	return idx
}
