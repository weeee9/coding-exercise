package removing

func removeStars(s string) string {
	result := make([]byte, 0, len(s))
	for i := range s {
		if s[i] == '*' {
			result = result[:len(result)-1]
			continue
		}
		result = append(result, s[i])
	}

	return string(result)
}
