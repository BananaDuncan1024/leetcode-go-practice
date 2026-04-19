package main

func main() {

	timeSeries := []int{1, 4, 9}

	duration := 2

	findPoisonedDuration(timeSeries, duration)
}

func findPoisonedDuration(timeSeries []int, duration int) int {

	totaltime := 0

	// 這邊需注意 array 要處理最後的終止條件 避免 超過array 後還取下一個位置導致memory leak
	for i := 0; i < len(timeSeries); i++ {

		if i == len(timeSeries)-1 {
			totaltime += duration
			break
		}

		// 這邊邏輯沒問題 只要下一個時間點還在中毒時間內 就把中毒時間延長到下一個時間點 否則就加上固定的中毒時間
		if (timeSeries[i] + duration) > timeSeries[i+1] {
			totaltime += timeSeries[i+1] - timeSeries[i]
		} else {
			totaltime += duration
		}
	}
	return totaltime
}
