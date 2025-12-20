func minOperations(nums []int) int {
    n := len(nums)
    freq := make(map[int]int)

    for _, v := range nums {
        freq[v]++
    }

    duplicates := 0
    for _, c := range freq {
        if c > 1 {
            duplicates++
        }
    }

    if duplicates == 0 {
        return 0
    }

    ops := 0
    i := 0

    for i < n && duplicates > 0 {
        for j := 0; j < 3 && i < n; j++ {
            val := nums[i]
            freq[val]--

 
            if freq[val] == 1 {
                duplicates--
            }

            i++
        }
        ops++
    }

    return ops
}