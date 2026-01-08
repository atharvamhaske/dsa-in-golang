func rotate(nums []int, k int)  {
    n := len(nums)
    k %= n
    
    reverse(nums, 0, n-1)
    reverse(nums, 0, k-1)
    reverse(nums, k, n-1)
    
}

func reverse(nums []int, st int, end int) {
    for st < end {
        nums[st], nums[end] = nums[end], nums[st]
        st++
        end--
    }
}