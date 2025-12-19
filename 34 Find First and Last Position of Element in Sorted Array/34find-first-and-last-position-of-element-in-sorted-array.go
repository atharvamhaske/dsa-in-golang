func searchRange(nums []int, target int) []int {
    first:= findFirst(nums, target)
    second := findSec(nums, target)
    return []int{first, second}
}

func findFirst(a []int, target int) int {
    left, right := 0, len(a)-1;

    result := -1

    for left <= right {
        mid := left + (right - left)/2

        if a[mid] == target {
            result = mid
            right = mid - 1
        } else if a[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return result
}

func findSec(b []int, target int) int {
    left, right := 0, len(b) - 1

    result := -1;

    for left <= right {
        mid := left + (right - left)/2

        if b[mid] == target {
            result = mid
            left = mid + 1
        } else if b[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return result

}