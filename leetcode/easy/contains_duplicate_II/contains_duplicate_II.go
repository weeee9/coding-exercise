package containsduplicate

func containsNearbyDuplicate(nums []int, k int) bool {
	set := map[int]struct{}{}

	for i, num := range nums {
		// if i is greater than k, it means we have moved beyond
		// the adjacent distance and need to remove the number
		// from the previous window
		if i > k {
			delete(set, nums[i-k-1])
		}

		if _, ok := set[num]; ok {
			return true
		}

		set[num] = struct{}{}
	}

	return false
}
