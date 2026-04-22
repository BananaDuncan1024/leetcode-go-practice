package main

func main() {

	obj := Constructor()
	obj.Ping(1)
	obj.Ping(100)
	obj.Ping(3000)
}

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	// 1. 新時間入隊
	this.queue = append(this.queue, t)

	// 2. 檢查最前面的資料，如果過期了 (小於 t - 3000)，就把它出隊 (Pop)
	for len(this.queue) > 0 && this.queue[0] < t-3000 {
		// 利用 Slice 的重新切片來模擬 Pop
		this.queue = this.queue[1:]
	}

	// 3. 剩下的長度就是答案
	return len(this.queue)
}

type RecentCounter2 struct {
	pings []int
	head  int // 指向當前 3000 毫秒窗口的起點
	tail  int // 指向下一個可以寫入的位置
}

func Constructor2() RecentCounter2 {
	// 題目保證最多 10000 次呼叫，直接分配好，避開所有 runtime 擴容開銷
	return RecentCounter2{
		pings: make([]int, 10000),
		head:  0,
		tail:  0,
	}
}

func (this *RecentCounter2) Ping2(t int) int {
	// 1. 將資料寫入 tail 位置，然後 tail 往後移
	this.pings[this.tail] = t
	this.tail++

	// 2. 如果 head 指向的時間已經過期，head 就往後移（收縮滑動窗口）
	for this.pings[this.head] < t-3000 {
		this.head++
	}

	// 3. 有效的數量就是 tail 減去 head
	return this.tail - this.head
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
