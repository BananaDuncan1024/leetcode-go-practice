package main

func main() {

	uniqueOccurrences([]int{1, 2, 2, 1, 1, 3})
}

func uniqueOccurrences(arr []int) bool {
	// 統計每個數字出現的次數
	counts := make(map[int]int, len(arr)/2)
	for _, num := range arr {
		counts[num]++
	}

	// 檢查次數是否唯一，使用 struct{} 節省記憶體
	occurrences := make(map[int]struct{}, len(counts))
	for _, count := range counts {
		if _, exists := occurrences[count]; exists {
			return false
		}
		occurrences[count] = struct{}{}
	}

	return true
}
