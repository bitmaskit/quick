package quick

func Sort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	sort(nums, 0, len(nums)-1)
}

func sort(nums []int, lo, hi int) {
	if lo < hi {
		p := partition(nums, lo, hi)
		sort(nums, lo, p-1)
		sort(nums, p+1, hi)
	}
}

func partition(nums []int, lo, hi int) int {
	i, pivot := lo, nums[hi]
	for j := lo; j < hi; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[hi] = nums[hi], nums[i]
	return i
}