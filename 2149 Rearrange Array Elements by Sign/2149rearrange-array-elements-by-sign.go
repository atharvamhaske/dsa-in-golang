func rearrangeArray(nums []int) []int {
    n := len(nums)
    posIdx, negIdx := 0, 1
    res := make([]int, n)

    for _, val := range nums {
        if val > 0 {
            res[posIdx] = val
            posIdx +=2
        } else {
            res[negIdx] = val
            negIdx +=2
        }
    }
    return res
}

//classic two pointers store at 0, 1 and one observation that make positive and negative alternative as
//2n and (2n + 1) and dont write else if write else just