func majorityElement(nums []int) int {
	el := 0
    ct := 0 

    for _, val := range nums {
        if ct == 0 {
            el = val
        } 

        if val == el {
            ct ++
        } else {
            ct --
        }
    }
    return el
}

//moores voting algorithm