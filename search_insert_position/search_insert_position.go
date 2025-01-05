package searchinsertposition

func SearchInsert(nums []int, target int) int {

	// nums := []int{1, 3, 5, 6}
	// target := 7

	leftIndex, rightIndex := 0, len(nums)-1

	for leftIndex <= rightIndex {
		mid := leftIndex + (rightIndex-leftIndex)/2

		if nums[mid] == target {
			return mid // Target found
		} else if nums[mid] < target { 
			leftIndex = mid + 1 // Search in the right half
		} else {
			rightIndex = mid - 1 // Search in the left half
		}
	}

	// // Target not found, return the insert position
	return leftIndex

}
