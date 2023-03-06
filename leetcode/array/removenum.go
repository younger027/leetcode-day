package array

func RemoveElement(nums []int, val int) int {

	/*
		left := 0
		    for _, v := range nums { // v 即 nums[right]
		        if v != val {
		            nums[left] = v
		            left++
		        }
		    }
		    return left

	*/
	j := len(nums) - 1
	for i := 0; i < j; i++ {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			i -= 1
			j -= 1
		}
	}

	return j + 1
}
