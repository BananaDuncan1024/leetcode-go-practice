package main


func main() {

	array := []int {1,1,1,3,3,4,3,2,4,2}

	containsDuplicate(array)
}

func containsDuplicate(nums []int) bool {

    m := make(map[int]bool)

    for i := 0; i < len(nums); i++ {
	
	_ , exist := m[nums[i]] 

		if (exist) {
			return true
		}

	m[nums[i]] = true
    }

    return false;
    
}

func containsDuplicate2(nums []int) bool {
    set := make(map[int]struct{})
    for _, num := range nums {
        if _, hasNum := set[num]; hasNum {
            return true
        }
        set[num] = struct{}{}
    }
    return false
}