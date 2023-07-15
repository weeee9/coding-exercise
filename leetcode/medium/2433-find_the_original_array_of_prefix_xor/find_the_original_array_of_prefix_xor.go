package find

func findArray(pref []int) []int {
	ans := 0
	for i := range pref {
		pref[i] = ans ^ pref[i]
		ans = ans ^ pref[i]
	}

	return pref
}
