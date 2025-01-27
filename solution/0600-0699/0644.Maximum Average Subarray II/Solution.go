func findMaxAverage(nums []int, k int) float64 {
	eps := 1e-5
	l, r := 1e9, -1e9
	for _, x := range nums {
		l = math.Min(l, float64(x))
		r = math.Max(r, float64(x))
	}
	check := func(v float64) bool {
		s := 0.0
		for _, x := range nums[:k] {
			s += float64(x) - v
		}
		if s >= 0 {
			return true
		}
		t := 0.0
		mi := 0.0
		for i := k; i < len(nums); i++ {
			s += float64(nums[i]) - v
			t += float64(nums[i-k]) - v
			mi = math.Min(mi, t)
			if s >= mi {
				return true
			}
		}
		return false
	}
	for r-l >= eps {
		mid := (l + r) / 2
		if check(mid) {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}