func sortColors(nums []int) {
	n := len(nums)
	low, mid := 0, 0
	high := n - 1

    // Dutch National Flag ALgorithm

	for mid <= high {
		if nums[mid] == 0 {
			nums[low], nums[mid] = nums[mid], nums[low]
			mid++
			low++
		} else if nums[mid] == 2 {
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		} else {
			mid++
		}
	}
}