package sqrt

func MySqrt(x int) int {

	if x < 2 {
		return x // The square root of 0 is 0 and 1 is 1
	}

	left, right := 2, x/2 // We can safely start searching from 2 to x/2

	for left <= right {
		mid := left + (right-left)/2
		midSquared := mid * mid

		if midSquared == x {
			return mid // Found the exact square root
		} else if midSquared < x {
			left = mid + 1 // Move to the right half
		} else {
			right = mid - 1 // Move to the left half
		}
	}

	return right // Right will be the largest integer such that right*right <= x

}
